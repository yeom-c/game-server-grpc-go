package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	db_static_data "github.com/yeomc/game-server-grpc-go/db/sqlc/static_data"
	"github.com/yeomc/game-server-grpc-go/enum"
	"github.com/yeomc/game-server-grpc-go/helper"
)

func (s *store) TxCreateRewards(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, reward *model_pb.Reward) error {
	var tx *sql.Tx
	var err error
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	var newMergeReward model_pb.Reward
	// asset reward
	for _, rAsset := range reward.Assets {
		err = s.CreateRewardAsset(ctx, txGameQueries, gameDb, accountUserId, rAsset)
		if err != nil {
			return err
		}
	}

	// character reward
	for _, rCharacter := range reward.Characters {
		newReward, err := s.CreateRewardCharacter(ctx, txGameQueries, gameDb, accountUserId, rCharacter)
		if err != nil {
			return err
		}
		newMergeReward = helper.MergeReward(newMergeReward, newReward)
	}

	// item reward
	for _, rItem := range reward.Items {
		err = s.CreateRewardItem(ctx, txGameQueries, gameDb, accountUserId, rItem)
		if err != nil {
			return err
		}
	}

	// fate card reward
	for _, rFateCard := range reward.FateCards {
		err = s.CreateRewardFateCard(ctx, txGameQueries, gameDb, accountUserId, rFateCard)
		if err != nil {
			return err
		}
	}

	// merge reward
	*reward = helper.MergeReward(*reward, newMergeReward)

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}
	return nil
}

