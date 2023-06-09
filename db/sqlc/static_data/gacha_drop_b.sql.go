// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: gacha_drop_b.sql

package db_static_data

import (
	"context"
)

const getGachaDropBByEnumId = `-- name: GetGachaDropBByEnumId :one
SELECT id, enum_id, drop_character, pool_id, created_at FROM ` + "`" + `gacha_drop_b` + "`" + ` WHERE enum_id = ?
`

func (q *Queries) GetGachaDropBByEnumId(ctx context.Context, enumID string) (GachaDropB, error) {
	row := q.db.QueryRowContext(ctx, getGachaDropBByEnumId, enumID)
	var i GachaDropB
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.DropCharacter,
		&i.PoolID,
		&i.CreatedAt,
	)
	return i, err
}

const getGachaDropBList = `-- name: GetGachaDropBList :many
SELECT id, enum_id, drop_character, pool_id, created_at FROM ` + "`" + `gacha_drop_b` + "`" + `
`

func (q *Queries) GetGachaDropBList(ctx context.Context) ([]GachaDropB, error) {
	rows, err := q.db.QueryContext(ctx, getGachaDropBList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GachaDropB{}
	for rows.Next() {
		var i GachaDropB
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
