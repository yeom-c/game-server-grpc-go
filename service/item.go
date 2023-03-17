package service

import (
	"context"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/item"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	"github.com/yeomc/game-server-grpc-go/helper"
)

type itemService struct {
	pb.ItemServiceServer
}

func NewItemService() *itemService {
	return &itemService{}
}

func (s *itemService) GetItems(ctx context.Context, _ *model_pb.Empty) (*pb.GetItemsRes, error) {
	var itemsData []*model_pb.ItemData
	mySession := db_session.GetMySession(ctx)

	items, err := db.Store().GameQueries[mySession.GameDb].GetItemListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, item := range items {
		itemsData = append(itemsData, &model_pb.ItemData{
			Id:     item.ID,
			EnumId: item.EnumID,
			Count:  item.Count,
		})
	}

	return &pb.GetItemsRes{
		Items: itemsData,
	}, nil
}

func (s *itemService) UseItem(ctx context.Context, req *pb.UseItemReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	err := db.Store().TxUseItems(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.TargetId, req.UseItems)
	if err != nil {
		return nil, err
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}
