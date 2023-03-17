package service

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/mail"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type mailService struct {
	pb.MailServiceServer
}

func NewMailService() *mailService {
	return &mailService{}
}

func (s *mailService) GetMails(ctx context.Context, _ *model_pb.Empty) (*pb.GetMailsRes, error) {
	var mailDataList []*model_pb.MailData
	mySession := db_session.GetMySession(ctx)

	// 만료 메일 삭제.
	_, err := db.Store().GameQueries[mySession.GameDb].DeleteAllExpiredMail(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 메일 목록.
	mailList, err := db.Store().GameQueries[mySession.GameDb].GetMailListByAccountUserId(ctx, db_game.GetMailListByAccountUserIdParams{
		AccountUserID: mySession.AccountUserId,
		Limit:         int32(enum.ServerEnum["max_confirm_mail_count"].(int)),
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, mail := range mailList {
		var expiredAt *timestamppb.Timestamp
		if mail.ExpiredAt.Valid {
			expiredAt = timestamppb.New(mail.ExpiredAt.Time)
		}

		mailDataList = append(mailDataList, &model_pb.MailData{
			Id:            mail.ID,
			AccountUserId: mail.AccountUserID,
			Sender:        mail.Sender,
			Type:          mail.Type,
			Status:        mail.Status,
			Attachment:    mail.Attachment.String,
			Title:         mail.Title,
			Message:       mail.Message.String,
			ExpiredAt:     expiredAt,
			CreatedAt:     timestamppb.New(mail.CreatedAt),
		})
	}

	return &pb.GetMailsRes{
		Mails: mailDataList,
	}, nil
}

func (s *mailService) ReadMail(ctx context.Context, req *pb.MailReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	// 메일 체크.
	mail, err := db.Store().GameQueries[mySession.GameDb].GetMail(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_MailErrNotFoundMail.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if mail.AccountUserID != mySession.AccountUserId {
		return nil, helper.ErrorWithStack(err_pb.Code_MailErrNotFoundMail.String())
	}

	// 읽음 상태 업데이트.
	_, err = db.Store().GameQueries[mySession.GameDb].UpdateMailStatus(ctx, db_game.UpdateMailStatusParams{ID: mail.ID, Status: int32(enum.Mail_Status_READ)})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}

func (s *mailService) ConfirmMail(ctx context.Context, req *pb.MailReq) (*pb.ConfirmMailRes, error) {
	mySession := db_session.GetMySession(ctx)

	// 메일 체크.
	mail, err := db.Store().GameQueries[mySession.GameDb].GetMail(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_MailErrNotFoundMail.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if mail.AccountUserID != mySession.AccountUserId {
		return nil, helper.ErrorWithStack(err_pb.Code_MailErrNotFoundMail.String())
	}

	// 메일 확인 트랜잭션 처리.
	reward, err := db.Store().TxConfirmMails(ctx, nil, mySession.GameDb, mySession.AccountUserId, []db_game.Mail{mail})
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmMailRes{
		Reward: &reward,
	}, nil
}

func (s *mailService) ConfirmAllMail(ctx context.Context, _ *model_pb.Empty) (*pb.ConfirmMailRes, error) {
	mySession := db_session.GetMySession(ctx)

	mailList, err := db.Store().GameQueries[mySession.GameDb].GetMailListByAccountUserId(ctx, db_game.GetMailListByAccountUserIdParams{
		AccountUserID: mySession.AccountUserId,
		Limit:         int32(enum.ServerEnum["max_confirm_mail_count"].(int)),
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	var reward model_pb.Reward
	if len(mailList) > 0 {
		// 메일 확인 트랜잭션 처리.
		reward, err = db.Store().TxConfirmMails(ctx, nil, mySession.GameDb, mySession.AccountUserId, mailList)
		if err != nil {
			return nil, err
		}
	}

	return &pb.ConfirmMailRes{
		Reward: &reward,
	}, nil
}

func (s *mailService) DeleteMail(ctx context.Context, req *pb.MailReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	_, error := db.Store().GameQueries[mySession.GameDb].DeleteConfirmMail(ctx, db_game.DeleteConfirmMailParams{AccountUserID: mySession.AccountUserId, ID: req.Id, Status: int32(enum.Mail_Status_CONFIRM)})
	if error != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}

func (s *mailService) DeleteAllMail(ctx context.Context, _ *model_pb.Empty) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	_, error := db.Store().GameQueries[mySession.GameDb].DeleteAllConfirmMail(ctx, db_game.DeleteAllConfirmMailParams{AccountUserID: mySession.AccountUserId, Status: int32(enum.Mail_Status_CONFIRM), DeleteAll: 1})
	if error != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}
