// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: vstory_oracle_broadcast_pool.sql

package db_static_data

import (
	"context"
)

const getVstoryOracleBroadcastPool = `-- name: GetVstoryOracleBroadcastPool :many
SELECT id, enum_id, character_enum_id, event, regular, irregular, no_get, created_at FROM vstory_oracle_broadcast_pool
`

func (q *Queries) GetVstoryOracleBroadcastPool(ctx context.Context) ([]VstoryOracleBroadcastPool, error) {
	rows, err := q.db.QueryContext(ctx, getVstoryOracleBroadcastPool)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VstoryOracleBroadcastPool{}
	for rows.Next() {
		var i VstoryOracleBroadcastPool
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.CharacterEnumID,
			&i.Event,
			&i.Regular,
			&i.Irregular,
			&i.NoGet,
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

const getVstoryOracleBroadcastPoolByEnumId = `-- name: GetVstoryOracleBroadcastPoolByEnumId :one
SELECT id, enum_id, character_enum_id, event, regular, irregular, no_get, created_at FROM vstory_oracle_broadcast_pool WHERE enum_id = ?
`

func (q *Queries) GetVstoryOracleBroadcastPoolByEnumId(ctx context.Context, enumID string) (VstoryOracleBroadcastPool, error) {
	row := q.db.QueryRowContext(ctx, getVstoryOracleBroadcastPoolByEnumId, enumID)
	var i VstoryOracleBroadcastPool
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.CharacterEnumID,
		&i.Event,
		&i.Regular,
		&i.Irregular,
		&i.NoGet,
		&i.CreatedAt,
	)
	return i, err
}
