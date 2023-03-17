package service

import (
	"context"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/story"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
)

type storyService struct {
	pb.StoryServiceServer
}

func NewStoryService() *storyService {
	return &storyService{}
}

func (s *storyService) ClearStory(ctx context.Context, req *pb.ClearStoryReq) (*pb.ClearStoryRes, error) {
	mySession := db_session.GetMySession(ctx)

	reward, storyIndex, err := db.Store().TxClearStory(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.StoryEnumId)
	if err != nil {
		return nil, err
	}

	return &pb.ClearStoryRes{
		Reward:     &reward,
		StoryIndex: storyIndex,
	}, nil
}
