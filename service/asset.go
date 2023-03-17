package service

import (
	"context"
	"database/sql"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/asset"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

type assetService struct {
	pb.AssetServiceServer
}

func NewAssetService() *assetService {
	return &assetService{}
}

func (s *assetService) GetAssets(ctx context.Context, _ *model_pb.Empty) (*pb.GetAssetsRes, error) {
	var assetsData []*model_pb.AssetData
	mySession := db_session.GetMySession(ctx)

	assets, err := db.Store().GameQueries[mySession.GameDb].GetAssetListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	for _, asset := range assets {
		assetsData = append(assetsData, &model_pb.AssetData{
			Id:      asset.ID,
			EnumId:  asset.EnumID,
			Type:    asset.Type,
			Balance: asset.Balance,
		})
	}

	return &pb.GetAssetsRes{
		Assets: assetsData,
	}, nil
}
