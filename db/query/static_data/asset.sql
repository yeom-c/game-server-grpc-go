-- name: GetAssets :many
SELECT * FROM asset;

-- name: GetAssetByEnumId :one
SELECT * FROM asset WHERE enum_id = ?;
