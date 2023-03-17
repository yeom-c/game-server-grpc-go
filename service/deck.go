package service

import (
	"context"
	"database/sql"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/deck"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	"github.com/yeomc/game-server-grpc-go/enum"
	"github.com/yeomc/game-server-grpc-go/helper"
)

type deckService struct {
	pb.DeckServiceServer
}

func NewDeckService() *deckService {
	return &deckService{}
}

func (s *deckService) GetDecks(ctx context.Context, _ *model_pb.Empty) (*pb.GetDecksRes, error) {
	var decksData []*model_pb.DeckData
	mySession := db_session.GetMySession(ctx)

	decks, err := db.Store().GameQueries[mySession.GameDb].GetDeckListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	for _, deck := range decks {
		decksData = append(decksData, &model_pb.DeckData{
			Id:         deck.ID,
			Index:      deck.Index,
			Name:       deck.Name,
			Characters: []int32{deck.CharacterID0.Int32, deck.CharacterID1.Int32, deck.CharacterID2.Int32, deck.CharacterID3.Int32, deck.CharacterID4.Int32},
		})
	}

	return &pb.GetDecksRes{
		Decks: decksData,
	}, nil
}

func (s *deckService) SaveDeck(ctx context.Context, req *pb.SaveDeckReq) (*pb.SaveDeckRes, error) {
	mySession := db_session.GetMySession(ctx)
	gameStore := db.Store().GameQueries[mySession.GameDb]

	var reqChars []int32
	for _, character := range req.Characters {
		if character > 0 {
			reqChars = append(reqChars, character)
		}
	}

	if len(reqChars) > 0 {
		characters, err := gameStore.GetCharacterListById(ctx, reqChars)
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		if len(characters) != len(reqChars) {
			return nil, helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
		for _, character := range characters {
			if character.AccountUserID != mySession.AccountUserId {
				return nil, helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
			}
		}
	}

	var deckId int32
	req.Name = strings.Trim(req.Name, " ")
	if req.Id.GetValue() == 0 {
		// 덱 생성.
		deckCount, err := gameStore.GetDeckCount(ctx, mySession.AccountUserId)
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		if deckCount >= int64(enum.ServerEnum["max_deck_count"].(int)) {
			return nil, helper.ErrorWithStack(err_pb.Code_DeckErrExceedCountLimit.String())
		}

		result, err := gameStore.CreateDeck(ctx, db_game.CreateDeckParams{
			AccountUserID: mySession.AccountUserId,
			Index:         req.Index,
			Name:          req.Name,
			CharacterID0:  sql.NullInt32{Int32: req.Characters[0], Valid: req.Characters[0] > 0},
			CharacterID1:  sql.NullInt32{Int32: req.Characters[1], Valid: req.Characters[1] > 0},
			CharacterID2:  sql.NullInt32{Int32: req.Characters[2], Valid: req.Characters[2] > 0},
			CharacterID3:  sql.NullInt32{Int32: req.Characters[3], Valid: req.Characters[3] > 0},
			CharacterID4:  sql.NullInt32{Int32: req.Characters[4], Valid: req.Characters[4] > 0},
		})
		if err != nil {
			if err.(*mysql.MySQLError).Number == 1062 {
				return nil, helper.ErrorWithStack(err_pb.Code_DeckErrDuplicatedDeck.String())
			}
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		deckId = int32(lastInsertId)
	} else {
		// 덱 수정.
		deckId = req.Id.Value
		_, err := gameStore.UpdateDeck(ctx, db_game.UpdateDeckParams{
			ID:           deckId,
			Name:         req.Name,
			CharacterID0: sql.NullInt32{Int32: req.Characters[0], Valid: req.Characters[0] > 0},
			CharacterID1: sql.NullInt32{Int32: req.Characters[1], Valid: req.Characters[1] > 0},
			CharacterID2: sql.NullInt32{Int32: req.Characters[2], Valid: req.Characters[2] > 0},
			CharacterID3: sql.NullInt32{Int32: req.Characters[3], Valid: req.Characters[3] > 0},
			CharacterID4: sql.NullInt32{Int32: req.Characters[4], Valid: req.Characters[4] > 0},
		})
		if err != nil {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	return &pb.SaveDeckRes{
		Deck: &model_pb.DeckData{
			Id:         deckId,
			Index:      req.Index,
			Name:       req.Name,
			Characters: req.Characters,
		},
	}, nil
}
