package service

import (
	"context"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/character_broadcast"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type characterBroadcastService struct {
	pb.CharacterBroadcastServiceServer
}

func NewCharacterBroadcastService() *characterBroadcastService {
	return &characterBroadcastService{}
}

func (s *characterBroadcastService) GetOnAirCharacterBroadcasts(ctx context.Context, _ *model_pb.Empty) (*pb.GetOnAirCharacterBroadcastsRes, error) {
	var characterBroadcastsData []*model_pb.CharacterBroadcastData
	mySession := db_session.GetMySession(ctx)

	characterBroadcastList, broadcastResetAt, err := db.Store().TxGetOnAirCharacterBroadcasts(ctx, nil, mySession.GameDb, mySession.AccountUserId)
	if err != nil {
		return nil, err
	}
	for _, characterBroadcast := range characterBroadcastList {
		characterBroadcastsData = append(characterBroadcastsData, &model_pb.CharacterBroadcastData{
			Id:              characterBroadcast.ID,
			CharacterEnumId: characterBroadcast.CharacterEnumID,
			TimelineEnumId:  characterBroadcast.TimelineEnumID,
			Type:            characterBroadcast.Type,
			OnAir:           characterBroadcast.OnAir,
			Complete:        characterBroadcast.Complete,
			BroadcastedAt:   timestamppb.New(characterBroadcast.BroadcastedAt),
		})
	}

	return &pb.GetOnAirCharacterBroadcastsRes{
		CharacterBroadcasts: characterBroadcastsData,
		BroadcastResetAt:    timestamppb.New(broadcastResetAt),
	}, nil
}

func (s *characterBroadcastService) GetCompletedCharacterBroadcasts(ctx context.Context, req *pb.GetCompletedCharacterBroadcastsReq) (*pb.GetCompletedCharacterBroadcastsRes, error) {
	var characterBroadcastsData []*model_pb.CharacterBroadcastData
	mySession := db_session.GetMySession(ctx)

	characterBroadcastList, err := db.Store().GameQueries[mySession.GameDb].GetCompletedCharacterBroadcastList(ctx, db_game.GetCompletedCharacterBroadcastListParams{
		AccountUserID:   mySession.AccountUserId,
		CharacterEnumID: req.CharacterEnumId,
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, characterBroadcast := range characterBroadcastList {
		characterBroadcastsData = append(characterBroadcastsData, &model_pb.CharacterBroadcastData{
			Id:              characterBroadcast.ID,
			CharacterEnumId: characterBroadcast.CharacterEnumID,
			TimelineEnumId:  characterBroadcast.TimelineEnumID,
			Type:            characterBroadcast.Type,
			OnAir:           characterBroadcast.OnAir,
			Complete:        characterBroadcast.Complete,
			BroadcastedAt:   timestamppb.New(characterBroadcast.BroadcastedAt),
		})
	}

	return &pb.GetCompletedCharacterBroadcastsRes{
		CharacterBroadcasts: characterBroadcastsData,
	}, nil
}

func (s *characterBroadcastService) CompleteCharacterBroadcast(ctx context.Context, req *pb.CompleteCharacterBroadcastReq) (*pb.CompleteCharacterBroadcastRes, error) {
	mySession := db_session.GetMySession(ctx)

	reward, err := db.Store().TxCompleteCharacterBroadcast(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.CompleteCharacterBroadcastRes{
		Reward: &reward,
	}, nil
}
