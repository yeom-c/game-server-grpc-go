package db

import (
	"context"
	"encoding/json"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	"strings"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	db_static_data "github.com/yeom-c/game-server-grpc-go/db/sqlc/static_data"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DropRewardData struct {
	TypeArr  []string
	DropArr  []string
	ValueArr []int32
	RateArr  []float32
}

type RewardStaticData struct {
	AssetsStatic     map[string]db_static_data.Asset
	CharactersStatic map[string]db_static_data.Character
	ItemsStatic      map[string]db_static_data.Item
	FateCardsStatic  map[string]db_static_data.FateCard
}

// TODO: drop table drop_next link 처리 필요.
func (s *store) GetDropRewardList(ctx context.Context, reward *model_pb.Reward, dropEnumIdList []string) error {
	// drop 데이터 확인.
	dropStaticList, err := s.StaticDataQueries.GetDropListByEnumId(ctx, dropEnumIdList)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrDropNotFound.String())
	}

	// drop 데이터 struct 로 변환.
	var typeArr []string
	var dropArr []string
	var valueArr []int32
	var rateArr []float32
	eachEnumIdDropRewardData := make(map[string]DropRewardData)
	for _, dropStatic := range dropStaticList {
		err = json.Unmarshal([]byte(dropStatic.CeCommonTypeDrop), &typeArr)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_StaticDataErrDrop.String())
		}
		err = json.Unmarshal([]byte(dropStatic.Drop), &dropArr)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_StaticDataErrDrop.String())
		}
		err = json.Unmarshal([]byte(dropStatic.Value), &valueArr)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_StaticDataErrDrop.String())
		}
		err = json.Unmarshal([]byte(dropStatic.Rate), &rateArr)
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_StaticDataErrDrop.String())
		}

		eachEnumIdDropRewardData[dropStatic.EnumID] = DropRewardData{
			TypeArr:  typeArr,
			DropArr:  dropArr,
			ValueArr: valueArr,
			RateArr:  rateArr,
		}
	}

	// drop item 데이터 취합을 위함.
	rewardStaticData := RewardStaticData{}
	winDropRewardData := DropRewardData{}

	// 당첨된 drop item 데이터 취합.
	for _, dropEnumId := range dropEnumIdList {
		dropRewardData := eachEnumIdDropRewardData[dropEnumId]
		for i, t := range dropRewardData.TypeArr {
			if t == "" {
				continue
			}
			rewardType := enum.GetCommon_Type(t)
			drop := dropRewardData.DropArr[i]
			value := dropRewardData.ValueArr[i]
			rate := dropRewardData.RateArr[i]
			if helper.IsWinRate(float64(rate)) {
				err = s.SetRewardStaticData(ctx, &rewardStaticData, rewardType, drop)
				if err != nil {
					return err
				}

				winDropRewardData.TypeArr = append(winDropRewardData.TypeArr, t)
				winDropRewardData.DropArr = append(winDropRewardData.DropArr, drop)
				winDropRewardData.ValueArr = append(winDropRewardData.ValueArr, value)
				winDropRewardData.RateArr = append(winDropRewardData.RateArr, rate)
			}
		}
	}

	// 당첨 drop item 을 reward 에 set.
	for i, t := range winDropRewardData.TypeArr {
		rewardType := enum.GetCommon_Type(t)
		drop := winDropRewardData.DropArr[i]
		value := winDropRewardData.ValueArr[i]
		err := s.SetRewardList(ctx, rewardStaticData, reward, rewardType, drop, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *store) SetRewardList(ctx context.Context, rewardStaticData RewardStaticData, reward *model_pb.Reward, rewardType enum.Common_Type, rewardValue string, rewardCount int32) error {
	if rewardCount < 0 {
		rewardCount *= -1
	}

	if rewardType == enum.Common_Type_ASSET {
		assetData, err := s.GetRewardAsset(ctx, rewardStaticData.AssetsStatic[rewardValue])
		if err != nil {
			return err
		}
		assetData.Balance = int64(rewardCount)

		isNew := true
		for _, existAsset := range reward.Assets {
			if existAsset.EnumId == assetData.EnumId {
				existAsset.Balance += assetData.Balance
				isNew = false
				break
			}
		}
		if isNew {
			reward.Assets = append(reward.Assets, assetData)
		}
	} else if rewardType == enum.Common_Type_CHARACTER {
		for i := 0; i < int(rewardCount); i++ {
			characterData, err := s.GetRewardCharacter(ctx, rewardStaticData.CharactersStatic[rewardValue])
			if err != nil {
				return err
			}

			reward.Characters = append(reward.Characters, characterData)
		}
	} else if rewardType == enum.Common_Type_ITEM {
		itemData, err := s.GetRewardItem(ctx, rewardStaticData.ItemsStatic[rewardValue])
		if err != nil {
			return err
		}
		itemData.Count = rewardCount

		isNew := true
		for _, existItem := range reward.Items {
			if existItem.EnumId == itemData.EnumId {
				existItem.Count += itemData.Count
				isNew = false
				break
			}
		}
		if isNew {
			reward.Items = append(reward.Items, itemData)
		}
	} else if rewardType == enum.Common_Type_FATE_CARD {
		for i := 0; i < int(rewardCount); i++ {
			fateCardData, err := s.GetRewardFateCard(ctx, rewardStaticData.FateCardsStatic[rewardValue])
			if err != nil {
				return err
			}

			reward.FateCards = append(reward.FateCards, fateCardData)
		}
	} else {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrDropNotFoundCommonType.String())
	}

	return nil
}

