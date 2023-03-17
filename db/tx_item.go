package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	db_static_data "github.com/yeomc/game-server-grpc-go/db/sqlc/static_data"
	"github.com/yeomc/game-server-grpc-go/enum"
	"github.com/yeomc/game-server-grpc-go/helper"
)

func (s *store) TxUseItems(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId, targetId int32, useItems []*model_pb.UseItem) (err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 사용 아이템 개수 취합.
	itemsId := []int32{}
	itemIdUseCount := map[int32]int32{}
	for _, useItem := range useItems {
		itemsId = append(itemsId, useItem.Id)
		if _, ok := itemIdUseCount[useItem.Id]; ok {
			itemIdUseCount[useItem.Id] += useItem.Count
		} else {
			itemIdUseCount[useItem.Id] = useItem.Count
		}
	}

	// 수량 체크, 스태틱 데이터 정보 가져오기.
	items, err := txGameQueries.GetItemListById(ctx, itemsId)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	itemsEnumId := []string{}
	itemIdEnumId := map[int32]string{}
	for _, item := range items {
		if item.AccountUserID != accountUserId {
			return helper.ErrorWithStack(err_pb.Code_ItemErrNotFoundItem.String())
		}
		if item.Count < itemIdUseCount[item.ID] {
			return helper.ErrorWithStack(err_pb.Code_ItemErrNotEnoughItem.String())
		}

		itemsEnumId = append(itemsEnumId, item.EnumID)
		itemIdEnumId[item.ID] = item.EnumID
	}

	itemsStatic, err := s.StaticDataQueries.GetItemListByEnumId(ctx, itemsEnumId)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if len(itemsStatic) == 0 {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrItemNotFound.String())
	}
	enumIdItemStatic := map[string]*db_static_data.Item{}
	for i, itemStatic := range itemsStatic {
		enumIdItemStatic[itemStatic.EnumID] = &itemsStatic[i]
	}

	// 증가 경험치 취합, 사용 아이템 세팅.
	var costGold int32
	subReward := model_pb.SubReward{}
	itemTypeSum := map[enum.Item_Sub]int32{}
	for itemId, useCount := range itemIdUseCount {
		enumId := itemIdEnumId[itemId]
		itemStatic := enumIdItemStatic[enumId]
		itemType := enum.GetItem_Sub(itemStatic.CeItemSub)

		if itemType == enum.Item_Sub_CHARACTER_EXP {
			if _, ok := itemTypeSum[itemType]; ok {
				itemTypeSum[itemType] += itemStatic.Value * useCount
			} else {
				itemTypeSum[itemType] = itemStatic.Value * useCount
			}
			costGold += itemStatic.CostValue * useCount
		}

		s.SetSubReward(&subReward, 0, enum.Common_Type_ITEM, enumId, useCount)
	}

	if costGold > 0 {
		goldData, err := s.StaticDataQueries.GetAssetByAssetEnum(ctx, enum.Asset_GOLD.String())
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_StaticDataErrAssetNotFound.String())
		}
		s.SetSubReward(&subReward, 0, enum.Common_Type_ASSET, goldData.EnumID, costGold)
	}

	// 사용 아이템 차감.
	err = s.TxSubRewards(ctx, txGameQueries, gameDb, accountUserId, &subReward)
	if err != nil {
		return err
	}

	for itemType, amount := range itemTypeSum {
		if itemType == enum.Item_Sub_CHARACTER_EXP {
			// 캐릭터 경험치 테이블.
			characterGrowthsStatic, err := s.StaticDataQueries.GetCharacterGrowths(ctx)
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterGrowthNotFound.String())
			}
			characterExpTable := []int{}
			for _, characterGrowthStatic := range characterGrowthsStatic {
				characterExpTable = append(characterExpTable, int(characterGrowthStatic.Exp))
			}
			characterMaxLevel := len(characterGrowthsStatic)
			characterMaxExp := int32(helper.GetMaxExp(characterExpTable, characterMaxLevel))

			character, err := txGameQueries.GetCharacter(ctx, targetId)
			if err != nil {
				if err == sql.ErrNoRows {
					return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
				}
				return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}

			if character.AccountUserID != accountUserId {
				return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
			}

			if character.Exp+amount > characterMaxExp {
				amount = characterMaxExp - character.Exp
			}
			if amount > 0 {
				_, err = txGameQueries.UpdateCharacterExp(ctx, db_game.UpdateCharacterExpParams{
					ID:  character.ID,
					Exp: amount,
				})
				if err != nil {
					return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
				}
			}
		}
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return nil
}
