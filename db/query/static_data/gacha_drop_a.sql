-- name: GetGachaDropAList :many
SELECT * FROM `gacha_drop_a`;

-- name: GetGachaDropAByEnumId :one
SELECT * FROM `gacha_drop_a` WHERE enum_id = ?;
