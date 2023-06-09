// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: gacha_drop_d.sql

package db_static_data

import (
	"context"
)

const getGachaDropDByEnumId = `-- name: GetGachaDropDByEnumId :one
SELECT id, enum_id, drop_character, pool_id, created_at FROM ` + "`" + `gacha_drop_d` + "`" + ` WHERE enum_id = ?
`

func (q *Queries) GetGachaDropDByEnumId(ctx context.Context, enumID string) (GachaDropD, error) {
	row := q.db.QueryRowContext(ctx, getGachaDropDByEnumId, enumID)
	var i GachaDropD
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.DropCharacter,
		&i.PoolID,
		&i.CreatedAt,
	)
	return i, err
}

const getGachaDropDList = `-- name: GetGachaDropDList :many
SELECT id, enum_id, drop_character, pool_id, created_at FROM ` + "`" + `gacha_drop_d` + "`" + `
`

func (q *Queries) GetGachaDropDList(ctx context.Context) ([]GachaDropD, error) {
	rows, err := q.db.QueryContext(ctx, getGachaDropDList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GachaDropD{}
	for rows.Next() {
		var i GachaDropD
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.DropCharacter,
			&i.PoolID,
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
