-- name: GetCharacterBaseStats :many
SELECT * FROM character_base_stat;

-- name: GetCharacterBaseStatByEnumId :one
SELECT * FROM character_base_stat WHERE enum_id = ?;