func (s *store) SetRewardStaticData(ctx context.Context, rewardStaticData *RewardStaticData, rewardType enum.Common_Type, rewardEnumId string) (err error) {
	if rewardStaticData.AssetsStatic == nil {
		rewardStaticData.AssetsStatic = map[string]db_static_data.Asset{}
	}
	if rewardStaticData.CharactersStatic == nil {
		rewardStaticData.CharactersStatic = map[string]db_static_data.Character{}
	}
	if rewardStaticData.ItemsStatic == nil {
		rewardStaticData.ItemsStatic = map[string]db_static_data.Item{}
	}
	if rewardStaticData.FateCardsStatic == nil {
		rewardStaticData.FateCardsStatic = map[string]db_static_data.FateCard{}
	}

	if rewardType == enum.Common_Type_ASSET {
		if _, ok := rewardStaticData.AssetsStatic[rewardEnumId]; !ok {
			rewardStaticData.AssetsStatic[rewardEnumId], err = s.StaticDataQueries.GetAssetByEnumId(ctx, rewardEnumId)
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_StaticDataErrAssetNotFound.String())
			}
		}
	} else if rewardType == enum.Common_Type_CHARACTER {
		if _, ok := rewardStaticData.CharactersStatic[rewardEnumId]; !ok {
			rewardStaticData.CharactersStatic[rewardEnumId], err = s.StaticDataQueries.GetCharacterByEnumId(ctx, rewardEnumId)
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterNotFound.String())
			}
		}
	} else if rewardType == enum.Common_Type_ITEM {
		if _, ok := rewardStaticData.ItemsStatic[rewardEnumId]; !ok {
			rewardStaticData.ItemsStatic[rewardEnumId], err = s.StaticDataQueries.GetItemByEnumId(ctx, rewardEnumId)
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_StaticDataErrItemNotFound.String())
			}
		}
	} else if rewardType == enum.Common_Type_FATE_CARD {
		if _, ok := rewardStaticData.FateCardsStatic[rewardEnumId]; !ok {
			rewardStaticData.FateCardsStatic[rewardEnumId], err = s.StaticDataQueries.GetFateCardByEnumId(ctx, rewardEnumId)
			if err != nil {
				return helper.ErrorWithStack(err_pb.Code_StaticDataErrFateCardNotFound.String())
			}
		}
	} else {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrDropNotFoundCommonType.String())
	}

	return nil
}

