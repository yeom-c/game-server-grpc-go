-- name: GetCharacterGrowths :many
SELECT * FROM character_growth;

-- name: GetCharacterGrowthByEnumId :one
SELECT * FROM character_growth WHERE enum_id = ?;
