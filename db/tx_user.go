package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/yeom-c/game-server-grpc-go/helper"

	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
)

func (s *store) TxCreateUser(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32) (userId int32, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return userId, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		// 생성된 tx queries 를 사용해야 같은 트랜잭션에서 처리됨
		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	result, err := txGameQueries.CreateUser(ctx, accountUserId)
	if err != nil {
		return userId, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return userId, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	userId = int32(lastId)

	startDataList, err := s.StaticDataQueries.GetStarts(ctx)
	if err != nil {
		return userId, err
	}

	// 계정 생성시 지급 아이템.
	rewardStaticData := RewardStaticData{}
	rewardData := DropRewardData{}

	for _, startData := range startDataList {
		rewardType := enum.GetCommon_Type(startData.CeCommonType)
		rewardEnumId := startData.RewardID
		rewardCount := startData.Value
		err = s.SetRewardStaticData(ctx, &rewardStaticData, rewardType, rewardEnumId)
		if err != nil {
			return userId, err
		}

		rewardData.TypeArr = append(rewardData.TypeArr, rewardType.String())
		rewardData.DropArr = append(rewardData.DropArr, rewardEnumId)
		rewardData.ValueArr = append(rewardData.ValueArr, rewardCount)
	}

	reward := model_pb.Reward{}
	for i, t := range rewardData.TypeArr {
		rewardType := enum.GetCommon_Type(t)
		rewardEnumId := rewardData.DropArr[i]
		rewardValue := rewardData.ValueArr[i]
		err := s.SetRewardList(ctx, rewardStaticData, &reward, rewardType, rewardEnumId, rewardValue)
		if err != nil {
			return userId, err
		}
	}
	err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &reward)
	if err != nil {
		return userId, err
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return userId, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}
	return userId, nil
}

func (s *store) TxDailyReset(ctx context.Context, txGameQueries *db_game.Queries, now time.Time, gameDb, accountUserId int32) error {
	var tx *sql.Tx
	var err error
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	user, err := txGameQueries.GetUserByAccountUserId(ctx, accountUserId)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 매일 초기화 처리.
	updated := false
	if _, ok := enum.ServerEnum["reset_time_daily_0"]; !ok {
		return helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}
	dailyResetTime0 := enum.ServerEnum["reset_time_daily_0"].(int)
	resetDate0 := time.Date(now.Year(), now.Month(), now.Day(), dailyResetTime0, 0, 0, 0, time.UTC)
	if user.DailyResetAt.Before(resetDate0) {
		updated = true
	}

	if updated {
		_, err = txGameQueries.UpdateUserDailyResetAtByAccountUserId(ctx, db_game.UpdateUserDailyResetAtByAccountUserIdParams{
			AccountUserID: accountUserId,
			DailyResetAt:  now,
		})
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}
	return nil
}

func (s *store) TxSaveTutorial(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, tutorialGroupEnumId string) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 저장 여부 확인.
	tutorialInfo, err := txGameQueries.GetUserTutorialInfoByAccountUserId(ctx, accountUserId)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if strings.Contains(tutorialInfo, tutorialGroupEnumId) {
		return newReward, helper.ErrorWithStack(err_pb.Code_UserErrAlreadySaveTutorial.String())
	}

	var tutorialGroupList []string
	err = json.Unmarshal([]byte(tutorialInfo), &tutorialGroupList)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 보상 지급.
	tutorialGroupStatic, err := s.StaticDataQueries.GetTutorialGroupByEnumId(ctx, tutorialGroupEnumId)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrTutorialGroupNotFound.String())
	}
	if tutorialGroupStatic.GroupReward != "" {
		err = s.GetDropRewardList(ctx, &newReward, []string{tutorialGroupStatic.GroupReward})
		if err != nil {
			return newReward, err
		}
		err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
		if err != nil {
			return newReward, err
		}
	}

	// 튜토리얼 저장.
	tutorialGroupList = append(tutorialGroupList, tutorialGroupEnumId)
	updateTutorialInfo, err := json.Marshal(tutorialGroupList)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}
	_, err = txGameQueries.UpdateUserTutorialInfoByAccountUserId(ctx, db_game.UpdateUserTutorialInfoByAccountUserIdParams{
		AccountUserID: accountUserId,
		TutorialInfo:  string(updateTutorialInfo),
	})
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}
	return newReward, nil
}
