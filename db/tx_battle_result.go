package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	"strconv"
	"time"

	db_battle "github.com/yeom-c/game-server-grpc-go/db/sqlc/battle"
	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) TxConfirmBattleResults(ctx context.Context, txBattleQueries *db_battle.Queries, txGameQueries *db_game.Queries, gameDb, accountUserId int32, battleResultList []db_battle.BattleResult) (matchPoint, addPoint int32, newReward model_pb.Reward, err error) {
	var battleTx *sql.Tx
	var gameTx *sql.Tx
	if txBattleQueries == nil {
		battleTx, err = s.BattleDb.Begin()
		if err != nil {
			return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer battleTx.Rollback()

		txBattleQueries = s.BattleQueries.WithTx(battleTx)
	}
	if txGameQueries == nil {
		gameTx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer gameTx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(gameTx)
	}

	if len(battleResultList) == 0 {
		return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 점수 조회.
	battleUser, err := txBattleQueries.GetUserByAccountUserId(ctx, accountUserId)
	if err != nil {
		return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 티어 정보.
	tierStaticList, err := s.StaticDataQueries.GetTiers(ctx)
	if err != nil {
		return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrTier.String())
	}
	if len(tierStaticList) == 0 {
		return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrTier.String())
	}
	tierPointTable := []int{}
	for _, tierStatic := range tierStaticList {
		tierPointTable = append(tierPointTable, int(tierStatic.RankUpNecessary))
	}
	tierStaticIndex := helper.GetLevel(tierPointTable, int(battleUser.MatchPoint)) - 1
	tierStatic := tierStaticList[tierStaticIndex]
	winPoint := tierStatic.WinPoint
	losePoint := tierStatic.LosePoint
	winDropEnumId := tierStatic.WinReward
	loseDropEnumId := tierStatic.LoseReward

	// 보상, 점수 업데이트.
	battleResultIdList := []int32{}
	dropEnumIdList := []string{}
	var addWin int32
	var addLose int32
	winCnt := battleUser.MatchWin
	for _, battleResult := range battleResultList {
		if battleResult.ConfirmedAt.Valid {
			continue
		}

		battleResultIdList = append(battleResultIdList, battleResult.ID)
		if battleResult.Result == 0 {
			dropEnumIdList = append(dropEnumIdList, loseDropEnumId)
			addPoint += losePoint
			addLose += 1
		} else if battleResult.Result == 1 {
			dropEnumIdList = append(dropEnumIdList, winDropEnumId)
			addPoint += winPoint
			addWin += 1
			winCnt += 1

			// TODO: costume condition check 처리 필요.(condition 기획에 따라 테이블 설계후 저장, 체크, 지급)
			// 코스튬 컨티션 체크 후 지급.
			deckCharacters, err := s.TxGetDeckCharacters(ctx, txGameQueries, gameDb, accountUserId, battleResult.DeckID)
			if err != nil {
				return matchPoint, addPoint, newReward, err
			}
			deckCharactersEnumId := []string{}
			for _, deckCharacter := range deckCharacters {
				deckCharactersEnumId = append(deckCharactersEnumId, deckCharacter.EnumID)
			}
			if len(deckCharactersEnumId) > 0 {
				costumeStaticList, err := s.StaticDataQueries.GetCostumesByConditionAndCharacterEnumIds(ctx, enum.Costume_Condition_DUEL_WIN.String(), deckCharactersEnumId)
				if err != nil {
					return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrCostumeNotFound.String())
				}
				createCostumeParams := []db_game.CreateCostumeParams{}
				for _, costumeStatic := range costumeStaticList {
					conditionValue, err := strconv.Atoi(costumeStatic.ConditionValue)
					if err != nil {
						return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrCostume.String())
					}
					if int32(conditionValue) <= winCnt {
						createCostumeParams = append(createCostumeParams, db_game.CreateCostumeParams{
							AccountUserID:   accountUserId,
							EnumID:          costumeStatic.EnumID,
							CharacterEnumID: costumeStatic.Character,
							State:           int32(enum.Equip_State_UNEQUIP),
						})
					}
				}
				if len(createCostumeParams) > 0 {
					_, err = txGameQueries.CreateCostumes(ctx, createCostumeParams)
					if err != nil {
						return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
					}
				}
			}
		}
	}

	// 보상 지급.
	if len(dropEnumIdList) > 0 {
		err = s.GetDropRewardList(ctx, &newReward, dropEnumIdList)
		if err != nil {
			return matchPoint, addPoint, newReward, err
		}
		err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
		if err != nil {
			return matchPoint, addPoint, newReward, err
		}
	}

	// 점수 업데이트.
	if battleUser.MatchPoint+addPoint < 0 {
		addPoint = -battleUser.MatchPoint
	}
	_, err = txBattleQueries.UpdateUserMatchResultByAccountUserId(ctx, db_battle.UpdateUserMatchResultByAccountUserIdParams{
		AccountUserID: accountUserId,
		AddMatchPoint: addPoint,
		AddMatchWin:   addWin,
		AddMatchLose:  addLose,
	})
	if err != nil {
		return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	matchPoint = battleUser.MatchPoint + addPoint

	// 확인 처리.
	_, err = txBattleQueries.UpdateBattleResultsConfirmedAt(ctx, sql.NullTime{Time: time.Now().UTC(), Valid: true}, battleResultIdList)
	if err != nil {
		return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	if battleTx != nil {
		err := battleTx.Commit()
		if err != nil {
			return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}
	if gameTx != nil {
		err := gameTx.Commit()
		if err != nil {
			return matchPoint, addPoint, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return matchPoint, addPoint, newReward, nil
}
