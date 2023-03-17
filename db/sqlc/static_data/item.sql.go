// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: item.sql

package db_static_data

import (
	"context"
)

const getItemByEnumId = `-- name: GetItemByEnumId :one
SELECT id, enum_id, ce_item, ce_item_sub, ce_item_grade, bundle_max, cost_value, grade_resource_thumbnail, grade_image_reference, grade_resource, image_reference, value, created_at FROM item WHERE enum_id = ?
`

func (q *Queries) GetItemByEnumId(ctx context.Context, enumID string) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItemByEnumId, enumID)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.CeItem,
		&i.CeItemSub,
		&i.CeItemGrade,
		&i.BundleMax,
		&i.CostValue,
		&i.GradeResourceThumbnail,
		&i.GradeImageReference,
		&i.GradeResource,
		&i.ImageReference,
		&i.Value,
		&i.CreatedAt,
	)
	return i, err
}

const getItems = `-- name: GetItems :many
SELECT id, enum_id, ce_item, ce_item_sub, ce_item_grade, bundle_max, cost_value, grade_resource_thumbnail, grade_image_reference, grade_resource, image_reference, value, created_at FROM item
`

func (q *Queries) GetItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.CeItem,
			&i.CeItemSub,
			&i.CeItemGrade,
			&i.BundleMax,
			&i.CostValue,
			&i.GradeResourceThumbnail,
			&i.GradeImageReference,
			&i.GradeResource,
			&i.ImageReference,
			&i.Value,
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
