// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: character_collection.sql

package db_game

import (
	"context"
	"database/sql"
)

const getCharacterCollectionCountByAccountUserId = `-- name: GetCharacterCollectionCountByAccountUserId :one
SELECT COUNT(*) FROM character_collection WHERE account_user_id = ?
`

func (q *Queries) GetCharacterCollectionCountByAccountUserId(ctx context.Context, accountUserID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, getCharacterCollectionCountByAccountUserId, accountUserID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getCharacterCollectionListByAccountUserId = `-- name: GetCharacterCollectionListByAccountUserId :many
SELECT id, account_user_id, character_enum_id, affection_exp, count, created_at FROM character_collection WHERE account_user_id = ?
`

func (q *Queries) GetCharacterCollectionListByAccountUserId(ctx context.Context, accountUserID int32) ([]CharacterCollection, error) {
	rows, err := q.db.QueryContext(ctx, getCharacterCollectionListByAccountUserId, accountUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterCollection{}
	for rows.Next() {
		var i CharacterCollection
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.CharacterEnumID,
			&i.AffectionExp,
			&i.Count,
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

const upsertCharacterCollection = `-- name: UpsertCharacterCollection :execresult
INSERT INTO character_collection (account_user_id, character_enum_id, affection_exp, ` + "`" + `count` + "`" + `) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE affection_exp = affection_exp + ?, ` + "`" + `count` + "`" + ` = ` + "`" + `count` + "`" + ` + ?
`

type UpsertCharacterCollectionParams struct {
	AccountUserID   int32  `json:"account_user_id"`
	CharacterEnumID string `json:"character_enum_id"`
	AffectionExp    int32  `json:"affection_exp"`
	Count           int32  `json:"count"`
}

func (q *Queries) UpsertCharacterCollection(ctx context.Context, arg UpsertCharacterCollectionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, upsertCharacterCollection,
		arg.AccountUserID,
		arg.CharacterEnumID,
		arg.AffectionExp,
		arg.Count,
		arg.AffectionExp,
		arg.Count,
	)
}