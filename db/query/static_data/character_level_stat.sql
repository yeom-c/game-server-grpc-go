-- name: GetCharacterLevelStats :many
SELECT * FROM character_level_stat;

-- name: GetCharacterLevelStatByEnumId :one
SELECT * FROM character_level_stat WHERE enum_id = ?;
