-- name: GetGachas :many
SELECT * FROM `gacha`;

-- name: GetGachaByEnumId :one
SELECT * FROM `gacha` WHERE enum_id = ?;
