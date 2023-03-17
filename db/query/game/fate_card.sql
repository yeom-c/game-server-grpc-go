-- name: GetFateCardListByAccountUserId :many
SELECT * FROM fate_card WHERE account_user_id = ?;

-- name: GetFateCard :one
SELECT * FROM fate_card WHERE id = ?;

-- name: GetFateCardByEnumId :one
SELECT * FROM fate_card WHERE account_user_id = ? AND enum_id = ?;

-- name: GetFateCardByCharacterEnumId :one
SELECT * FROM fate_card WHERE account_user_id = ? AND character_enum_id = ? LIMIT 1;

-- name: CreateFateCard :execresult
INSERT INTO fate_card (account_user_id, enum_id) VALUES (?, ?);

-- name: UpdateFateCardCharacterEnumId :execresult
UPDATE fate_card SET character_enum_id = ? WHERE id = ?;

-- name: UnequipFateCardByCharacterEnumId :execresult
UPDATE fate_card SET character_enum_id = NULL WHERE account_user_id = ? AND character_enum_id = ?;
