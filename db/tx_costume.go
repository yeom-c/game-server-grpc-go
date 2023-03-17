package db

import (
	"context"
	"database/sql"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"

	db_game "github.com/yeom-c/game-server-grpc-go/db/sqlc/game"
	"github.com/yeom-c/game-server-grpc-go/enum"
	"github.com/yeom-c/game-server-grpc-go/helper"
)

func (s *store) TxEquipCostume(ctx context.Context, txGameQueries *db_game.Queries, gameDb, accountUserId, costumeId int32) (err error) {
	var tx *sql.Tx
	if txGameQueries == nil {
		tx, err = s.GameDb[gameDb].Begin()
		if err != nil {
			return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
		defer tx.Rollback()

		txGameQueries = s.GameQueries[gameDb].WithTx(tx)
	}

	// 착용 코스튬 확인.
	costume, err := txGameQueries.GetCostume(ctx, costumeId)
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_CostumeErrNotFoundCostume.String())
	}
	if costume.AccountUserID != accountUserId {
		return helper.ErrorWithStack(err_pb.Code_CostumeErrNotFoundCostume.String())
	}

	// 기존 착용 코스튬 해제.
	_, err = txGameQueries.UpdateCostumeStateByCharacterEnumId(ctx, db_game.UpdateCostumeStateByCharacterEnumIdParams{
		AccountUserID:   accountUserId,
		CharacterEnumID: costume.CharacterEnumID,
		State:           int32(enum.Equip_State_UNEQUIP),
	})
	if err != nil {
		return helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 착용.
	_, err = txGameQueries.UpdateCostumeState(ctx, db_game.UpdateCostumeStateParams{
		ID:    costumeId,
		State: int32(enum.Equip_State_EQUIP),
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
