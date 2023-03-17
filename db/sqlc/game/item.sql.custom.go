package db_game

import (
	"context"
	"fmt"
	"strings"
)

const getItemListById = `-- name: GetItemListById :many
SELECT id, account_user_id, enum_id, count, created_at FROM item WHERE id IN (%s)
`

func (q *Queries) GetItemListById(ctx context.Context, itemsId []int32) ([]Item, error) {
	ids := []string{}
	for _, id := range itemsId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(getItemListById, strings.Join(ids, ","))
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
			&i.AccountUserID,
			&i.EnumID,
			&i.Count,
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

const getItemsByEnumId = `-- name: GetItemsByEnumId :many
SELECT id, account_user_id, enum_id, count, created_at FROM item WHERE account_user_id = %d AND enum_id IN ("%s")
`

func (q *Queries) GetItemsByEnumId(ctx context.Context, accountUserId int32, itemsEnumId []string) ([]Item, error) {
	query := fmt.Sprintf(getItemsByEnumId, accountUserId, strings.Join(itemsEnumId, "\",\""))
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
			&i.AccountUserID,
			&i.EnumID,
			&i.Count,
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
