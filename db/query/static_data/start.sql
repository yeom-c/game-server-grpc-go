-- name: GetStarts :many
SELECT * FROM `start`;

-- name: GetStartByEnumId :one
SELECT * FROM start WHERE enum_id = ?;
