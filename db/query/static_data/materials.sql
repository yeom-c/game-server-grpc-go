-- name: GetMaterials :many
SELECT * FROM materials;

-- name: GetMaterialsByEnumId :one
SELECT * FROM materials WHERE enum_id = ?;
