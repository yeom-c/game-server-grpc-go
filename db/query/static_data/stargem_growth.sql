-- name: GetStargemGrowths :many
SELECT * FROM stargem_growth;

-- name: GetStargemGrowthByEnumId :one
SELECT * FROM stargem_growth WHERE enum_id = ?;
