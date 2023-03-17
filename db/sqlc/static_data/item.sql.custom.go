package db_static_data

import (
	"context"
	"fmt"
	"strings"
)

const getItemListByEnumId = `-- name: GetItemListByEnumId :many
SELECT id, enum_id, ce_item, ce_item_sub, ce_item_grade, bundle_max, cost_value, grade_resource_thumbnail, grade_resource, image_reference, value, created_at FROM item WHERE enum_id IN ("%s")
`

func (q *Queries) GetItemListByEnumId(ctx context.Context, enumId []string) ([]Item, error) {
	query := fmt.Sprintf(getItemListByEnumId, strings.Join(enumId, "\",\""))
	rows, err := q.db.QueryContext(ctx, query)
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
