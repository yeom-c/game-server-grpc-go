-- name: GetItems :many
SELECT * FROM item;

-- name: GetItemByEnumId :one
SELECT * FROM item WHERE enum_id = ?;
