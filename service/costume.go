package service

import (
	"context"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/costume"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type costumeService struct {
	pb.CostumeServiceServer
}

func NewCostumeService() *costumeService {
	return &costumeService{}
}

func (s *costumeService) GetCostumes(ctx context.Context, _ *model_pb.Empty) (*pb.GetCostumesRes, error) {
	var costumesData []*model_pb.CostumeData
	mySession := db_session.GetMySession(ctx)

	costumes, err := db.Store().GameQueries[mySession.GameDb].GetCostumeListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, costume := range costumes {
		costumesData = append(costumesData, &model_pb.CostumeData{
			Id:              costume.ID,
			EnumId:          costume.EnumID,
			CharacterEnumId: costume.CharacterEnumID,
			State:           costume.State,
			CreatedAt:       timestamppb.New(costume.CreatedAt),
		})
	}

	return &pb.GetCostumesRes{
		Costumes: costumesData,
	}, nil
}

func (s *costumeService) GetCharacterCostumes(ctx context.Context, req *pb.GetCharacterCostumesReq) (*pb.GetCostumesRes, error) {
	var costumesData []*model_pb.CostumeData
	mySession := db_session.GetMySession(ctx)

	costumes, err := db.Store().GameQueries[mySession.GameDb].GetCostumeListByCharacterEnumId(ctx, db_game.GetCostumeListByCharacterEnumIdParams{
		AccountUserID:   mySession.AccountUserId,
		CharacterEnumID: req.CharacterEnumId,
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, costume := range costumes {
		costumesData = append(costumesData, &model_pb.CostumeData{
			Id:              costume.ID,
			EnumId:          costume.EnumID,
			CharacterEnumId: costume.CharacterEnumID,
			State:           costume.State,
			CreatedAt:       timestamppb.New(costume.CreatedAt),
		})
	}

	return &pb.GetCostumesRes{
		Costumes: costumesData,
	}, nil
}

func (s *costumeService) EquipCostume(ctx context.Context, req *pb.EquipCostumeReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	err := db.Store().TxEquipCostume(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.Id)
	if err != nil {
		return nil, err
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}
