-- name: GetGachaDropBList :many
SELECT * FROM `gacha_drop_b`;

-- name: GetGachaDropBByEnumId :one
SELECT * FROM `gacha_drop_b` WHERE enum_id = ?;
