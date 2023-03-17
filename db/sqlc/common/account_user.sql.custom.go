package db_common

import (
	"context"
	"fmt"
	"strings"
)

const getAccountUserListById = `-- name: GetAccountUserListById :many
SELECT id, account_id, game_db, nickname, created_at FROM account_user WHERE id IN (%s)
`

func (q *Queries) GetAccountUserListById(ctx context.Context, accountUsersId []int32) ([]AccountUser, error) {
	ids := []string{}
	for _, id := range accountUsersId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(getAccountUserListById, strings.Join(ids, ","))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AccountUser{}
	for rows.Next() {
		var i AccountUser
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.GameDb,
			&i.Nickname,
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
