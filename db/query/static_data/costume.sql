-- name: GetCostumes :many
SELECT * FROM `costume`;

-- name: GetCostumeByEnumId :one
SELECT * FROM `costume` WHERE enum_id = ?;