func (s *store) SetSubReward(subReward *model_pb.SubReward, subId int32, subType enum.Common_Type, subEnumId string, subCount int32) {
	if subCount < 0 {
		subCount *= -1
	}

	subValue := model_pb.SubValue{
		Id:     subId,
		Type:   int32(subType),
		EnumId: subEnumId,
		Count:  subCount,
	}

	if subType == enum.Common_Type_ASSET {
		isNew := true
		for _, existValue := range subReward.Assets {
			if existValue.EnumId == subValue.EnumId {
				existValue.Count += subCount
				isNew = false
				break
			}
		}
		if isNew {
			subReward.Assets = append(subReward.Assets, &subValue)
		}
	} else if subType == enum.Common_Type_CHARACTER {
		subReward.Characters = append(subReward.Characters, &subValue)
	} else if subType == enum.Common_Type_ITEM {
		isNew := true
		for _, existValue := range subReward.Items {
			if existValue.EnumId == subValue.EnumId {
				existValue.Count += subCount
				isNew = false
				break
			}
		}
		if isNew {
			subReward.Items = append(subReward.Items, &subValue)
		}
	}
}

func (s *store) GetRewardAsset(ctx context.Context, assetStatic db_static_data.Asset) (*model_pb.AssetData, error) {
	return &model_pb.AssetData{
		EnumId: assetStatic.EnumID,
		Type:   int32(enum.GetAsset(assetStatic.CeAsset)),
	}, nil
}

func (s *store) GetRewardCharacter(ctx context.Context, characterStatic db_static_data.Character) (*model_pb.CharacterData, error) {
	return &model_pb.CharacterData{
		EnumId: characterStatic.EnumID,
	}, nil
}

func (s *store) GetRewardItem(ctx context.Context, itemStatic db_static_data.Item) (*model_pb.ItemData, error) {
	return &model_pb.ItemData{
		EnumId: itemStatic.EnumID,
	}, nil
}

func (s *store) GetRewardFateCard(ctx context.Context, fateCardStatic db_static_data.FateCard) (*model_pb.FateCardData, error) {
	return &model_pb.FateCardData{
		EnumId: fateCardStatic.EnumID,
	}, nil
}

func (s *store) CreateRewardAsset(ctx context.Context, gameQueries *db_game.Queries, gameDb, accountUserId int32, rewardAsset *model_pb.AssetData) error {
	if gameQueries == nil {
		gameQueries = s.GameQueries[gameDb]
	}

	result, err := gameQueries.UpsertAsset(ctx, db_game.UpsertAssetParams{
		AccountUserID: accountUserId,
		EnumID:        rewardAsset.EnumId,
		Type:          rewardAsset.Type,
		Amount:        rewardAsset.Balance,
	})
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	lastInsertId, _ := result.LastInsertId()
	if lastInsertId > 0 {
		rewardAsset.Id = int32(lastInsertId)
	}

	return nil
}

func (s *store) CreateRewardCharacter(ctx context.Context, gameQueries *db_game.Queries, gameDb, accountUserId int32, rewardCharacter *model_pb.CharacterData) (newReward model_pb.Reward, err error) {
	if gameQueries == nil {
		gameQueries = s.GameQueries[gameDb]
	}

	characterId, _, newReward, err := s.TxCreateCharacter(ctx, gameQueries, gameDb, accountUserId, rewardCharacter.EnumId)
	if err != nil {
		return newReward, err
	}
	if characterId > 0 {
		rewardCharacter.Id = characterId
		rewardCharacter.CreatedAt = timestamppb.Now()
	}

	return newReward, nil
}

func (s *store) CreateRewardItem(ctx context.Context, gameQueries *db_game.Queries, gameDb, accountUserId int32, rewardItem *model_pb.ItemData) error {
	if gameQueries == nil {
		gameQueries = s.GameQueries[gameDb]
	}

	result, err := gameQueries.UpsertItem(ctx, db_game.UpsertItemParams{
		AccountUserID: accountUserId,
		EnumID:        rewardItem.EnumId,
		Count:         rewardItem.Count,
	})
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	lastInsertId, _ := result.LastInsertId()
	if lastInsertId > 0 {
		rewardItem.Id = int32(lastInsertId)
	}

	return nil
}

