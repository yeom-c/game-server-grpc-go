package db

import (
	"context"
	"database/sql"
	"encoding/json"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	db_static_data "github.com/yeomc/game-server-grpc-go/db/sqlc/static_data"
	"github.com/yeomc/game-server-grpc-go/enum"
	"github.com/yeomc/game-server-grpc-go/helper"
)

func (s *store) TxCreateCharacter(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, enumId string) (characterId, collectionId int32, newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	characterStatic, err := s.StaticDataQueries.GetCharacterByEnumId(ctx, enumId)
	if err != nil {
		return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterNotFound.String())
	}

	// 캐릭터 생성.
	result, err := txGameQueries.CreateCharacter(ctx, db_game.CreateCharacterParams{
		AccountUserID: accountUserId,
		EnumID:        enumId,
	})
	if err != nil {
		return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	characterId = int32(lastInsertId)

	// 기본 코스튬 생성.
	if characterStatic.CostumeBundle != "" {
		costumeBundleStatic, err := s.StaticDataQueries.GetCharacterCostumeBundleByEnumId(ctx, characterStatic.CostumeBundle)
		if err != nil {
			return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterCostumeBundleNotFound.String())
		}

		costumes := []string{}
		err = json.Unmarshal([]byte(costumeBundleStatic.Costume), &costumes)
		if err != nil {
			return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterCostumeBundle.String())
		}
		if len(costumes) == 0 {
			return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterCostumeBundle.String())
		}

		_, err = txGameQueries.CreateCostume(ctx, db_game.CreateCostumeParams{
			AccountUserID:   accountUserId,
			EnumID:          costumes[0],
			CharacterEnumID: enumId,
			State:           int32(enum.Equip_State_EQUIP),
		})
		if err != nil {
			return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	// 도감 생성.
	_, err = txGameQueries.UpsertCharacterCollection(ctx, db_game.UpsertCharacterCollectionParams{
		AccountUserID:   accountUserId,
		CharacterEnumID: enumId,
		AffectionExp:    0,
		Count:           1,
	})
	if err != nil {
		return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return characterId, collectionId, newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return characterId, collectionId, newReward, nil
}

func (s *store) TxLevelUpSignatureWeapon(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, targetCharacterId int32, materialCharactersId []int32) (err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 대상 캐릭터 정보.
	targetCharacter, err := txGameQueries.GetCharacter(ctx, targetCharacterId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if targetCharacter.AccountUserID != accountUserId {
		return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
	}

	characterStatic, err := s.StaticDataQueries.GetCharacterByEnumId(ctx, targetCharacter.EnumID)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrCharacterNotFound.String())
	}

	// 재료 캐릭터 덱 세팅 확인.
	deckList, err := txGameQueries.GetDeckListByAccountUserId(ctx, accountUserId)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	mapDeckCharacterId := map[int32]int32{}
	for _, deck := range deckList {
		if deck.CharacterID0.Int32 > 0 {
			mapDeckCharacterId[deck.CharacterID0.Int32] = deck.CharacterID0.Int32
		}
		if deck.CharacterID1.Int32 > 0 {
			mapDeckCharacterId[deck.CharacterID1.Int32] = deck.CharacterID1.Int32
		}
		if deck.CharacterID2.Int32 > 0 {
			mapDeckCharacterId[deck.CharacterID2.Int32] = deck.CharacterID2.Int32
		}
		if deck.CharacterID3.Int32 > 0 {
			mapDeckCharacterId[deck.CharacterID3.Int32] = deck.CharacterID3.Int32
		}
		if deck.CharacterID4.Int32 > 0 {
			mapDeckCharacterId[deck.CharacterID4.Int32] = deck.CharacterID4.Int32
		}
	}
	for _, materialCharacterId := range materialCharactersId {
		if _, has := mapDeckCharacterId[materialCharacterId]; has {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
	}

	// 재료 캐릭터 정보.
	enumIdMaterialCharactersId := map[string][]int32{}
	materialCharacters, err := txGameQueries.GetCharacterListById(ctx, materialCharactersId)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, materialCharacter := range materialCharacters {
		if materialCharacter.AccountUserID != accountUserId {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
		if materialCharacter.ID == targetCharacterId {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
		if _, ok := enumIdMaterialCharactersId[materialCharacter.EnumID]; !ok {
			enumIdMaterialCharactersId[materialCharacter.EnumID] = []int32{}
		}

		enumIdMaterialCharactersId[materialCharacter.EnumID] = append(enumIdMaterialCharactersId[materialCharacter.EnumID], materialCharacter.ID)
	}

	// 전용 장비 정보.
	signatureWeaponStatic, err := s.StaticDataQueries.GetSignatureWeaponByEnumId(ctx, characterStatic.SignatureWeapon)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrSignatureWeaponNotFound.String())
	}

	// 재료 차감.
	materialsStatic, err := s.StaticDataQueries.GetMaterialsByEnumId(ctx, signatureWeaponStatic.WeaponGrowthMaterials)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterialsNotFound.String())
	}
	var typeArr []string
	var valueArr []string
	var amountArr []int32
	err = json.Unmarshal([]byte(materialsStatic.CeCommonTypeMaterial), &typeArr)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterials.String())
	}
	err = json.Unmarshal([]byte(materialsStatic.Material), &valueArr)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterials.String())
	}
	err = json.Unmarshal([]byte(materialsStatic.MaterialValue), &amountArr)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_StaticDataErrMaterials.String())
	}

	subReward := model_pb.SubReward{}
	for i, t := range typeArr {
		if t == "" {
			continue
		}

		subType := enum.GetCommon_Type(t)
		subValue := valueArr[i]
		subAmount := amountArr[i]
		if subType == enum.Common_Type_CHARACTER {
			if len(enumIdMaterialCharactersId[subValue]) < int(subAmount) {
				return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
			}
			for j := 0; j < int(subAmount); j++ {
				materialId := enumIdMaterialCharactersId[subValue][0]
				s.SetSubReward(&subReward, materialId, subType, subValue, 1)

				enumIdMaterialCharactersId[subValue] = enumIdMaterialCharactersId[subValue][1:]
			}
		} else {
			s.SetSubReward(&subReward, 0, subType, subValue, subAmount)
		}
	}
	err = s.TxSubRewards(ctx, txGameQueries, gameDb, accountUserId, &subReward)
	if err != nil {
		return err
	}

	// 레벨업.
	_, err = txGameQueries.UpdateCharacterEquipmentLevel(ctx, db_game.UpdateCharacterEquipmentLevelParams{
		ID:       targetCharacterId,
		AddLevel: 1,
	})
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return nil
}

