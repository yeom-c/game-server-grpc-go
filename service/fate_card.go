package service

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/fate_card"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	"github.com/yeomc/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type fateCardService struct {
	pb.FateCardServiceServer
}

func NewFateCardService() *fateCardService {
	return &fateCardService{}
}

func (s *fateCardService) GetFateCards(ctx context.Context, _ *model_pb.Empty) (*pb.GetFateCardsRes, error) {
	var fateCardsData []*model_pb.FateCardData
	mySession := db_session.GetMySession(ctx)

	fateCardList, err := db.Store().GameQueries[mySession.GameDb].GetFateCardListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, fateCard := range fateCardList {
		fateCardsData = append(fateCardsData, &model_pb.FateCardData{
			Id:              fateCard.ID,
			EnumId:          fateCard.EnumID,
			CharacterEnumId: fateCard.CharacterEnumID.String,
			CreatedAt:       timestamppb.New(fateCard.CreatedAt),
		})
	}

	return &pb.GetFateCardsRes{
		FateCards: fateCardsData,
	}, nil
}

func (s *fateCardService) EquipFateCard(ctx context.Context, req *pb.EquipFateCardReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	err := db.Store().TxEquipFateCard(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.CharacterId, req.FateCardId)
	if err != nil {
		return nil, err
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}

func (s *fateCardService) UnquipFateCard(ctx context.Context, req *pb.UnequipFateCardReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	fateCard, err := db.Store().GameQueries[mySession.GameDb].GetFateCard(ctx, req.FateCardId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_FateCardErrNotFoundFateCard.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if fateCard.AccountUserID != mySession.AccountUserId {
		return nil, helper.ErrorWithStack(err_pb.Code_FateCardErrNotFoundFateCard.String())
	}

	_, err = db.Store().GameQueries[mySession.GameDb].UpdateFateCardCharacterEnumId(ctx, db_game.UpdateFateCardCharacterEnumIdParams{
		ID:              req.FateCardId,
		CharacterEnumID: sql.NullString{},
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &model_pb.Result{
		Result: 1,
	}, nil
}
