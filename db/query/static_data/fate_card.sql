-- name: GetFateCards :many
SELECT * FROM `fate_card`;

-- name: GetFateCardByEnumId :one
SELECT * FROM `fate_card` WHERE enum_id = ?;
