// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: account.sql

package db_common

import (
	"context"
	"database/sql"
)

const createAccount = `-- name: CreateAccount :execresult
INSERT INTO account (uuid, world_id, profile_idx) VALUES (?, ?, ?)
`

type CreateAccountParams struct {
	Uuid       string        `json:"uuid"`
	WorldID    int32         `json:"world_id"`
	ProfileIdx sql.NullInt32 `json:"profile_idx"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAccount, arg.Uuid, arg.WorldID, arg.ProfileIdx)
}

const getAccount = `-- name: GetAccount :one
SELECT id, uuid, world_id, profile_idx, created_at FROM account WHERE id = ?
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.WorldID,
		&i.ProfileIdx,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountByUuid = `-- name: GetAccountByUuid :one
SELECT id, uuid, world_id, profile_idx, created_at FROM account WHERE uuid = ? limit 1
`

func (q *Queries) GetAccountByUuid(ctx context.Context, uuid string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUuid, uuid)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.WorldID,
		&i.ProfileIdx,
		&i.CreatedAt,
	)
	return i, err
}
