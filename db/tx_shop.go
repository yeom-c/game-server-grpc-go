package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_common "github.com/yeom-c/game-server-grpc-go/db/sqlc/common"
	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) TxBuyGoods(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, shopGoods db_common.ShopGoods) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 비용 차감.
	subReward := model_pb.SubReward{}
	s.SetSubReward(&subReward, 0, enum.Common_Type(shopGoods.CostType), shopGoods.CostEnumID, shopGoods.Cost)
	err = s.TxSubRewards(ctx, txGameQueries, gameDb, accountUserId, &subReward)
	if err != nil {
		return newReward, err
	}

	// 상품 지급.
	if enum.Common_Type_GACHA == enum.Common_Type(shopGoods.Type) {
		newReward, err = s.TxCreateRewardGacha(ctx, txGameQueries, gameDb, accountUserId, shopGoods.ShopCategoryID, shopGoods.EnumID, shopGoods.Count, shopGoods.Info.String)
		if err != nil {
			return newReward, err
		}
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return newReward, nil
}
