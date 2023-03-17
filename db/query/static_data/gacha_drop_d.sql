-- name: GetGachaDropDList :many
SELECT * FROM `gacha_drop_d`;

-- name: GetGachaDropDByEnumId :one
SELECT * FROM `gacha_drop_d` WHERE enum_id = ?;
