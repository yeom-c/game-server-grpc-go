-- name: GetItemListByAccountUserId :many
SELECT * FROM item WHERE account_user_id = ?;

-- name: GetItem :one
SELECT * FROM item WHERE id = ?;

-- name: GetItemByEnumId :one
SELECT * FROM item WHERE account_user_id = ? and enum_id = ?;

-- name: UpsertItem :execresult
INSERT INTO item (account_user_id, enum_id, `count`) VALUES (?, ?, sqlc.arg(count)) ON DUPLICATE KEY UPDATE `count` = `count` + sqlc.arg(count);

-- name: AddItemCount :execresult
UPDATE item SET `count` = `count` + sqlc.arg(amount) WHERE id = ?;
