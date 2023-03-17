-- name: GetCostumeListByAccountUserId :many
SELECT * FROM costume WHERE account_user_id = ?;

-- name: GetCostumeListByCharacterEnumId :many
SELECT * FROM costume WHERE account_user_id = ? AND character_enum_id = ?;

-- name: GetCostume :one
SELECT * FROM costume WHERE id = ?;

-- name: GetCostumeByEnumId :one
SELECT * FROM costume WHERE account_user_id = ? AND enum_id = ?;

-- name: CreateCostume :execresult
INSERT IGNORE INTO costume (account_user_id, enum_id, character_enum_id, `state`) VALUES (?, ?, ?, ?);

-- name: UpdateCostumeState :execresult
UPDATE costume SET `state` = ? WHERE id = ?;

-- name: UpdateCostumeStateByCharacterEnumId :execresult
UPDATE costume SET `state` = ? WHERE account_user_id = ? AND character_enum_id = ?;
