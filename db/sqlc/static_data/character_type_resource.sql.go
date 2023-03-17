// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: character_type_resource.sql

package db_static_data

import (
	"context"
)

const getCharacterTypeResourceByEnumId = `-- name: GetCharacterTypeResourceByEnumId :one
SELECT id, enum_id, common_enum_value, image_ref, created_at FROM character_type_resource WHERE enum_id = ?
`

func (q *Queries) GetCharacterTypeResourceByEnumId(ctx context.Context, enumID string) (CharacterTypeResource, error) {
	row := q.db.QueryRowContext(ctx, getCharacterTypeResourceByEnumId, enumID)
	var i CharacterTypeResource
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.CommonEnumValue,
		&i.ImageRef,
		&i.CreatedAt,
	)
	return i, err
}

const getCharacterTypeResources = `-- name: GetCharacterTypeResources :many
SELECT id, enum_id, common_enum_value, image_ref, created_at FROM character_type_resource
`

func (q *Queries) GetCharacterTypeResources(ctx context.Context) ([]CharacterTypeResource, error) {
	rows, err := q.db.QueryContext(ctx, getCharacterTypeResources)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterTypeResource{}
	for rows.Next() {
		var i CharacterTypeResource
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.CommonEnumValue,
			&i.ImageRef,
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
