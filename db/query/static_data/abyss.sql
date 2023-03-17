-- name: GetAbysses :many
SELECT * FROM abyss;

-- name: GetAbyssByEnumId :one
SELECT * FROM abyss WHERE enum_id = ?;
