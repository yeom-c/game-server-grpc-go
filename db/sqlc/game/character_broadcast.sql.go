// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: character_broadcast.sql

package db_game

import (
	"context"
	"database/sql"
)

const createCharacterBroadcast = `-- name: CreateCharacterBroadcast :execresult
INSERT INTO character_broadcast (account_user_id, character_enum_id, timeline_enum_id, ` + "`" + `type` + "`" + `, on_air, complete) VALUES (?, ?, ?, ?, ?, ?)
`

type CreateCharacterBroadcastParams struct {
	AccountUserID   int32  `json:"account_user_id"`
	CharacterEnumID string `json:"character_enum_id"`
	TimelineEnumID  string `json:"timeline_enum_id"`
	Type            int32  `json:"type"`
	OnAir           int32  `json:"on_air"`
	Complete        int32  `json:"complete"`
}

func (q *Queries) CreateCharacterBroadcast(ctx context.Context, arg CreateCharacterBroadcastParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createCharacterBroadcast,
		arg.AccountUserID,
		arg.CharacterEnumID,
		arg.TimelineEnumID,
		arg.Type,
		arg.OnAir,
		arg.Complete,
	)
}

const getAllCompletedCharacterBroadcastList = `-- name: GetAllCompletedCharacterBroadcastList :many
SELECT id, account_user_id, character_enum_id, timeline_enum_id, type, on_air, complete, broadcasted_at FROM character_broadcast WHERE account_user_id = ? AND complete = 1
`

func (q *Queries) GetAllCompletedCharacterBroadcastList(ctx context.Context, accountUserID int32) ([]CharacterBroadcast, error) {
	rows, err := q.db.QueryContext(ctx, getAllCompletedCharacterBroadcastList, accountUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterBroadcast{}
	for rows.Next() {
		var i CharacterBroadcast
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.CharacterEnumID,
			&i.TimelineEnumID,
			&i.Type,
			&i.OnAir,
			&i.Complete,
			&i.BroadcastedAt,
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

const getCompletedCharacterBroadcastList = `-- name: GetCompletedCharacterBroadcastList :many
SELECT id, account_user_id, character_enum_id, timeline_enum_id, type, on_air, complete, broadcasted_at FROM character_broadcast WHERE account_user_id = ? AND character_enum_id = ? AND complete = 1
`

type GetCompletedCharacterBroadcastListParams struct {
	AccountUserID   int32  `json:"account_user_id"`
	CharacterEnumID string `json:"character_enum_id"`
}

func (q *Queries) GetCompletedCharacterBroadcastList(ctx context.Context, arg GetCompletedCharacterBroadcastListParams) ([]CharacterBroadcast, error) {
	rows, err := q.db.QueryContext(ctx, getCompletedCharacterBroadcastList, arg.AccountUserID, arg.CharacterEnumID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterBroadcast{}
	for rows.Next() {
		var i CharacterBroadcast
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.CharacterEnumID,
			&i.TimelineEnumID,
			&i.Type,
			&i.OnAir,
			&i.Complete,
			&i.BroadcastedAt,
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

const getOnAirCharacterBroadcast = `-- name: GetOnAirCharacterBroadcast :one
SELECT id, account_user_id, character_enum_id, timeline_enum_id, type, on_air, complete, broadcasted_at FROM character_broadcast WHERE id = ?
`

func (q *Queries) GetOnAirCharacterBroadcast(ctx context.Context, id int32) (CharacterBroadcast, error) {
	row := q.db.QueryRowContext(ctx, getOnAirCharacterBroadcast, id)
	var i CharacterBroadcast
	err := row.Scan(
		&i.ID,
		&i.AccountUserID,
		&i.CharacterEnumID,
		&i.TimelineEnumID,
		&i.Type,
		&i.OnAir,
		&i.Complete,
		&i.BroadcastedAt,
	)
	return i, err
}

const getOnAirCharacterBroadcastList = `-- name: GetOnAirCharacterBroadcastList :many
SELECT id, account_user_id, character_enum_id, timeline_enum_id, type, on_air, complete, broadcasted_at FROM character_broadcast WHERE account_user_id = ? AND on_air = 1
`

func (q *Queries) GetOnAirCharacterBroadcastList(ctx context.Context, accountUserID int32) ([]CharacterBroadcast, error) {
	rows, err := q.db.QueryContext(ctx, getOnAirCharacterBroadcastList, accountUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterBroadcast{}
	for rows.Next() {
		var i CharacterBroadcast
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.CharacterEnumID,
			&i.TimelineEnumID,
			&i.Type,
			&i.OnAir,
			&i.Complete,
			&i.BroadcastedAt,
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

const updateCharacterBroadcastComplete = `-- name: UpdateCharacterBroadcastComplete :execresult
UPDATE character_broadcast SET complete = ? WHERE id = ?
`

type UpdateCharacterBroadcastCompleteParams struct {
	Complete int32 `json:"complete"`
	ID       int32 `json:"id"`
}

func (q *Queries) UpdateCharacterBroadcastComplete(ctx context.Context, arg UpdateCharacterBroadcastCompleteParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCharacterBroadcastComplete, arg.Complete, arg.ID)
}

const updateCharacterBroadcastOnAirByAccountUserId = `-- name: UpdateCharacterBroadcastOnAirByAccountUserId :execresult
UPDATE character_broadcast SET on_air = ? WHERE account_user_id = ?
`

type UpdateCharacterBroadcastOnAirByAccountUserIdParams struct {
	OnAir         int32 `json:"on_air"`
	AccountUserID int32 `json:"account_user_id"`
}

func (q *Queries) UpdateCharacterBroadcastOnAirByAccountUserId(ctx context.Context, arg UpdateCharacterBroadcastOnAirByAccountUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateCharacterBroadcastOnAirByAccountUserId, arg.OnAir, arg.AccountUserID)
}

const upsertCharacterBroadcast = `-- name: UpsertCharacterBroadcast :execresult
INSERT INTO character_broadcast (account_user_id, character_enum_id, timeline_enum_id, ` + "`" + `type` + "`" + `, on_air, complete) VALUES (?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE on_air = ?, complete = ?, broadcasted_at = NOW()
`

type UpsertCharacterBroadcastParams struct {
	AccountUserID   int32  `json:"account_user_id"`
	CharacterEnumID string `json:"character_enum_id"`
	TimelineEnumID  string `json:"timeline_enum_id"`
	Type            int32  `json:"type"`
	OnAir           int32  `json:"on_air"`
	Complete        int32  `json:"complete"`
}

func (q *Queries) UpsertCharacterBroadcast(ctx context.Context, arg UpsertCharacterBroadcastParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertCharacterBroadcast,
		arg.AccountUserID,
		arg.CharacterEnumID,
		arg.TimelineEnumID,
		arg.Type,
		arg.OnAir,
		arg.Complete,
		arg.OnAir,
		arg.Complete,
	)
}
