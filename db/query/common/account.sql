-- name: GetAccount :one
SELECT * FROM account WHERE id = ?;

-- name: GetAccountByUuid :one
SELECT * FROM account WHERE uuid = ? limit 1;

-- name: CreateAccount :execresult
INSERT INTO account (uuid, world_id, profile_idx) VALUES (sqlc.arg(uuid), sqlc.arg(world_id), sqlc.arg(profile_idx));
