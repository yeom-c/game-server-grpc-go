// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: character.sql

package db_game

import (
	"context"
	"database/sql"
)

const createCharacter = `-- name: CreateCharacter :execresult
INSERT INTO ` + "`" + `character` + "`" + ` (account_user_id, enum_id, exp) VALUES (?, ?, ?)
`

type CreateCharacterParams struct {
	AccountUserID int32  `json:"account_user_id"`
	EnumID        string `json:"enum_id"`
	Exp           int32  `json:"exp"`
}

func (q *Queries) CreateCharacter(ctx context.Context, arg CreateCharacterParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCharacter, arg.AccountUserID, arg.EnumID, arg.Exp)
}

const deleteCharacter = `-- name: DeleteCharacter :execresult
DELETE FROM ` + "`" + `character` + "`" + ` WHERE id = ?
`

func (q *Queries) DeleteCharacter(ctx context.Context, id int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteCharacter, id)
}

const getCharacter = `-- name: GetCharacter :one
SELECT id, account_user_id, enum_id, exp, equipment_level, created_at FROM ` + "`" + `character` + "`" + ` WHERE id = ?
`

func (q *Queries) GetCharacter(ctx context.Context, id int32) (Character, error) {
	row := q.db.QueryRowContext(ctx, getCharacter, id)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.AccountUserID,
		&i.EnumID,
		&i.Exp,
		&i.EquipmentLevel,
		&i.CreatedAt,
	)
	return i, err
}

const getCharacterByEnumId = `-- name: GetCharacterByEnumId :one
SELECT id, account_user_id, enum_id, exp, equipment_level, created_at FROM ` + "`" + `character` + "`" + ` WHERE account_user_id = ? AND enum_id = ? LIMIT 1
`

type GetCharacterByEnumIdParams struct {
	AccountUserID int32  `json:"account_user_id"`
	EnumID        string `json:"enum_id"`
}

func (q *Queries) GetCharacterByEnumId(ctx context.Context, arg GetCharacterByEnumIdParams) (Character, error) {
	row := q.db.QueryRowContext(ctx, getCharacterByEnumId, arg.AccountUserID, arg.EnumID)
	var i Character
	err := row.Scan(
		&i.ID,
		&i.AccountUserID,
		&i.EnumID,
		&i.Exp,
		&i.EquipmentLevel,
		&i.CreatedAt,
	)
	return i, err
}

const getCharacterListByAccountUserId = `-- name: GetCharacterListByAccountUserId :many
SELECT id, account_user_id, enum_id, exp, equipment_level, created_at FROM ` + "`" + `character` + "`" + ` WHERE account_user_id = ?
`

func (q *Queries) GetCharacterListByAccountUserId(ctx context.Context, accountUserID int32) ([]Character, error) {
	rows, err := q.db.QueryContext(ctx, getCharacterListByAccountUserId, accountUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Character{}
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.EnumID,
			&i.Exp,
			&i.EquipmentLevel,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCharacterEquipmentLevel = `-- name: UpdateCharacterEquipmentLevel :execresult
UPDATE ` + "`" + `character` + "`" + ` SET equipment_level = equipment_level + ? WHERE id = ?
`

type UpdateCharacterEquipmentLevelParams struct {
	AddLevel int32 `json:"add_level"`
	ID       int32 `json:"id"`
}

func (q *Queries) UpdateCharacterEquipmentLevel(ctx context.Context, arg UpdateCharacterEquipmentLevelParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCharacterEquipmentLevel, arg.AddLevel, arg.ID)
}

const updateCharacterExp = `-- name: UpdateCharacterExp :execresult
UPDATE ` + "`" + `character` + "`" + ` SET exp = exp + ? WHERE id = ?
`

type UpdateCharacterExpParams struct {
	Exp int32 `json:"exp"`
	ID  int32 `json:"id"`
}

func (q *Queries) UpdateCharacterExp(ctx context.Context, arg UpdateCharacterExpParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCharacterExp, arg.Exp, arg.ID)
}

const updateCharacterExpLimitMax = `-- name: UpdateCharacterExpLimitMax :execresult
UPDATE ` + "`" + `character` + "`" + ` SET exp = IF(exp >= ?, exp, IF(exp + ? >= ?, ?, exp + ?)) WHERE id = ?
`

type UpdateCharacterExpLimitMaxParams struct {
	MaxExp int32 `json:"max_exp"`
	AddExp int32 `json:"add_exp"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateCharacterExpLimitMax(ctx context.Context, arg UpdateCharacterExpLimitMaxParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCharacterExpLimitMax,
		arg.MaxExp,
		arg.AddExp,
		arg.MaxExp,
		arg.MaxExp,
		arg.AddExp,
		arg.ID,
	)
}
