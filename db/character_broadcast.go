package db

import (
	"context"
	"encoding/json"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	"github.com/yeomc/game-server-grpc-go/enum"
	"github.com/yeomc/game-server-grpc-go/helper"
)

type broadcastPool struct {
	event     map[string][]string
	regular   map[string][]string
	irregular map[string][]string
	noGet     map[string][]string
}

func (s *store) GetAvailableCharacterBroadcastPoolData(ctx context.Context, gameQueries *db_game.Queries, gameDb, accountUserId int32) (*broadcastPool, error) {
	if gameQueries == nil {
		gameQueries = s.GameQueries[gameDb]
	}
	broadcastPoolData, err := s.GetCharacterBroadcastPoolData(ctx)
	if err != nil {
		return nil, err
	}

	// 획득 캐릭터별 regular, irregular 방송 유지, no_get 방송 제거.
	hasRegular := map[string][]string{}
	hasIrregular := map[string][]string{}
	characterCollectionList, err := gameQueries.GetCharacterCollectionListByAccountUserId(ctx, accountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, characterCollection := range characterCollectionList {
		if _, has := broadcastPoolData.regular[characterCollection.CharacterEnumID]; has {
			hasRegular[characterCollection.CharacterEnumID] = broadcastPoolData.regular[characterCollection.CharacterEnumID]
		}
		if _, has := broadcastPoolData.irregular[characterCollection.CharacterEnumID]; has {
			hasIrregular[characterCollection.CharacterEnumID] = broadcastPoolData.irregular[characterCollection.CharacterEnumID]
		}
		delete(broadcastPoolData.noGet, characterCollection.CharacterEnumID)
	}
	broadcastPoolData.regular = hasRegular
	broadcastPoolData.irregular = hasIrregular

	// 완료한 방송 제거. (event, regular)
	completedBroadcastList, err := gameQueries.GetAllCompletedCharacterBroadcastList(ctx, accountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, completedBroadcast := range completedBroadcastList {
		pool := map[string][]string{}
		broadcastType := enum.Oracle_Broadcast_Type(completedBroadcast.Type)
		if enum.Oracle_Broadcast_Type_EVENT == broadcastType {
			pool = broadcastPoolData.event
		} else if enum.Oracle_Broadcast_Type_REGULAR == broadcastType {
			pool = broadcastPoolData.regular
		} else {
			continue
		}

		if timelineList, has := pool[completedBroadcast.CharacterEnumID]; has {
			findIndex := -1
			for i, timelineEnumId := range timelineList {
				if timelineEnumId == completedBroadcast.TimelineEnumID {
					findIndex = i
					break
				}
			}
			if findIndex >= 0 {
				pool[completedBroadcast.CharacterEnumID] = append(timelineList[:findIndex], timelineList[findIndex+1:]...)
			}
			if len(pool[completedBroadcast.CharacterEnumID]) == 0 {
				delete(pool, completedBroadcast.CharacterEnumID)
			}
		}
	}

	return broadcastPoolData, nil
}

func (s *store) GetCharacterBroadcastPoolData(ctx context.Context) (*broadcastPool, error) {
	broadcastPoolStaticList, err := s.StaticDataQueries.GetVstoryOracleBroadcastPool(ctx)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	broadcastPool := &broadcastPool{
		event:     map[string][]string{},
		regular:   map[string][]string{},
		irregular: map[string][]string{},
		noGet:     map[string][]string{},
	}
	for _, broadcastPoolStatic := range broadcastPoolStaticList {
		var eventArr []string
		var regularArr []string
		var irregularArr []string
		var noGetArr []string
		if broadcastPoolStatic.Event != "" {
			err = json.Unmarshal([]byte(broadcastPoolStatic.Event), &eventArr)
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_StaticDataErrVstoryOracleBroadcastPool.String())
			}
			eventArr = helper.RemoveEmpty(eventArr)
		}
		if broadcastPoolStatic.Regular != "" {
			err = json.Unmarshal([]byte(broadcastPoolStatic.Regular), &regularArr)
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_StaticDataErrVstoryOracleBroadcastPool.String())
			}
			regularArr = helper.RemoveEmpty(regularArr)
		}
		if broadcastPoolStatic.Irregular != "" {
			err = json.Unmarshal([]byte(broadcastPoolStatic.Irregular), &irregularArr)
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_StaticDataErrVstoryOracleBroadcastPool.String())
			}
			irregularArr = helper.RemoveEmpty(irregularArr)
		}
		if broadcastPoolStatic.NoGet != "" {
			err = json.Unmarshal([]byte(broadcastPoolStatic.NoGet), &noGetArr)
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_StaticDataErrVstoryOracleBroadcastPool.String())
			}
			noGetArr = helper.RemoveEmpty(noGetArr)
		}

		broadcastPool.event[broadcastPoolStatic.CharacterEnumID] = eventArr
		broadcastPool.regular[broadcastPoolStatic.CharacterEnumID] = regularArr
		broadcastPool.irregular[broadcastPoolStatic.CharacterEnumID] = irregularArr
		broadcastPool.noGet[broadcastPoolStatic.CharacterEnumID] = noGetArr
	}

	return broadcastPool, nil
}
