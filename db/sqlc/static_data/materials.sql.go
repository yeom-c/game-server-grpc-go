// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: materials.sql

package db_static_data

import (
	"context"
)

const getMaterials = `-- name: GetMaterials :many
SELECT id, enum_id, ce_common_type_material, material, material_value, created_at FROM materials
`

func (q *Queries) GetMaterials(ctx context.Context) ([]Materials, error) {
	rows, err := q.db.QueryContext(ctx, getMaterials)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Materials{}
	for rows.Next() {
		var i Materials
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.CeCommonTypeMaterial,
			&i.Material,
			&i.MaterialValue,
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

const getMaterialsByEnumId = `-- name: GetMaterialsByEnumId :one
SELECT id, enum_id, ce_common_type_material, material, material_value, created_at FROM materials WHERE enum_id = ?
`

func (q *Queries) GetMaterialsByEnumId(ctx context.Context, enumID string) (Materials, error) {
	row := q.db.QueryRowContext(ctx, getMaterialsByEnumId, enumID)
	var i Materials
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.CeCommonTypeMaterial,
		&i.Material,
		&i.MaterialValue,
		&i.CreatedAt,
	)
	return i, err
}
