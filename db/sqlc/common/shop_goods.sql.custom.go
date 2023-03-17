package db_common

import (
	"context"
	"fmt"
	"strings"
)

const getShopGoodsListByShopCategoryIdList = `-- name: GetShopGoodsListByShopCategoryIdList :many
SELECT id, shop_category_id, type, enum_id, info, name, ` + "`" + `desc` + "`" + `, cost_type, cost_enum_id, cost, original_cost, count, visible, start_at, end_at FROM shop_goods WHERE shop_category_id IN (%s)
`

func (q *Queries) GetShopGoodsListByShopCategoryIdList(ctx context.Context, shopCategoryID []int32) ([]ShopGoods, error) {
	ids := []string{}
	for _, id := range shopCategoryID {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(getShopGoodsListByShopCategoryIdList, strings.Join(ids, ","))
	rows, err := q.db.QueryContext(ctx, query)
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
