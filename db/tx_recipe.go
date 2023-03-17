package db

import (
	"context"
	"database/sql"
	"encoding/json"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) TxMakeRecipe(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, recipeEnumId string, count int32) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	if count <= 0 {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	recipeStatic, err := s.StaticDataQueries.GetRecipeByEnumId(ctx, recipeEnumId)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrRecipeNotFound.String())
	}

	// 재료 차감.
	materialsStatic, err := s.StaticDataQueries.GetMaterialsByEnumId(ctx, recipeStatic.RecipeMaterial)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterialsNotFound.String())
	}
	var typeArr []string
	var valueArr []string
	var amountArr []int32
	err = json.Unmarshal([]byte(materialsStatic.CeCommonTypeMaterial), &typeArr)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterials.String())
	}
	err = json.Unmarshal([]byte(materialsStatic.Material), &valueArr)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterials.String())
	}
	err = json.Unmarshal([]byte(materialsStatic.MaterialValue), &amountArr)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterials.String())
	}

	subReward := model_pb.SubReward{}
	for i, t := range typeArr {
		if t == "" {
			continue
		}

		subType := enum.GetCommon_Type(t)
		subValue := valueArr[i]
		subAmount := amountArr[i] * count
		s.SetSubReward(&subReward, 0, subType, subValue, subAmount)
	}
	err = s.TxSubRewards(ctx, txGameQueries, gameDb, accountUserId, &subReward)
	if err != nil {
		return newReward, err
	}

	// 제작 상품 지급.
	resultStaticData := RewardStaticData{}
	resultType := enum.GetCommon_Type(recipeStatic.CeCommonTypeResult)
	resultEnumId := recipeStatic.Result
	err = s.SetRewardStaticData(ctx, &resultStaticData, resultType, resultEnumId)
	if err != nil {
		return newReward, err
	}
	err = s.SetRewardList(ctx, resultStaticData, &newReward, resultType, resultEnumId, count)
	if err != nil {
		return newReward, err
	}
	err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
	if err != nil {
		return newReward, err
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return newReward, nil
}
