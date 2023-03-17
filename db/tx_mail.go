package db

import (
	"context"
	"database/sql"
	"encoding/json"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	"time"

	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
	"github.com/yeom-c/game-server-grpc-go/model"

	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
)

func (s *store) TxConfirmMails(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, mailList []db_game.Mail) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 첨부 상품 취합.
	rewardStaticData := RewardStaticData{}
	rewardData := DropRewardData{}
	mailIdList := []int32{}
	for _, mail := range mailList {
		// 이미 수령 or 만료 메일 제외.
		if mail.Status == int32(enum.Mail_Status_CONFIRM) {
			continue
		} else if mail.ExpiredAt.Valid {
			if mail.ExpiredAt.Time.Before(time.Now()) {
				mailIdList = append(mailIdList, mail.ID)
				continue
			}
		}

		if mail.Attachment.Valid {
			var attachments []model.MailAttachment
			err := json.Unmarshal([]byte(mail.Attachment.String), &attachments)
			if err != nil {
				return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
			}

			for _, attachment := range attachments {
				rewardType := enum.Common_Type(attachment.RewardType)
				rewardValue := attachment.RewardValue
				rewardCount := attachment.RewardCount
				if rewardValue == "" || rewardCount == 0 {
					continue
				}

				err = s.SetRewardStaticData(ctx, &rewardStaticData, rewardType, rewardValue)
				if err != nil {
					return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
				}

				rewardData.TypeArr = append(rewardData.TypeArr, rewardType.String())
				rewardData.DropArr = append(rewardData.DropArr, rewardValue)
				rewardData.ValueArr = append(rewardData.ValueArr, rewardCount)
				mailIdList = append(mailIdList, mail.ID)
			}
		} else {
			mailIdList = append(mailIdList, mail.ID)
		}
	}

	for i, t := range rewardData.TypeArr {
		rewardType := enum.GetCommon_Type(t)
		rewardValue := rewardData.DropArr[i]
		rewardCount := rewardData.ValueArr[i]
		err := s.SetRewardList(ctx, rewardStaticData, &newReward, rewardType, rewardValue, rewardCount)
		if err != nil {
			return newReward, err
		}
	}

	// 첨부 상품 지급.
	err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
	if err != nil {
		return newReward, err
	}

	// 메일 수령 상태 업데이트.
	if len(mailIdList) > 0 {
		_, err = txGameQueries.UpdateMailsStatus(ctx, int32(enum.Mail_Status_CONFIRM), mailIdList)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return newReward, nil
}
