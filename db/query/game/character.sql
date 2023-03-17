-- name: GetCharacterListByAccountUserId :many
SELECT * FROM `character` WHERE account_user_id = ?;

-- name: GetCharacter :one
SELECT * FROM `character` WHERE id = ?;

-- name: GetCharacterByEnumId :one
SELECT * FROM `character` WHERE account_user_id = ? AND enum_id = ? LIMIT 1;

-- name: CreateCharacter :execresult
INSERT INTO `character` (account_user_id, enum_id, exp) VALUES (?, ?, ?);

-- name: UpdateCharacterExp :execresult
UPDATE `character` SET exp = exp + ? WHERE id = ?;

-- name: UpdateCharacterExpLimitMax :execresult
UPDATE `character` SET exp = IF(exp >= sqlc.arg(max_exp), exp, IF(exp + sqlc.arg(add_exp) >= sqlc.arg(max_exp), sqlc.arg(max_exp), exp + sqlc.arg(add_exp))) WHERE id = ?;

-- name: UpdateCharacterEquipmentLevel :execresult
UPDATE `character` SET equipment_level = equipment_level + sqlc.arg(add_level) WHERE id = ?;

-- name: DeleteCharacter :execresult
DELETE FROM `character` WHERE id = ?;