func (s *store) CreateRewardFateCard(ctx context.Context, gameQueries *db_game.Queries, gameDb, accountUserId int32, rewardFateCard *model_pb.FateCardData) error {
	if gameQueries == nil {
		gameQueries = s.GameQueries[gameDb]
	}

	result, err := gameQueries.CreateFateCard(ctx, db_game.CreateFateCardParams{
		AccountUserID: accountUserId,
		EnumID:        rewardFateCard.EnumId,
	})
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	lastInsertId, _ := result.LastInsertId()
	if lastInsertId > 0 {
		rewardFateCard.Id = int32(lastInsertId)
		rewardFateCard.CreatedAt = timestamppb.Now()
	}

	return nil
}

func (s *store) GetGachaDropCharacter(ctx context.Context, poolGrade enum.Gacha_Pool_Grade, isPickup bool, poolId, pickupPoolId string, gachaDropDList []db_static_data.GachaDropD, gachaDropCList []db_static_data.GachaDropC, gachaDropBList []db_static_data.GachaDropB, gachaDropAList []db_static_data.GachaDropA) (string, error) {
	poolCharacterList := []string{}
	if poolGrade == enum.Gacha_Pool_Grade_D {
		for _, gachaDrop := range gachaDropDList {
			// poolId 가 있으면 포함.
			if strings.Contains(gachaDrop.PoolID, poolId) {
				// pickup 이 아닌데 pickupPoolId 가 포함되있으면 제외.
				if !isPickup && strings.Contains(gachaDrop.PoolID, pickupPoolId) {
					continue
				}

				poolCharacterList = append(poolCharacterList, gachaDrop.DropCharacter)
			}
		}
	} else if poolGrade == enum.Gacha_Pool_Grade_C {
		for _, gachaDrop := range gachaDropCList {
			// poolId 가 있으면 포함.
			if strings.Contains(gachaDrop.PoolID, poolId) {
				// pickup 이 아닌데 pickupPoolId 가 포함되있으면 제외.
				if !isPickup && strings.Contains(gachaDrop.PoolID, pickupPoolId) {
					continue
				}

				poolCharacterList = append(poolCharacterList, gachaDrop.DropCharacter)
			}
		}
	} else if poolGrade == enum.Gacha_Pool_Grade_B {
		for _, gachaDrop := range gachaDropBList {
			// poolId 가 있으면 포함.
			if strings.Contains(gachaDrop.PoolID, poolId) {
				// pickup 이 아닌데 pickUpPoolId 가 포함되있으면 제외.
				if !isPickup && strings.Contains(gachaDrop.PoolID, pickupPoolId) {
					continue
				}

				poolCharacterList = append(poolCharacterList, gachaDrop.DropCharacter)
			}
		}
	} else if poolGrade == enum.Gacha_Pool_Grade_A {
		for _, gachaDrop := range gachaDropAList {
			// poolId 가 있으면 포함.
			if strings.Contains(gachaDrop.PoolID, poolId) {
				// pickup 이 아닌데 pickUpPoolId 가 포함되있으면 제외.
				if !isPickup && strings.Contains(gachaDrop.PoolID, pickupPoolId) {
					continue
				}

				poolCharacterList = append(poolCharacterList, gachaDrop.DropCharacter)
			}
		}
	} else {
		return "", helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}

	if len(poolCharacterList) < 1 {
		return "", helper.ErrorWithStack(err_pb.Code_ServerErr.String())
	}
	winCharacterIndex := helper.GetRandInt(0, len(poolCharacterList)-1)
	winCharacterEnumId := poolCharacterList[winCharacterIndex]

	return winCharacterEnumId, nil
}
