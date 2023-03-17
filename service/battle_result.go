package service

import (
	"context"
	"database/sql"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/battle_result"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	db_battle "github.com/yeom-c/game-server-grpc-go/db/sqlc/battle"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

type battleResultService struct {
	pb.BattleResultServiceServer
}

func NewBattleResultService() *battleResultService {
	return &battleResultService{}
}

func (s *battleResultService) ConfirmBattleResult(ctx context.Context, req *pb.ConfirmBattleResultReq) (*pb.ConfirmBattleResultRes, error) {
	mySession := db_session.GetMySession(ctx)

	// TODO: 결과 받아서 저장 임시처리 (추후 배틀서버에서 결과 저장하고 해당 코드 삭제 필요).
	_, err := db.Store().BattleQueries.UpdateBattleResultResultByChannelId(ctx, db_battle.UpdateBattleResultResultByChannelIdParams{
		AccountUserID: mySession.AccountUserId,
		ChannelID:     req.BattleChannelId,
		Result:        req.BattleResult,
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	battleResult, err := db.Store().BattleQueries.GetBattleResultByChannelId(ctx, db_battle.GetBattleResultByChannelIdParams{
		AccountUserID: mySession.AccountUserId,
		ChannelID:     req.BattleChannelId,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_BattleResultErrNotFoundBattleResult.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if battleResult.ConfirmedAt.Valid {
		return nil, helper.ErrorWithStack(err_pb.Code_BattleResultErrAlreadyConfirmed.String())
	}

	// 확인 트랜잭션 처리.
	matchPoint, addPoint, reward, err := db.Store().TxConfirmBattleResults(ctx, nil, nil, mySession.GameDb, mySession.AccountUserId, []db_battle.BattleResult{battleResult})
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmBattleResultRes{
		MatchPoint: matchPoint,
		AddPoint:   addPoint,
		Reward:     &reward,
	}, nil
}
