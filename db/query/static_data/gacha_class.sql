-- name: GetGachaClasses :many
SELECT * FROM `gacha_class`;

-- name: GetGachaClassByEnumId :one
SELECT * FROM `gacha_class` WHERE enum_id = ?;
