package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/cheat"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	"time"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	"github.com/yeomc/game-server-grpc-go/enum"
	"github.com/yeomc/game-server-grpc-go/helper"
	"github.com/yeomc/game-server-grpc-go/model"
)

type cheatService struct {
	pb.CheatServiceServer
}

func NewCheatService() *cheatService {
	return &cheatService{}
}

func (s *cheatService) CreateMails(ctx context.Context, _ *model_pb.Empty) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	startDataList, err := db.Store().StaticDataQueries.GetStarts(ctx)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_StaticDataErrStartNotFound.String())
	}

	mailList := []db_game.CreateMailParams{}
	for i, startData := range startDataList {
		attachment := []model.MailAttachment{
			{
				RewardType:  int32(enum.GetCommon_Type(startData.CeCommonType)),
				RewardValue: startData.RewardID,
				RewardCount: startData.Value,
			},
		}
		attachmentB, err := json.Marshal(attachment)
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}

		// 메세지 메일.
		mailList = append(mailList, db_game.CreateMailParams{
			AccountUserID: mySession.AccountUserId,
			Sender:        "운영자",
			Type:          0,
			Status:        int32(enum.Mail_Status_NEW),
			DeleteAll:     1,
			Attachment:    sql.NullString{},
			Title:         fmt.Sprintf("일반 메세지 %d", i),
			Message:       sql.NullString{Valid: true, String: fmt.Sprintf("일반 메세지 %d", i)},
			ExpiredAt:     sql.NullTime{},
		})

		// 아이템 메일 만료.
		mailList = append(mailList, db_game.CreateMailParams{
			AccountUserID: mySession.AccountUserId,
			Sender:        "운영자",
			Type:          0,
			Status:        int32(enum.Mail_Status_NEW),
			DeleteAll:     1,
			Attachment:    sql.NullString{Valid: true, String: string(attachmentB)},
			Title:         fmt.Sprintf("아이템 메일 만료(자동삭제) %d", i),
			Message:       sql.NullString{Valid: true, String: fmt.Sprintf("아이템 메일 만료(자동삭제) %d", i)},
			ExpiredAt:     sql.NullTime{Valid: true, Time: time.Now().AddDate(0, 0, -7)},
		})

		// 아이템 메일 유기한.
		mailList = append(mailList, db_game.CreateMailParams{
			AccountUserID: mySession.AccountUserId,
			Sender:        "운영자",
			Type:          0,
			Status:        int32(enum.Mail_Status_NEW),
			DeleteAll:     1,
			Attachment:    sql.NullString{Valid: true, String: string(attachmentB)},
			Title:         fmt.Sprintf("아이템 메일 유기한 %d", i),
			Message:       sql.NullString{Valid: true, String: fmt.Sprintf("아이템 메일 유기한 %d", i)},
			ExpiredAt:     sql.NullTime{Valid: true, Time: time.Now().AddDate(0, 0, 7)},
		})

		// 아이템 메일 무기한.
		mailList = append(mailList, db_game.CreateMailParams{
			AccountUserID: mySession.AccountUserId,
			Sender:        "운영자",
			Type:          0,
			Status:        int32(enum.Mail_Status_NEW),
			DeleteAll:     1,
			Attachment:    sql.NullString{Valid: true, String: string(attachmentB)},
			Title:         fmt.Sprintf("아이템 메일 무기한 %d", i),
			Message:       sql.NullString{Valid: true, String: fmt.Sprintf("아이템 메일 무기한 %d", i)},
			ExpiredAt:     sql.NullTime{},
		})
	}

	_, err = db.Store().GameQueries[mySession.GameDb].CreateMails(ctx, mailList)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &model_pb.Result{Result: 1}, nil
}

func (s *cheatService) CreateAsset(ctx context.Context, req *pb.CreateAssetReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	assetStatic, err := db.Store().StaticDataQueries.GetAssetByEnumId(ctx, req.EnumId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_StaticDataErrAssetNotFound.String())
	}

	db.Store().GameQueries[mySession.GameDb].UpsertAsset(ctx, db_game.UpsertAssetParams{
		AccountUserID: mySession.AccountUserId,
		EnumID:        assetStatic.EnumID,
		Type:          int32(enum.GetAsset(assetStatic.CeAsset)),
		Amount:        req.Amount,
	})

	return &model_pb.Result{Result: 1}, nil
}
