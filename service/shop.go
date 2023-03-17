package service

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/shop"
	"time"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	db_common "github.com/yeomc/game-server-grpc-go/db/sqlc/common"
	"github.com/yeomc/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type shopService struct {
	pb.ShopServiceServer
}

func NewShopService() *shopService {
	return &shopService{}
}

func (s *shopService) GetShop(ctx context.Context, req *pb.GetShopReq) (*pb.GetShopRes, error) {
	shop, err := db.Store().CommonQueries.GetShopByTypeAndVisible(ctx, db_common.GetShopByTypeAndVisibleParams{
		ShopVisible: 1,
		ShopType:    req.ShopType,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_ShopErrNotFoundShop.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	var startAt, endAt *timestamppb.Timestamp
	if shop.StartAt.Valid {
		startAt = timestamppb.New(shop.StartAt.Time)
	}
	if shop.EndAt.Valid {
		endAt = timestamppb.New(shop.EndAt.Time)
	}

	categoryList := []*model_pb.ShopCategoryData{}
	shopCategoryList, err := db.Store().CommonQueries.GetShopCategoryListByShopId(ctx, shop.ID)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, shopCategory := range shopCategoryList {
		categoryList = append(categoryList, &model_pb.ShopCategoryData{
			Id:     shopCategory.ID,
			ShopId: shopCategory.ShopID,
			Name:   shopCategory.Name,
			Order:  shopCategory.Order,
		})
	}

	return &pb.GetShopRes{
		Shop: &model_pb.ShopData{
			Id:      shop.ID,
			Type:    shop.Type,
			Visible: shop.Visible,
			Name:    shop.Name,
			Desc:    shop.Desc,
			StartAt: startAt,
			EndAt:   endAt,
		},
		CategoryList: categoryList,
	}, nil
}

func (s *shopService) GetShopGoods(ctx context.Context, req *pb.GetShopGoodsReq) (*pb.GetShopGoodsRes, error) {
	mySession := db_session.GetMySession(ctx)

	userShopInfo, err := db.Store().GameQueries[mySession.GameDb].GetUserShopInfoByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	goodsList := []*model_pb.ShopGoodsData{}
	shopGoodsList, err := db.Store().CommonQueries.GetShopGoodsListByShopCategoryIdList(ctx, req.CategoryIdList)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	now := time.Now()
	for _, shopGoods := range shopGoodsList {
		// 판매 상태 체크.
		if shopGoods.Visible == 0 {
			continue
		}

		// 판매 시간 체크.
		var startAt, endAt *timestamppb.Timestamp
		if shopGoods.StartAt.Valid {
			if now.Before(shopGoods.StartAt.Time) {
				continue
			}
			startAt = timestamppb.New(shopGoods.StartAt.Time)
		}
		if shopGoods.EndAt.Valid {
			if now.After(shopGoods.EndAt.Time) {
				continue
			}
			endAt = timestamppb.New(shopGoods.EndAt.Time)
		}
		goodsList = append(goodsList, &model_pb.ShopGoodsData{
			Id:             shopGoods.ID,
			ShopCategoryId: shopGoods.ShopCategoryID,
			Type:           shopGoods.Type,
			EnumId:         shopGoods.EnumID,
			Info:           shopGoods.Info.String,
			Name:           shopGoods.Name,
			Desc:           shopGoods.Desc,
			CostType:       shopGoods.CostType,
			CostEnumId:     shopGoods.CostEnumID,
			Cost:           shopGoods.Cost,
			OriginalCost:   shopGoods.OriginalCost,
			Count:          shopGoods.Count,
			Visible:        shopGoods.Visible,
			StartAt:        startAt,
			EndAt:          endAt,
		})
	}

	return &pb.GetShopGoodsRes{
		UserShopInfo: userShopInfo.String,
		GoodsList:    goodsList,
	}, nil
}

func (s *shopService) BuyShopGoods(ctx context.Context, req *pb.BuyShopGoodsReq) (*pb.BuyShopGoodsRes, error) {
	mySession := db_session.GetMySession(ctx)
	now := time.Now()

	shopGoods, err := db.Store().CommonQueries.GetShopGoods(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_ShopGoodsErrNotFoundShopGoods.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if shopGoods.EnumID != req.EnumId {
		return nil, helper.ErrorWithStack(err_pb.Code_ShopGoodsErrNotFoundShopGoods.String())
	}

	// 활성 상태 체크.
	if shopGoods.Visible == 0 {
		return nil, helper.ErrorWithStack(err_pb.Code_ShopGoodsErrNotFoundShopGoods.String())
	}

	// 판매 시간 체크.
	if shopGoods.StartAt.Valid {
		if now.Before(shopGoods.StartAt.Time) {
			return nil, helper.ErrorWithStack(err_pb.Code_ShopGoodsErrNotFoundShopGoods.String())
		}
	}
	if shopGoods.EndAt.Valid {
		if now.After(shopGoods.EndAt.Time) {
			return nil, helper.ErrorWithStack(err_pb.Code_ShopGoodsErrNotFoundShopGoods.String())
		}
	}

	// 구매 처리.
	reward, err := db.Store().TxBuyGoods(ctx, nil, mySession.GameDb, mySession.AccountUserId, shopGoods)
	if err != nil {
		return nil, err
	}

	return &pb.BuyShopGoodsRes{
		Reward: &reward,
	}, nil
}
