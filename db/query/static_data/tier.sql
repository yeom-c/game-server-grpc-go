-- name: GetTiers :many
SELECT * FROM tier;

-- name: GetTierByEnumId :one
SELECT * FROM tier WHERE enum_id = ?;
