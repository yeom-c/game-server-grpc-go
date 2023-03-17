-- name: GetGachaLogCategoryListByAccountUserId :many
SELECT enum_id FROM gacha_log where account_user_id = ? GROUP BY enum_id ORDER BY enum_id;

-- name: GetGachaLogListByAccountUserId :many
SELECT * FROM gacha_log WHERE account_user_id = ? ORDER BY created_at DESC, id DESC LIMIT ?, ?;

-- name: GetGachaLogListByEnumId :many
SELECT * FROM gacha_log WHERE account_user_id = ? AND enum_id = ? ORDER BY created_at DESC, id DESC LIMIT ?, ?;

-- name: GetGachaLogListCountByAccountUserId :one
SELECT COUNT(*) FROM gacha_log WHERE account_user_id = ?;

-- name: GetGachaLogListCountByEnumId :one
SELECT COUNT(*) FROM gacha_log WHERE account_user_id = ? AND enum_id = ?;

-- name: CreateGachaLog :execresult
INSERT INTO gacha_log (account_user_id, enum_id, character_enum_id) VALUES (?, ?, ?);
