-- name: GetDrops :many
SELECT * FROM `drop`;

-- name: GetDropByEnumId :one
SELECT * FROM `drop` WHERE enum_id = ?;
