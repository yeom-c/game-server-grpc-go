package service

import (
	"context"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/character_collection"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	"github.com/yeom-c/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type characterCollectionService struct {
	pb.CharacterCollectionServiceServer
}

func NewCharacterCollectionService() *characterCollectionService {
	return &characterCollectionService{}
}

func (s *characterCollectionService) GetCharacterCollections(ctx context.Context, _ *model_pb.Empty) (*pb.GetCharacterCollectionsRes, error) {
	var characterCollectionsData []*model_pb.CharacterCollectionData
	mySession := db_session.GetMySession(ctx)

	characterCollectionList, err := db.Store().GameQueries[mySession.GameDb].GetCharacterCollectionListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, characterCollection := range characterCollectionList {
		characterCollectionsData = append(characterCollectionsData, &model_pb.CharacterCollectionData{
			Id:              characterCollection.ID,
			CharacterEnumId: characterCollection.CharacterEnumID,
			AffectionExp:    characterCollection.AffectionExp,
			Count:           characterCollection.Count,
			CreatedAt:       timestamppb.New(characterCollection.CreatedAt),
		})
	}

	return &pb.GetCharacterCollectionsRes{
		CharacterCollections: characterCollectionsData,
	}, nil
}
