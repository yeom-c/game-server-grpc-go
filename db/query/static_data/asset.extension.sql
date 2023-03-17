-- name: GetAssetByAssetEnum :one
SELECT * FROM asset WHERE ce_asset = ? LIMIT 1;