func (s *store) TxGetDeckCharacters(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId, deckId int32) (characters []db_game.Character, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return characters, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	deck, err := txGameQueries.GetDeck(ctx, deckId)
	if err != nil {
		if err == sql.ErrNoRows {
			return characters, helper.ErrorWithStack(err_pb.Code_DeckErrNotFoundDeck.String())
		}
		return characters, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if deck.AccountUserID != accountUserId {
		return characters, helper.ErrorWithStack(err_pb.Code_DeckErrNotFoundDeck.String())
	}

	charactersId := helper.RemoveZero([]int32{deck.CharacterID0.Int32, deck.CharacterID1.Int32, deck.CharacterID2.Int32, deck.CharacterID3.Int32, deck.CharacterID4.Int32})
	if len(charactersId) > 0 {
		characters, err = txGameQueries.GetCharacterListByIdAndAccountUserId(ctx, charactersId, accountUserId)
		if err != nil {
			return characters, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return characters, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return characters, nil
}

func (s *store) TxExtinctCharacter(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, targetCharactersId []int32) (newReward model_pb.Reward, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	if len(targetCharactersId) == 0 {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 덱 정보.
	deckList, err := txGameQueries.GetDeckListByAccountUserId(ctx, accountUserId)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	deckCharacterMap := map[int32]int32{}
	for _, deck := range deckList {
		if deck.CharacterID0.Valid {
			deckCharacterMap[deck.CharacterID0.Int32] = deck.ID
		}
		if deck.CharacterID1.Valid {
			deckCharacterMap[deck.CharacterID1.Int32] = deck.ID
		}
		if deck.CharacterID2.Valid {
			deckCharacterMap[deck.CharacterID2.Int32] = deck.ID
		}
		if deck.CharacterID3.Valid {
			deckCharacterMap[deck.CharacterID3.Int32] = deck.ID
		}
		if deck.CharacterID4.Valid {
			deckCharacterMap[deck.CharacterID4.Int32] = deck.ID
		}
	}

	// 캐릭터 정보 체크.
	characterList, err := txGameQueries.GetCharacterListById(ctx, targetCharactersId)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if len(characterList) == 0 {
		return newReward, helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
	}
	characterEnumIdList := []string{}
	for _, character := range characterList {
		// 다른유저 캐릭터 불가.
		if character.AccountUserID != accountUserId {
			return newReward, helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
		// 덱 세팅 캐릭터 불가.
		if _, has := deckCharacterMap[character.ID]; has {
			return newReward, helper.ErrorWithStack(err_pb.Code_CharacterErrNotExtinctInDeckCharacter.String())
		}
		characterEnumIdList = append(characterEnumIdList, character.EnumID)
	}
	characterStaticList, err := s.StaticDataQueries.GetCharactersByEnumIds(ctx, characterEnumIdList)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	characterStaticMap := map[string]*db_static_data.Character{}
	for _, characterStatic := range characterStaticList {
		characterStaticMap[characterStatic.EnumID] = &characterStatic
	}

	// 캐릭터 삭제.
	_, err = txGameQueries.DeleteCharacters(ctx, targetCharactersId)
	if err != nil {
		return newReward, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 보상 지급.
	rewardType := enum.ServerEnum["extinct_character_reward_type"].(enum.Common_Type)
	rewardValue := enum.ServerEnum["extinct_character_reward_value"].(string)
	var rewardCount int32
	for _, character := range characterList {
		characterStatic := characterStaticMap[character.EnumID]
		characterGrade := enum.GetCharacter_Grade(characterStatic.CeCharacterGrade)
		if characterGrade == enum.Character_Grade_GRADE_D {
			rewardCount += 1
		} else if characterGrade == enum.Character_Grade_GRADE_C {
			rewardCount += 2
		} else if characterGrade == enum.Character_Grade_GRADE_B {
			rewardCount += 4
		} else if characterGrade == enum.Character_Grade_GRADE_A {
			rewardCount += 8
		}
	}
	if rewardCount > 0 {
		resultStaticData := RewardStaticData{}
		err = s.SetRewardStaticData(ctx, &resultStaticData, rewardType, rewardValue)
		if err != nil {
			return newReward, err
		}
		err = s.SetRewardList(ctx, resultStaticData, &newReward, rewardType, rewardValue, rewardCount)
		if err != nil {
			return newReward, err
		}
		err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
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
