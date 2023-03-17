package db_game

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const getCharacterListById = `-- name: GetCharacterListById :many
SELECT id, account_user_id, enum_id, exp, equipment_level, created_at FROM ` + "`" + `character` + "`" + ` WHERE id IN (%s)
`

func (q *Queries) GetCharacterListById(ctx context.Context, charactersId []int32) ([]Character, error) {
	ids := []string{}
	for _, id := range charactersId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(getCharacterListById, strings.Join(ids, ","))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Character{}
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.EnumID,
			&i.Exp,
			&i.EquipmentLevel,
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

const getCharacterListByIdAndAccountUserId = `-- name: GetCharacterListByIdAndAccountUserId :many
SELECT id, account_user_id, enum_id, exp, equipment_level, created_at FROM ` + "`" + `character` + "`" + ` WHERE id IN (%s) AND account_user_id = %d
`

func (q *Queries) GetCharacterListByIdAndAccountUserId(ctx context.Context, charactersId []int32, accountUserId int32) ([]Character, error) {
	ids := []string{}
	for _, id := range charactersId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(getCharacterListByIdAndAccountUserId, strings.Join(ids, ","), accountUserId)
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Character{}
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.AccountUserID,
			&i.EnumID,
			&i.Exp,
			&i.EquipmentLevel,
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

const deleteCharacters = `-- name: DeleteCharacters :execresult
DELETE FROM ` + "`" + `character` + "`" + ` WHERE id IN (%s)
`

func (q *Queries) DeleteCharacters(ctx context.Context, charactersId []int32) (sql.Result, error) {
	ids := []string{}
	for _, id := range charactersId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(deleteCharacters, strings.Join(ids, ","))

	return q.db.ExecContext(ctx, query)
}
