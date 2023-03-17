// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: character_base_stat.sql

package db_static_data

import (
	"context"
)

const getCharacterBaseStatByEnumId = `-- name: GetCharacterBaseStatByEnumId :one
SELECT id, enum_id, base_stat, created_at FROM character_base_stat WHERE enum_id = ?
`

func (q *Queries) GetCharacterBaseStatByEnumId(ctx context.Context, enumID string) (CharacterBaseStat, error) {
	row := q.db.QueryRowContext(ctx, getCharacterBaseStatByEnumId, enumID)
	var i CharacterBaseStat
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.BaseStat,
		&i.CreatedAt,
	)
	return i, err
}

const getCharacterBaseStats = `-- name: GetCharacterBaseStats :many
SELECT id, enum_id, base_stat, created_at FROM character_base_stat
`

func (q *Queries) GetCharacterBaseStats(ctx context.Context) ([]CharacterBaseStat, error) {
	rows, err := q.db.QueryContext(ctx, getCharacterBaseStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterBaseStat{}
	for rows.Next() {
		var i CharacterBaseStat
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.BaseStat,
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
