-- name: GetCharacterCostumeBundleList :many
SELECT * FROM `character_costume_bundle`;

-- name: GetCharacterCostumeBundleByEnumId :one
SELECT * FROM `character_costume_bundle` WHERE enum_id = ?;
