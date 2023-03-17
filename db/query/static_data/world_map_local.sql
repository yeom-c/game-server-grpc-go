-- name: GetWorldMapLocals :many
SELECT * FROM world_map_local;

-- name: GetWorldMapLocalByEnumId :one
SELECT * FROM world_map_local WHERE enum_id = ?;
