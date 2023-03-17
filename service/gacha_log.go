package service

import (
	"context"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/gacha_log"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	"github.com/yeomc/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type gachaLogService struct {
	pb.GachaLogServiceServer
}

func NewGachaLogService() *gachaLogService {
	return &gachaLogService{}
}

func (s *gachaLogService) GetGachaLogCategories(ctx context.Context, _ *model_pb.Empty) (*pb.GetGachaLogCategoriesRes, error) {
	mySession := db_session.GetMySession(ctx)

	gachaLogCategoriesData, err := db.Store().GameQueries[mySession.GameDb].GetGachaLogCategoryListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &pb.GetGachaLogCategoriesRes{
		GachaLogCategories: gachaLogCategoriesData,
	}, nil
}

func (s *gachaLogService) GetGachaLogs(ctx context.Context, req *pb.GetGachaLogsReq) (*pb.GetGachaLogsRes, error) {
	gachaLogsData := []*model_pb.GachaLogData{}
	mySession := db_session.GetMySession(ctx)

	page := req.Page
	if page < 1 {
		page = 1
	}

	limit := int32(10)
	offset := (page - 1) * limit
	gachaLogs := []db_game.GachaLog{}
	var totalCount int64
	if req.EnumId == "" {
		var err error
		gachaLogs, err = db.Store().GameQueries[mySession.GameDb].GetGachaLogListByAccountUserId(ctx, db_game.GetGachaLogListByAccountUserIdParams{
			AccountUserID: mySession.AccountUserId,
			Offset:        offset,
			Limit:         limit,
		})
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}

		totalCount, err = db.Store().GameQueries[mySession.GameDb].GetGachaLogListCountByAccountUserId(ctx, mySession.AccountUserId)
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	} else {
		var err error
		gachaLogs, err = db.Store().GameQueries[mySession.GameDb].GetGachaLogListByEnumId(ctx, db_game.GetGachaLogListByEnumIdParams{
			AccountUserID: mySession.AccountUserId,
			EnumID:        req.EnumId,
			Offset:        offset,
			Limit:         limit,
		})
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}

		totalCount, err = db.Store().GameQueries[mySession.GameDb].GetGachaLogListCountByEnumId(ctx, db_game.GetGachaLogListCountByEnumIdParams{
			AccountUserID: mySession.AccountUserId,
			EnumID:        req.EnumId,
		})
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	totalPage := totalCount / int64(limit)
	if totalCount%int64(limit) > 0 {
		totalPage += 1
	}

	for _, gachaLog := range gachaLogs {
		gachaLogsData = append(gachaLogsData, &model_pb.GachaLogData{
			Id:              gachaLog.ID,
			EnumId:          gachaLog.EnumID,
			CharacterEnumId: gachaLog.CharacterEnumID,
			CreatedAt:       timestamppb.New(gachaLog.CreatedAt),
		})
	}

	return &pb.GetGachaLogsRes{
		GachaLogs: gachaLogsData,
		TotalPage: totalPage,
	}, nil
}
