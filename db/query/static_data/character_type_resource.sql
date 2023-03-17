-- name: GetCharacterTypeResources :many
SELECT * FROM character_type_resource;

-- name: GetCharacterTypeResourceByEnumId :one
SELECT * FROM character_type_resource WHERE enum_id = ?;
