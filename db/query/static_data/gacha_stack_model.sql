-- name: GetGachaStackModelList :many
SELECT * FROM `gacha_stack_model`;

-- name: GetGachaStackModelByEnumId :one
SELECT * FROM `gacha_stack_model` WHERE enum_id = ?;
