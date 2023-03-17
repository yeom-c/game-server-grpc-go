package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) TxEquipFateCard(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId, characterId, fateCardId int32) (err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 캐릭터 확인.
	character, err := txGameQueries.GetCharacter(ctx, characterId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
		}
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if character.AccountUserID != accountUserId {
		return helper.ErrorWithStack(err_pb.Code_CharacterErrNotFoundCharacter.String())
	}

	// 인연카드 확인.
	fateCard, err := txGameQueries.GetFateCard(ctx, fateCardId)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.ErrorWithStack(err_pb.Code_FateCardErrNotFoundFateCard.String())
		}
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if fateCard.AccountUserID != accountUserId {
		return helper.ErrorWithStack(err_pb.Code_FateCardErrNotFoundFateCard.String())
	}

	// 기존 착용 인연카드 해제.
	_, err = txGameQueries.UnequipFateCardByCharacterEnumId(ctx, db_game.UnequipFateCardByCharacterEnumIdParams{
		AccountUserID:   accountUserId,
		CharacterEnumID: sql.NullString{String: character.EnumID, Valid: true},
	})
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 착용.
	_, err = txGameQueries.UpdateFateCardCharacterEnumId(ctx, db_game.UpdateFateCardCharacterEnumIdParams{
		ID:              fateCardId,
		CharacterEnumID: sql.NullString{String: character.EnumID, Valid: true},
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
