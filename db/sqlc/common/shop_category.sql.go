// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: shop_category.sql

package db_common

import (
	"context"
)

const getShopCategoryListByShopId = `-- name: GetShopCategoryListByShopId :many
SELECT id, shop_id, visible, name, ` + "`" + `order` + "`" + ` FROM shop_category WHERE shop_id = ? ORDER BY ` + "`" + `order` + "`" + ` ASC, id DESC
`

func (q *Queries) GetShopCategoryListByShopId(ctx context.Context, shopID int32) ([]ShopCategory, error) {
	rows, err := q.db.QueryContext(ctx, getShopCategoryListByShopId, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ShopCategory{}
	for rows.Next() {
		var i ShopCategory
		if err := rows.Scan(
			&i.ID,
			&i.ShopID,
			&i.Visible,
			&i.Name,
			&i.Order,
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
