-- name: GetCharacterResources :many
SELECT * FROM character_resource;

-- name: GetCharacterResourceByEnumId :one
SELECT * FROM character_resource WHERE enum_id = ?;
