package db_game

import (
	"context"
	"fmt"
	"strings"
)

const getAssetListByEnumId = `-- name: GetAssetListByEnumId :many
SELECT id, account_user_id, enum_id, type, balance FROM asset WHERE account_user_id = %d AND enum_id IN ("%s")
`

func (q *Queries) GetAssetListByEnumId(ctx context.Context, accountUserId int32, assetsEnumId []string) ([]Asset, error) {
	query := fmt.Sprintf(getAssetListByEnumId, accountUserId, strings.Join(assetsEnumId, "\",\""))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Asset{}
	for rows.Next() {
		var i Asset
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.EnumID,
			&i.Type,
			&i.Balance,
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
