package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	"github.com/yeomc/game-server-grpc-go/helper"
)

func (s *store) TxClearStory(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId int32, storyEnumId string) (newReward model_pb.Reward, storyIndex int32, err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 스토리 데이터.
	storyStageStatic, err := s.StaticDataQueries.GetStoryStageByEnumId(ctx, storyEnumId)
	if err != nil {
		return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_StaticDataErrStoryStageNotFound.String())
	}
	storyIndex = storyStageStatic.StageIndex

	// 이전 스토리 클리어 체크.
	user, err := txGameQueries.GetUserByAccountUserId(ctx, accountUserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_UserErrNotFoundUser.String())
		}
		return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if user.StoryIndex < storyStageStatic.StageIndex-1 {
		return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_StoryErrNotClearedBeforeStory.String())
	}

	// 보상 지급.
	rewardDropEnumId := storyStageStatic.Reward
	if rewardDropEnumId != "" {
		err = s.GetDropRewardList(ctx, &newReward, []string{rewardDropEnumId})
		if err != nil {
			return newReward, storyIndex, err
		}
		err = s.TxCreateRewards(ctx, txGameQueries, gameDb, accountUserId, &newReward)
		if err != nil {
			return newReward, storyIndex, err
		}
	}

	// 스토리 인덱스 저장.
	if user.StoryIndex < storyIndex {
		_, err = txGameQueries.UpdateUserStoryIndexByAccountUserId(ctx, db_game.UpdateUserStoryIndexByAccountUserIdParams{
			AccountUserID: accountUserId,
			StoryIndex:    storyIndex,
		})
		if err != nil {
			return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	if tx != nil {
		err := tx.Commit()
		if err != nil {
			return newReward, storyIndex, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return newReward, storyIndex, nil
}
