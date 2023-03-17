-- name: GetAccountGrowths :many
SELECT * FROM account_growth;

-- name: GetAccountGrowthByEnumId :one
SELECT * FROM account_growth WHERE enum_id = ?;
