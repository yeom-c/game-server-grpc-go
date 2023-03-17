-- name: GetGachaDropCList :many
SELECT * FROM `gacha_drop_c`;

-- name: GetGachaDropCByEnumId :one
SELECT * FROM `gacha_drop_c` WHERE enum_id = ?;
