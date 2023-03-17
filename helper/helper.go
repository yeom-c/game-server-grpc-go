package helper

import (
	"encoding/json"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ErrorWithStack(err string) error {
	return errors.WithStack(errors.New(err))
}

func ProtoToMap(message proto.Message) (map[string]interface{}, error) {
	var result map[string]interface{}
	marshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}
	messageByte, err := marshaler.Marshal(message)

	json.Unmarshal(messageByte, &result)

	return result, err
}

func Contains(slice []string, find string) bool {
	for _, item := range slice {
		if item == find {
			return true
		}
	}

	return false
}

func Remove(slice []interface{}, i int) []interface{} {
	return append(slice[:i], slice[i+1:]...)
}

func RemoveZero(slice []int32) []int32 {
	result := []int32{}
	for _, item := range slice {
		if item != 0 {
			result = append(result, item)
		}
	}

	return result
}

func RemoveEmpty(slice []string) []string {
	result := []string{}
	for _, item := range slice {
		if item != "" {
			result = append(result, item)
		}
	}

	return result
}

func GetLevel(table []int, exp int) int {
	level := 0
	for i, start := range table {
		if exp >= start {
			level = i + 1
			continue
		}
		break
	}

	return level
}

func GetMaxExp(table []int, level int) int {
	maxExp := 0
	if len(table) > 0 && level > 0 {
		if level >= len(table) {
			maxExp = table[len(table)-1]
		} else {
			maxExp = table[level] - 1
		}
	}

	return maxExp
}

func MergeReward(reward1, reward2 model_pb.Reward) model_pb.Reward {
	merge := model_pb.Reward{}

	eachEnumIdAssetData := map[string]*model_pb.AssetData{}
	for _, asset := range reward1.Assets {
		if _, ok := eachEnumIdAssetData[asset.EnumId]; !ok {
			eachEnumIdAssetData[asset.EnumId] = &model_pb.AssetData{
				Id:      asset.Id,
				EnumId:  asset.EnumId,
				Type:    asset.Type,
				Balance: asset.Balance,
			}
		} else {
			eachEnumIdAssetData[asset.EnumId].Balance += asset.Balance
		}
	}
	for _, asset := range reward2.Assets {
		if _, ok := eachEnumIdAssetData[asset.EnumId]; !ok {
			eachEnumIdAssetData[asset.EnumId] = &model_pb.AssetData{
				Id:      asset.Id,
				EnumId:  asset.EnumId,
				Type:    asset.Type,
				Balance: asset.Balance,
			}
		} else {
			eachEnumIdAssetData[asset.EnumId].Balance += asset.Balance
		}
	}
	for _, assetData := range eachEnumIdAssetData {
		merge.Assets = append(merge.Assets, assetData)
	}

	for _, character := range reward1.Characters {
		merge.Characters = append(merge.Characters, &model_pb.CharacterData{
			Id:             character.Id,
			EnumId:         character.EnumId,
			Exp:            character.Exp,
			EquipmentLevel: character.EquipmentLevel,
			CreatedAt:      character.CreatedAt,
		})
	}
	for _, character := range reward2.Characters {
		merge.Characters = append(merge.Characters, &model_pb.CharacterData{
			Id:             character.Id,
			EnumId:         character.EnumId,
			Exp:            character.Exp,
			EquipmentLevel: character.EquipmentLevel,
			CreatedAt:      character.CreatedAt,
		})
	}

	eachEnumIdItemData := map[string]*model_pb.ItemData{}
	for _, item := range reward1.Items {
		if _, ok := eachEnumIdItemData[item.EnumId]; !ok {
			eachEnumIdItemData[item.EnumId] = &model_pb.ItemData{
				Id:     item.Id,
				EnumId: item.EnumId,
				Count:  item.Count,
			}
		} else {
			eachEnumIdItemData[item.EnumId].Count += item.Count
		}
	}
	for _, item := range reward2.Items {
		if _, ok := eachEnumIdItemData[item.EnumId]; !ok {
			eachEnumIdItemData[item.EnumId] = &model_pb.ItemData{
				Id:     item.Id,
				EnumId: item.EnumId,
				Count:  item.Count,
			}
		} else {
			eachEnumIdItemData[item.EnumId].Count += item.Count
		}
	}
	for _, itemData := range eachEnumIdItemData {
		merge.Items = append(merge.Items, itemData)
	}

	for _, fateCard := range reward1.FateCards {
		merge.FateCards = append(merge.FateCards, &model_pb.FateCardData{
			Id:              fateCard.Id,
			EnumId:          fateCard.EnumId,
			CharacterEnumId: fateCard.CharacterEnumId,
			CreatedAt:       fateCard.CreatedAt,
		})
	}
	for _, fateCard := range reward2.FateCards {
		merge.FateCards = append(merge.FateCards, &model_pb.FateCardData{
			Id:              fateCard.Id,
			EnumId:          fateCard.EnumId,
			CharacterEnumId: fateCard.CharacterEnumId,
			CreatedAt:       fateCard.CreatedAt,
		})
	}

	return merge
}
