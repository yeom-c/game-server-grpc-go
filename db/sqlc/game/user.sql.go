// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db_game

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO ` + "`" + `user` + "`" + ` (` + "`" + `account_user_id` + "`" + `) VALUES (?)
`

func (q *Queries) CreateUser(ctx context.Context, accountUserID int32) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser, accountUserID)
}

const getUser = `-- name: GetUser :one
SELECT id, account_user_id, story_index, tutorial_info, shop_info, daily_reset_at, broadcast_reset_at, created_at FROM ` + "`" + `user` + "`" + ` WHERE id = ?
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountUserID,
		&i.StoryIndex,
		&i.TutorialInfo,
		&i.ShopInfo,
		&i.DailyResetAt,
		&i.BroadcastResetAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByAccountUserId = `-- name: GetUserByAccountUserId :one
SELECT id, account_user_id, story_index, tutorial_info, shop_info, daily_reset_at, broadcast_reset_at, created_at FROM ` + "`" + `user` + "`" + ` WHERE account_user_id = ? LIMIT 1
`

func (q *Queries) GetUserByAccountUserId(ctx context.Context, accountUserID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByAccountUserId, accountUserID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.AccountUserID,
		&i.StoryIndex,
		&i.TutorialInfo,
		&i.ShopInfo,
		&i.DailyResetAt,
		&i.BroadcastResetAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUserShopInfoByAccountUserId = `-- name: GetUserShopInfoByAccountUserId :one
SELECT shop_info FROM ` + "`" + `user` + "`" + ` WHERE account_user_id = ?
`

func (q *Queries) GetUserShopInfoByAccountUserId(ctx context.Context, accountUserID int32) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getUserShopInfoByAccountUserId, accountUserID)
	var shop_info sql.NullString
	err := row.Scan(&shop_info)
	return shop_info, err
}

const getUserTutorialInfoByAccountUserId = `-- name: GetUserTutorialInfoByAccountUserId :one
SELECT tutorial_info FROM ` + "`" + `user` + "`" + ` WHERE account_user_id = ?
`

func (q *Queries) GetUserTutorialInfoByAccountUserId(ctx context.Context, accountUserID int32) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserTutorialInfoByAccountUserId, accountUserID)
	var tutorial_info string
	err := row.Scan(&tutorial_info)
	return tutorial_info, err
}

const updateUserBroadcastResetAtByAccountUserId = `-- name: UpdateUserBroadcastResetAtByAccountUserId :execresult
UPDATE ` + "`" + `user` + "`" + ` SET broadcast_reset_at = ? WHERE account_user_id = ?
`

type UpdateUserBroadcastResetAtByAccountUserIdParams struct {
	BroadcastResetAt time.Time `json:"broadcast_reset_at"`
	AccountUserID    int32     `json:"account_user_id"`
}

func (q *Queries) UpdateUserBroadcastResetAtByAccountUserId(ctx context.Context, arg UpdateUserBroadcastResetAtByAccountUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserBroadcastResetAtByAccountUserId, arg.BroadcastResetAt, arg.AccountUserID)
}

const updateUserDailyResetAtByAccountUserId = `-- name: UpdateUserDailyResetAtByAccountUserId :execresult
UPDATE ` + "`" + `user` + "`" + ` SET daily_reset_at = ? WHERE account_user_id = ?
`

type UpdateUserDailyResetAtByAccountUserIdParams struct {
	DailyResetAt  time.Time `json:"daily_reset_at"`
	AccountUserID int32     `json:"account_user_id"`
}

func (q *Queries) UpdateUserDailyResetAtByAccountUserId(ctx context.Context, arg UpdateUserDailyResetAtByAccountUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserDailyResetAtByAccountUserId, arg.DailyResetAt, arg.AccountUserID)
}

const updateUserShopInfoByAccountUserId = `-- name: UpdateUserShopInfoByAccountUserId :execresult
UPDATE ` + "`" + `user` + "`" + ` SET shop_info = ? WHERE account_user_id = ?
`

type UpdateUserShopInfoByAccountUserIdParams struct {
	ShopInfo      sql.NullString `json:"shop_info"`
	AccountUserID int32          `json:"account_user_id"`
}

func (q *Queries) UpdateUserShopInfoByAccountUserId(ctx context.Context, arg UpdateUserShopInfoByAccountUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserShopInfoByAccountUserId, arg.ShopInfo, arg.AccountUserID)
}

const updateUserStoryIndexByAccountUserId = `-- name: UpdateUserStoryIndexByAccountUserId :execresult
UPDATE ` + "`" + `user` + "`" + ` SET story_index = ? WHERE account_user_id = ?
`

type UpdateUserStoryIndexByAccountUserIdParams struct {
	StoryIndex    int32 `json:"story_index"`
	AccountUserID int32 `json:"account_user_id"`
}

func (q *Queries) UpdateUserStoryIndexByAccountUserId(ctx context.Context, arg UpdateUserStoryIndexByAccountUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserStoryIndexByAccountUserId, arg.StoryIndex, arg.AccountUserID)
}

const updateUserTutorialInfoByAccountUserId = `-- name: UpdateUserTutorialInfoByAccountUserId :execresult
UPDATE ` + "`" + `user` + "`" + ` SET tutorial_info = ? WHERE account_user_id = ?
`

type UpdateUserTutorialInfoByAccountUserIdParams struct {
	TutorialInfo  string `json:"tutorial_info"`
	AccountUserID int32  `json:"account_user_id"`
}

func (q *Queries) UpdateUserTutorialInfoByAccountUserId(ctx context.Context, arg UpdateUserTutorialInfoByAccountUserIdParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUserTutorialInfoByAccountUserId, arg.TutorialInfo, arg.AccountUserID)
}