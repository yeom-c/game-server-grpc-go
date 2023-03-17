package service

import (
	"context"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/recipe"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
)

type recipeService struct {
	pb.RecipeServiceServer
}

func NewRecipeService() *recipeService {
	return &recipeService{}
}

func (s *recipeService) MakeRecipe(ctx context.Context, req *pb.MakeRecipeReq) (*pb.MakeRecipeRes, error) {
	mySession := db_session.GetMySession(ctx)

	reward, err := db.Store().TxMakeRecipe(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.EnumId, req.Count)
	if err != nil {
		return nil, err
	}

	return &pb.MakeRecipeRes{
		Reward: &reward,
	}, nil
}
