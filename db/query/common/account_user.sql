-- name: GetAccountUser :one
SELECT * FROM account_user WHERE id = ?;

-- name: GetAccountUserByAccountId :one
SELECT * FROM account_user WHERE account_id = ? limit 1;

-- name: GetAccountUserByNickname :one
SELECT * FROM account_user WHERE nickname = ? limit 1;

-- name: CreateAccountUser :execresult
INSERT INTO account_user (account_id, game_db, nickname) VALUES (?, ?, ?);

-- name: UpdateAccountUserSignedIn :execresult
UPDATE account_user SET signed_in_at = ? WHERE id = ?;
