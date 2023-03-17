package db

import (
	"context"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	db_common "github.com/yeom-c/game-server-grpc-go/db/sqlc/common"
	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) GetEachAccountUserCharacters(ctx context.Context, eachAccountUserCharactersId map[int32][]int32) (map[int32]db_common.AccountUser, map[int32]map[int32]db_game.Character, error) {
	eachAccountUsers := map[int32]db_common.AccountUser{}                // accountUserId -> accountUser
	eachAccountUserCharacters := map[int32]map[int32]db_game.Character{} // accountUserId -> characterId -> character

	accountUsersId := []int32{}
	for accountUserId, charactersId := range eachAccountUserCharactersId {
		if len(charactersId) > 0 {
			accountUsersId = append(accountUsersId, accountUserId)
		}
	}

	if len(accountUsersId) > 0 {
		accountUsers, err := s.CommonQueries.GetAccountUserListById(ctx, accountUsersId)
		if err != nil {
			return nil, nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		if len(accountUsersId) != len(accountUsers) {
			return nil, nil, helper.ErrorWithStack(err_pb.Code_AccountErrNotFoundAccount.String())
		}

		// id -> accountUser map 변환
		for _, accountUser := range accountUsers {
			eachAccountUsers[accountUser.ID] = accountUser
		}

		// gameDb -> charactersId map 변환
		eachDbCharactersIdMap := map[int32][]int32{}
		for accountUserId, eachDbCharactersId := range eachAccountUserCharactersId {
			if len(eachDbCharactersId) == 0 {
				continue
			}
			gameDb := eachAccountUsers[accountUserId].GameDb
			eachDbCharactersIdMap[gameDb] = append(eachDbCharactersIdMap[gameDb], eachDbCharactersId...)
		}

		// get all characters
		allCharacters := []db_game.Character{}
		for gameDb, charactersId := range eachDbCharactersIdMap {
			if len(charactersId) > 0 {
				characters, err := s.GameQueries[gameDb].GetCharacterListById(ctx, charactersId)
				if err != nil {
					return nil, nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
				}

				allCharacters = append(allCharacters, characters...)
			}
		}

		// accountUserId -> characterId -> character map 변환
		for _, character := range allCharacters {
			if _, ok := eachAccountUserCharacters[character.AccountUserID]; !ok {
				eachAccountUserCharacters[character.AccountUserID] = map[int32]db_game.Character{}
			}
			eachAccountUserCharacters[character.AccountUserID][character.ID] = character
		}
	}

	return eachAccountUsers, eachAccountUserCharacters, nil
}
