// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: shop_goods.sql

package db_common

import (
	"context"
)

const getShopGoods = `-- name: GetShopGoods :one
SELECT id, shop_category_id, type, enum_id, info, name, ` + "`" + `desc` + "`" + `, cost_type, cost_enum_id, cost, original_cost, count, visible, start_at, end_at FROM shop_goods WHERE id = ?
`

func (q *Queries) GetShopGoods(ctx context.Context, id int32) (ShopGoods, error) {
	row := q.db.QueryRowContext(ctx, getShopGoods, id)
	var i ShopGoods
	err := row.Scan(
		&i.ID,
		&i.ShopCategoryID,
		&i.Type,
		&i.EnumID,
		&i.Info,
		&i.Name,
		&i.Desc,
		&i.CostType,
		&i.CostEnumID,
		&i.Cost,
		&i.OriginalCost,
		&i.Count,
		&i.Visible,
		&i.StartAt,
		&i.EndAt,
	)
	return i, err
}

const getShopGoodsListByShopCategoryId = `-- name: GetShopGoodsListByShopCategoryId :many
SELECT id, shop_category_id, type, enum_id, info, name, ` + "`" + `desc` + "`" + `, cost_type, cost_enum_id, cost, original_cost, count, visible, start_at, end_at FROM shop_goods WHERE shop_category_id = ?
`

func (q *Queries) GetShopGoodsListByShopCategoryId(ctx context.Context, shopCategoryID int32) ([]ShopGoods, error) {
	rows, err := q.db.QueryContext(ctx, getShopGoodsListByShopCategoryId, shopCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ShopGoods{}
	for rows.Next() {
		var i ShopGoods
		if err := rows.Scan(
			&i.ID,
			&i.ShopCategoryID,
			&i.Type,
			&i.EnumID,
			&i.Info,
			&i.Name,
			&i.Desc,
			&i.CostType,
			&i.CostEnumID,
			&i.Cost,
			&i.OriginalCost,
			&i.Count,
			&i.Visible,
			&i.StartAt,
			&i.EndAt,
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
