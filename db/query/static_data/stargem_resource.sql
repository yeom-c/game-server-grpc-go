-- name: GetStargemResources :many
SELECT * FROM stargem_resource;

-- name: GetStargemResourceByEnumId :one
SELECT * FROM stargem_resource WHERE enum_id = ?;
