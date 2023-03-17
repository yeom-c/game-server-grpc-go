package db_static_data

import (
	"context"
	"fmt"
	"strings"
)

const getGachaPoolGroupListByEnumId = `-- name: GetGachaPoolGroupListByEnumId :many
SELECT id, enum_id, pool_condition, pool_id, pool_id_rate, created_at FROM ` + "`" + `gacha_pool_group` + "`" + ` WHERE enum_id IN ("%s")
`

func (q *Queries) GetGachaPoolGroupListByEnumId(ctx context.Context, enumId []string) ([]GachaPoolGroup, error) {
	query := fmt.Sprintf(getGachaPoolGroupListByEnumId, strings.Join(enumId, "\",\""))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GachaPoolGroup{}
	for rows.Next() {
		var i GachaPoolGroup
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.PoolCondition,
			&i.PoolID,
			&i.PoolIDRate,
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
