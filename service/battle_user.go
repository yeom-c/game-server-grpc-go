package service

import (
	"context"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_battle "github.com/yeomc/game-server-grpc-go/db/sqlc/battle"
	db_common "github.com/yeomc/game-server-grpc-go/db/sqlc/common"
	"github.com/yeomc/game-server-grpc-go/helper"

	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/battle_user"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	"github.com/yeomc/game-server-grpc-go/db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type battleUserService struct {
	pb.BattleUserServiceServer
}

func NewBattleUserService() *battleUserService {
	return &battleUserService{}
}

func (s *battleUserService) GetRankerList(ctx context.Context, _ *model_pb.Empty) (*pb.GetRankerListRes, error) {
	rankerList := []*model_pb.Ranker{}
	page := int32(1)
	limit := int32(10)
	offset := (page - 1) * limit

	rankerUserList, err := db.Store().BattleQueries.GetRankerList(ctx, db_battle.GetRankerListParams{Limit: limit, Offset: offset})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	if len(rankerUserList) > 0 {
		// common db 에서 nickname 을 얻음.
		rankerAccountUserIdList := []int32{}
		for _, rankerUser := range rankerUserList {
			rankerAccountUserIdList = append(rankerAccountUserIdList, rankerUser.AccountUserID)
		}
		rankerAccountUserList, err := db.Store().CommonQueries.GetAccountUserListById(ctx, rankerAccountUserIdList)
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		mapRankerAccountUser := map[int32]db_common.AccountUser{}
		for _, rankerAccountUser := range rankerAccountUserList {
			mapRankerAccountUser[rankerAccountUser.ID] = rankerAccountUser
		}

		// 전달 목록 처리.
		for _, rankerUser := range rankerUserList {
			rankerList = append(rankerList, &model_pb.Ranker{
				AccountUserId: rankerUser.AccountUserID,
				Nickname:      mapRankerAccountUser[rankerUser.AccountUserID].Nickname,
				MatchPoint:    rankerUser.MatchPoint,
				MatchWin:      rankerUser.MatchWin,
				MatchLose:     rankerUser.MatchLose,
				CreatedAt:     timestamppb.New(rankerUser.CreatedAt),
			})
		}
	}

	return &pb.GetRankerListRes{
		RankerList: rankerList,
	}, nil
}
