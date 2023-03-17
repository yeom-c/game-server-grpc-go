-- name: GetCharacterCollectionListByAccountUserId :many
SELECT * FROM character_collection WHERE account_user_id = ?;

-- name: GetCharacterCollectionCountByAccountUserId :one
SELECT COUNT(*) FROM character_collection WHERE account_user_id = ?;

-- name: UpsertCharacterCollection :execresult
INSERT INTO character_collection (account_user_id, character_enum_id, affection_exp, `count`) VALUES (?, ?, sqlc.arg(affection_exp), sqlc.arg(count)) ON DUPLICATE KEY UPDATE affection_exp = affection_exp + sqlc.arg(affection_exp), `count` = `count` + sqlc.arg(count);
