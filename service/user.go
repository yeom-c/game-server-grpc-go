package service

import (
	"context"
	"database/sql"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeomc/game-server-grpc-go/helper"

	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/user"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userService struct {
	pb.UserServiceServer
}

func NewUserService() *userService {
	return &userService{}
}

func (s *userService) GetUser(ctx context.Context, _ *model_pb.Empty) (*pb.GetUserRes, error) {
	var userData *model_pb.UserData
	mySession := db_session.GetMySession(ctx)
	gameDb := mySession.GameDb

	user, err := db.Store().GameQueries[gameDb].GetUserByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_UserErrNotFoundUser.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	battleUser, err := db.Store().BattleQueries.GetUserByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	userData = &model_pb.UserData{
		Id:           user.ID,
		MatchPoint:   battleUser.MatchPoint,
		MatchWin:     battleUser.MatchWin,
		MatchLose:    battleUser.MatchLose,
		StoryIndex:   user.StoryIndex,
		TutorialInfo: user.TutorialInfo,
		CreatedAt:    timestamppb.New(user.CreatedAt),
	}

	return &pb.GetUserRes{
		User: userData,
	}, nil
}

func (s *userService) SaveTutorial(ctx context.Context, req *pb.SaveTutorialReq) (*pb.SaveTutorialRes, error) {
	mySession := db_session.GetMySession(ctx)

	reward, err := db.Store().TxSaveTutorial(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.EnumId)
	if err != nil {
		return nil, err
	}

	return &pb.SaveTutorialRes{
		Reward: &reward,
	}, nil
}
