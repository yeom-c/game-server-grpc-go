-- name: GetBalances :many
SELECT * FROM balance;

-- name: GetBalanceByEnumId :one
SELECT * FROM balance WHERE enum_id = ?;
