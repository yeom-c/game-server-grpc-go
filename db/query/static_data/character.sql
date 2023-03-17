-- name: GetCharacters :many
SELECT * FROM `character`;

-- name: GetCharacterByEnumId :one
SELECT * FROM `character` WHERE enum_id = ?;
