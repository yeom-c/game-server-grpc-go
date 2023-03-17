package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	"time"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) TxGetOnAirCharacterBroadcasts(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32) (broadcasts []db_game.CharacterBroadcast, broadcastResetAt time.Time, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 방송 목록 갱신.
	resetIntervalMinute := float64(20)
	now := time.Now().UTC()
	user, err := txGameQueries.GetUserByAccountUserId(ctx, accountUserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_UserErrNotFoundUser.String())
		}
		return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	broadcastResetAt = user.BroadcastResetAt
	if now.Sub(broadcastResetAt).Minutes() >= resetIntervalMinute {
		// 기존 방송 종료.
		_, err := txGameQueries.UpdateCharacterBroadcastOnAirByAccountUserId(ctx, db_game.UpdateCharacterBroadcastOnAirByAccountUserIdParams{
			AccountUserID: accountUserId,
			OnAir:         0,
		})
		if err != nil {
			return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}

		broadcastResetAt = now
		broadcastPoolData, err := s.GetAvailableCharacterBroadcastPoolData(ctx, txGameQueries, gameDb, accountUserId)
		if err != nil {
			return broadcasts, broadcastResetAt, err
		}

		// 이벤트 방송.
		maxEventCnt := 6
		for i := 0; i < maxEventCnt; i++ {
			if len(broadcastPoolData.event) > 0 {
				randIndex := helper.GetRandInt(0, len(broadcastPoolData.event)-1)
				for characterEnumId, timelineEnumIdList := range broadcastPoolData.event {
					if randIndex == 0 {
						if len(timelineEnumIdList) > 0 {
							timelineEnumId := timelineEnumIdList[helper.GetRandInt(0, len(timelineEnumIdList)-1)]
							_, err = txGameQueries.UpsertCharacterBroadcast(ctx, db_game.UpsertCharacterBroadcastParams{
								AccountUserID:   accountUserId,
								CharacterEnumID: characterEnumId,
								TimelineEnumID:  timelineEnumId,
								Type:            int32(enum.Oracle_Broadcast_Type_EVENT),
								OnAir:           1,
								Complete:        0,
							})
							if err != nil {
								return broadcasts, broadcastResetAt, err
							}

							// 뽑힌 캐릭터 다른 방송에서 제외.
							delete(broadcastPoolData.regular, characterEnumId)
							delete(broadcastPoolData.irregular, characterEnumId)
							delete(broadcastPoolData.noGet, characterEnumId)
						}
						delete(broadcastPoolData.event, characterEnumId)
						break
					}
					randIndex--
				}
			} else {
				break
			}
		}

		// 보유 캐릭터 방송.
		maxRegularCnt := 3
		if len(broadcastPoolData.regular) >= 20 {
			maxRegularCnt = 5
		} else if len(broadcastPoolData.regular) >= 10 {
			maxRegularCnt = 4
		}
		for i := 0; i < maxRegularCnt; i++ {
			if len(broadcastPoolData.regular) > 0 {
				randIndex := helper.GetRandInt(0, len(broadcastPoolData.regular)-1)
				for characterEnumId, timelineEnumIdList := range broadcastPoolData.regular {
					if randIndex == 0 {
						if len(timelineEnumIdList) > 0 {
							// regular timeline 은 순서대로 방송이라서 첫번째를 뽑음.
							timelineEnumId := timelineEnumIdList[0]
							_, err = txGameQueries.UpsertCharacterBroadcast(ctx, db_game.UpsertCharacterBroadcastParams{
								AccountUserID:   accountUserId,
								CharacterEnumID: characterEnumId,
								TimelineEnumID:  timelineEnumId,
								Type:            int32(enum.Oracle_Broadcast_Type_REGULAR),
								OnAir:           1,
								Complete:        0,
							})
							if err != nil {
								return broadcasts, broadcastResetAt, err
							}

							// 뽑힌 캐릭터 다른 방송에서 제외.
							delete(broadcastPoolData.irregular, characterEnumId)
							delete(broadcastPoolData.noGet, characterEnumId)
						}
						delete(broadcastPoolData.regular, characterEnumId)
						break
					}
					randIndex--
				}
			} else if len(broadcastPoolData.irregular) > 0 {
				randIndex := helper.GetRandInt(0, len(broadcastPoolData.irregular)-1)
				for characterEnumId, timelineEnumIdList := range broadcastPoolData.irregular {
					if randIndex == 0 {
						if len(timelineEnumIdList) > 0 {
							timelineEnumId := timelineEnumIdList[helper.GetRandInt(0, len(timelineEnumIdList)-1)]
							_, err = txGameQueries.UpsertCharacterBroadcast(ctx, db_game.UpsertCharacterBroadcastParams{
								AccountUserID:   accountUserId,
								CharacterEnumID: characterEnumId,
								TimelineEnumID:  timelineEnumId,
								Type:            int32(enum.Oracle_Broadcast_Type_IRREGULAR),
								OnAir:           1,
								Complete:        0,
							})
							if err != nil {
								return broadcasts, broadcastResetAt, err
							}

							// 뽑힌 캐릭터 다른 방송에서 제외.
							delete(broadcastPoolData.noGet, characterEnumId)
						}
						delete(broadcastPoolData.irregular, characterEnumId)
						break
					}
					randIndex--
				}
			} else {
				break
			}
		}

		// 미보유 캐릭터 방송.
		maxNoGetCnt := 1
		if len(broadcastPoolData.noGet) >= 10 {
			maxNoGetCnt = 3
		} else if len(broadcastPoolData.noGet) >= 5 {
			maxNoGetCnt = 2
		}
		for i := 0; i < maxNoGetCnt; i++ {
			if len(broadcastPoolData.noGet) > 0 {
				randIndex := helper.GetRandInt(0, len(broadcastPoolData.noGet)-1)
				for characterEnumId, timelineEnumIdList := range broadcastPoolData.noGet {
					if randIndex == 0 {
						if len(timelineEnumIdList) > 0 {
							timelineEnumId := timelineEnumIdList[helper.GetRandInt(0, len(timelineEnumIdList)-1)]
							_, err = txGameQueries.UpsertCharacterBroadcast(ctx, db_game.UpsertCharacterBroadcastParams{
								AccountUserID:   accountUserId,
								CharacterEnumID: characterEnumId,
								TimelineEnumID:  timelineEnumId,
								Type:            int32(enum.Oracle_Broadcast_Type_NOGET),
								OnAir:           1,
								Complete:        0,
							})
							if err != nil {
								return broadcasts, broadcastResetAt, err
							}
						}
						delete(broadcastPoolData.noGet, characterEnumId)
						break
					}
					randIndex--
				}
			} else {
				break
			}
		}

		// 갱신 시간 업데이트.
		_, err = txGameQueries.UpdateUserBroadcastResetAtByAccountUserId(ctx, db_game.UpdateUserBroadcastResetAtByAccountUserIdParams{
			AccountUserID:    accountUserId,
			BroadcastResetAt: broadcastResetAt,
		})
		if err != nil {
			return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	broadcasts, err = txGameQueries.GetOnAirCharacterBroadcastList(ctx, accountUserId)
	if err != nil {
		return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return broadcasts, broadcastResetAt, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return broadcasts, broadcastResetAt, nil
}

func (s *store) TxCompleteCharacterBroadcast(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId, characterBroadcastId int32) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 방송 정보.
	broadcast, err := txGameQueries.GetOnAirCharacterBroadcast(ctx, characterBroadcastId)
	if err != nil {
		if err == sql.ErrNoRows {
			return newReward, helper.ErrorWithStack(err_pb.Code_CharacterBroadcastErrNotFoundCharacterBroadcast.String())
		}
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if broadcast.AccountUserID != accountUserId {
		return newReward, helper.ErrorWithStack(err_pb.Code_CharacterBroadcastErrNotFoundCharacterBroadcast.String())
	}

	// 방송중 확인
	if broadcast.OnAir != 1 {
		return newReward, helper.ErrorWithStack(err_pb.Code_CharacterBroadcastErrNotOnAir.String())
	}
	// 미완료 확인
	if broadcast.Complete != 0 {
		return newReward, helper.ErrorWithStack(err_pb.Code_CharacterBroadcastErrAlreadyCompleted.String())
	}

	// 보상 지급.
	vstoryTimelineStatic, err := s.StaticDataQueries.GetVstoryTimelineByEnumId(ctx, broadcast.TimelineEnumID)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrVstoryTimelinePoolNotFound.String())
	}
	reward := vstoryTimelineStatic.StoryReward
	if reward != "" {
		err = s.GetDropRewardList(ctx, &newReward, []string{reward})
		if err != nil {
			return newReward, err
		}
		err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
		if err != nil {
			return newReward, err
		}
	}

	// 완료 업데이트.
	_, err = txGameQueries.UpdateCharacterBroadcastComplete(ctx, db_game.UpdateCharacterBroadcastCompleteParams{
		ID:       characterBroadcastId,
		Complete: 1,
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
