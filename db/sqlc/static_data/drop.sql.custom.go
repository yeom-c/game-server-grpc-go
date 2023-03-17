package db_static_data

import (
	"context"
	"fmt"
	"strings"
)

const getDropListByEnumId = `-- name: GetDropListByEnumId :many
SELECT id, enum_id, content_type, ce_common_type_drop, icon, ` + "`" + `drop` + "`" + `, rate, value, drop_next, created_at FROM ` + "`" + `drop` + "`" + ` WHERE enum_id IN ("%s")
`

func (q *Queries) GetDropListByEnumId(ctx context.Context, enumId []string) ([]Drop, error) {
	query := fmt.Sprintf(getDropListByEnumId, strings.Join(enumId, "\",\""))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Drop{}
	for rows.Next() {
		var i Drop
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.ContentType,
			&i.CeCommonTypeDrop,
			&i.Icon,
			&i.Drop,
			&i.Rate,
			&i.Value,
			&i.DropNext,
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
