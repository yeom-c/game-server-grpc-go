-- name: GetGachaPoolGroupList :many
SELECT * FROM `gacha_pool_group`;

-- name: GetGachaPoolGroupByEnumId :one
SELECT * FROM `gacha_pool_group` WHERE enum_id = ?;