func (s *store) TxSubRewards(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, subReward *model_pb.SubReward) error {
	var tx *sql.Tx
	var err error
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	if len(subReward.Assets) > 0 {
		var assetsEnumId []string
		for _, subAsset := range subReward.Assets {
			assetsEnumId = append(assetsEnumId, subAsset.EnumId)
		}

		assets, err := txGameQueries.GetAssetListByEnumId(ctx, accountUserId, assetsEnumId)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_AssetErrNotFoundAsset.String())
		}

		curAssets := map[string]db_game.Asset{}
		for _, asset := range assets {
			curAssets[asset.EnumID] = asset
		}

		for _, subAsset := range subReward.Assets {
			if _, ok := curAssets[subAsset.EnumId]; !ok {
				return helper.ErrorWithStack(err_pb.Code_AssetErrNotFoundAsset.String())
			}
			if curAssets[subAsset.EnumId].Balance < int64(subAsset.Count) {
				return helper.ErrorWithStack(err_pb.Code_AssetErrNotEnoughAsset.String())
			}
			subAsset.Id = curAssets[subAsset.EnumId].ID

			_, err := txGameQueries.AddAssetBalance(ctx, db_game.AddAssetBalanceParams{
				ID:     subAsset.Id,
				Amount: -int64(subAsset.Count),
			})
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}
		}
	}

	if len(subReward.Characters) > 0 {
		subCharactersId := []int32{}
		for _, subCharacter := range subReward.Characters {
			if subCharacter.Id == 0 {
				// TODO: 임의 character 재료 선정 처리 필요.
				return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
				//character, err := txGameQueries.GetCharacterByEnumId(ctx, db_game.GetCharacterByEnumIdParams{
				//	AccountUserID: accountUserId,
				//	EnumID:        subCharacter.EnumId,
				//})
				//if err != nil {
				//	return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
				//}
				//subCharacter.Id = character.ID
			}

			subCharactersId = append(subCharactersId, subCharacter.Id)
		}

		subCharacters, err := txGameQueries.GetCharacterListByIdAndAccountUserId(ctx, subCharactersId, accountUserId)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		if len(subCharacters) == 0 {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		} else if len(subCharacters) != len(subCharactersId) {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}

		_, err = txGameQueries.DeleteCharacters(ctx, subCharactersId)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	if len(subReward.Items) > 0 {
		var itemsEnumId []string
		for _, subItem := range subReward.Items {
			itemsEnumId = append(itemsEnumId, subItem.EnumId)
		}

		items, err := txGameQueries.GetItemsByEnumId(ctx, accountUserId, itemsEnumId)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_ItemErrNotFoundItem.String())
		}

		curItems := map[string]db_game.Item{}
		for _, item := range items {
			curItems[item.EnumID] = item
		}

		for _, subItem := range subReward.Items {
			if _, ok := curItems[subItem.EnumId]; !ok {
				return helper.ErrorWithStack(err_pb.Code_ItemErrNotFoundItem.String())
			}
			if curItems[subItem.EnumId].Count < subItem.Count {
				return helper.ErrorWithStack(err_pb.Code_ItemErrNotEnoughItem.String())
			}
			subItem.Id = curItems[subItem.EnumId].ID

			_, err := txGameQueries.AddItemCount(ctx, db_game.AddItemCountParams{
				ID:     subItem.Id,
				Amount: -subItem.Count,
			})
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
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

func (s *store) TxCreateRewardGacha(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, goodsCategoryId int32, gachaEnumId string, count int32, goodsInfo string) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 등급 보정 스택, 픽업 스택 정보 가져오기.
	var gradeStack, pickupStack int32
	userShopInfoMap := map[string]map[string]int32{}
	userShopInfo, err := txGameQueries.GetUserShopInfoByAccountUserId(ctx, accountUserId)
	if err != nil {
		if err != sql.ErrNoRows {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}
	if userShopInfo.Valid {
		err = json.Unmarshal([]byte(userShopInfo.String), &userShopInfoMap)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}
		categoryInfoMap := userShopInfoMap[fmt.Sprint(goodsCategoryId)]
		if categoryInfoMap != nil {
			if _, ok := categoryInfoMap["grade_stack"]; ok {
				gradeStack = categoryInfoMap["grade_stack"]
			}
			if _, ok := categoryInfoMap["pickup_stack"]; ok {
				pickupStack = categoryInfoMap["pickup_stack"]
			}
		}
	}

	// 가챠 스택 데이터.
	gachaStackModelEnumId := ""
	if goodsInfo != "" {
		gachaInfo := map[string]interface{}{}
		err = json.Unmarshal([]byte(goodsInfo), &gachaInfo)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}
		gachaStackModelEnumId = gachaInfo["gacha_stack_model"].(string)
	}
	gachaStackModelData, err := s.StaticDataQueries.GetGachaStackModelByEnumId(ctx, gachaStackModelEnumId)
	if err != nil {
		if err == sql.ErrNoRows {
			return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrGachaStackModelNotFound.String())
		}
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 가챠 데이터.
	gachaData, err := s.StaticDataQueries.GetGachaByEnumId(ctx, gachaEnumId)
	if err != nil {
		if err == sql.ErrNoRows {
			return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrGachaNotFound.String())
		}
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 가챠 클래스 데이터(등급별 확률).
	gachaClassData, err := s.StaticDataQueries.GetGachaClassByEnumId(ctx, gachaData.GachaClass)
	if err != nil {
		if err == sql.ErrNoRows {
			return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrGachaClassNotFound.String())
		}
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	var groupProb []float64
	var poolGroup []string
	err = json.Unmarshal([]byte(gachaClassData.Prob), &groupProb)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}
	err = json.Unmarshal([]byte(gachaClassData.PoolGroup), &poolGroup)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}
	// 최고등급 확률 보정을 위해 해당 확률만 따로 빼냄, 최고 등급 확률로 먼저 선택하고 실패시 나머지 등급들로 선택.
	groupAProb := groupProb[enum.Gacha_Pool_Grade_A]
	groupProb[enum.Gacha_Pool_Grade_A] = 0

	// 가챠풀 그룹 데이터.
	gachaPoolGroupList, err := s.StaticDataQueries.GetGachaPoolGroupListByEnumId(ctx, poolGroup)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if len(gachaPoolGroupList) == 0 {
		return newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrGachaPoolGroupNotFound.String())
	}

	enumIdGachaPoolGroup := map[string]db_static_data.GachaPoolGroup{}
	for _, gachaPoolGroup := range gachaPoolGroupList {
		enumIdGachaPoolGroup[gachaPoolGroup.EnumID] = gachaPoolGroup
	}

	// 가챠풀 그룹별 캐릭터 목록.
	gachaDropAList, err := s.StaticDataQueries.GetGachaDropAList(ctx)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	gachaDropBList, err := s.StaticDataQueries.GetGachaDropBList(ctx)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	gachaDropCList, err := s.StaticDataQueries.GetGachaDropCList(ctx)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	gachaDropDList, err := s.StaticDataQueries.GetGachaDropDList(ctx)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 가챠 로직.
	needChangeGradeB := false
	if count == 11 {
		needChangeGradeB = true
	}
	winCharacterList := []string{}
	for i := 0; i < int(count); i++ {
		// 최고 등급 확률 보정.
		var addCorrectionRate float64
		if gradeStack >= gachaStackModelData.Stack100 {
			addCorrectionRate = 1
		} else if gradeStack >= gachaStackModelData.StackCorrectionStart {
			addCorrectionRate = float64(gradeStack-gachaStackModelData.StackCorrectionStart+1) * gachaStackModelData.AProbCorrection
		}

		// 등급 결정.
		gradeARate := groupAProb + addCorrectionRate
		winIndex := 0
		if helper.IsWinRate(gradeARate) { // A 등급 확인.
			winIndex = int(enum.Gacha_Pool_Grade_A)
			gradeStack = 0
		} else { // 나머지 등급 확인.
			winIndex = helper.GetRandIndexByFloatWeights(groupProb)
			gradeStack += 1
		}

		poolGrade := enum.Gacha_Pool_Grade(winIndex)
		if poolGrade >= enum.Gacha_Pool_Grade_B {
			needChangeGradeB = false
		}
		poolGroupEnumId := poolGroup[winIndex]
		gachaPoolGroup := enumIdGachaPoolGroup[poolGroupEnumId]

		// 가챠풀 결정.
		var isPickup bool
		var poolIdList []string
		var poolProb []float64
		err = json.Unmarshal([]byte(gachaPoolGroup.PoolID), &poolIdList)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}
		err = json.Unmarshal([]byte(gachaPoolGroup.PoolIDRate), &poolProb)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}

		winPoolIndex := helper.GetRandIndexByFloatWeights(poolProb)
		// A 등급 픽업풀 확률 보정.
		if poolGrade == enum.Gacha_Pool_Grade_A {
			if pickupStack >= int32(enum.ServerEnum["max_gacha_pickup_stack"].(int)) {
				winPoolIndex = 0
			}

			if winPoolIndex == 0 {
				pickupStack = 0
			} else {
				pickupStack += 1
			}
		}

		if winPoolIndex == 0 {
			isPickup = true
		}
		poolId := poolIdList[winPoolIndex]
		pickupPoolId := poolIdList[0]

		// 캐릭터 결정.
		winCharacterEnumId, err := s.GetGachaDropCharacter(ctx, poolGrade, isPickup, poolId, pickupPoolId, gachaDropDList, gachaDropCList, gachaDropBList, gachaDropAList)
		if err != nil {
			return newReward, err
		}
		winCharacterList = append(winCharacterList, winCharacterEnumId)
	}

	// 11회차 C, D 등급만 있으면 마지막 뽑힌 캐릭터 B 등급으로 지급.
	if needChangeGradeB {
		poolGrade := enum.Gacha_Pool_Grade_B
		poolGroupEnumId := poolGroup[poolGrade]
		gachaPoolGroup := enumIdGachaPoolGroup[poolGroupEnumId]

		// 가챠풀 결정.
		var isPickup bool
		var poolIdList []string
		var poolProb []float64
		err = json.Unmarshal([]byte(gachaPoolGroup.PoolID), &poolIdList)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}
		err = json.Unmarshal([]byte(gachaPoolGroup.PoolIDRate), &poolProb)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
		}

		winPoolIndex := helper.GetRandIndexByFloatWeights(poolProb)
		if winPoolIndex == 0 {
			isPickup = true
		}
		poolId := poolIdList[winPoolIndex]
		pickupPoolId := poolIdList[0]

		// 캐릭터 결정.
		gradeBCharacterEnumId, err := s.GetGachaDropCharacter(ctx, poolGrade, isPickup, poolId, pickupPoolId, gachaDropDList, gachaDropCList, gachaDropBList, gachaDropAList)
		if err != nil {
			return newReward, err
		}

		winCharacterList[10] = gradeBCharacterEnumId
	}

	// 캐릭터 생성.
	rewardStaticData := RewardStaticData{}
	rewardData := DropRewardData{}
	gachaLogParams := []db_game.CreateGachaLogParams{}

	for _, characterEnumId := range winCharacterList {
		rewardType := enum.Common_Type_CHARACTER
		rewardEnumId := characterEnumId
		rewardCount := 1
		err = s.SetRewardStaticData(ctx, &rewardStaticData, rewardType, rewardEnumId)
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}

		rewardData.TypeArr = append(rewardData.TypeArr, rewardType.String())
		rewardData.DropArr = append(rewardData.DropArr, rewardEnumId)
		rewardData.ValueArr = append(rewardData.ValueArr, int32(rewardCount))

		gachaLogParams = append(gachaLogParams, db_game.CreateGachaLogParams{
			AccountUserID:   accountUserId,
			EnumID:          gachaEnumId,
			CharacterEnumID: rewardEnumId,
		})
	}

	for i, t := range rewardData.TypeArr {
		rewardType := enum.GetCommon_Type(t)
		rewardEnumId := rewardData.DropArr[i]
		rewardValue := rewardData.ValueArr[i]
		err := s.SetRewardList(ctx, rewardStaticData, &newReward, rewardType, rewardEnumId, rewardValue)
		if err != nil {
			return newReward, err
		}
	}
	err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
	if err != nil {
		return newReward, err
	}

	// 가챠 기록 저장.
	_, err = txGameQueries.CreateGachaLogs(ctx, gachaLogParams)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 가챠 보정 스택, 픽업 스택 정보 저장.
	if _, ok := userShopInfoMap[fmt.Sprint(goodsCategoryId)]; !ok {
		userShopInfoMap[fmt.Sprint(goodsCategoryId)] = map[string]int32{}
	}
	userShopInfoMap[fmt.Sprint(goodsCategoryId)]["grade_stack"] = gradeStack
	userShopInfoMap[fmt.Sprint(goodsCategoryId)]["pickup_stack"] = pickupStack
	userShopInfoByte, err := json.Marshal(userShopInfoMap)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}
	if len(userShopInfoByte) > 0 {
		userShopInfo.String = string(userShopInfoByte)
		userShopInfo.Valid = true
		_, err = txGameQueries.UpdateUserShopInfoByAccountUserId(ctx, db_game.UpdateUserShopInfoByAccountUserIdParams{
			AccountUserID: accountUserId,
			ShopInfo:      userShopInfo,
		})
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_ServerErr.String())
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
