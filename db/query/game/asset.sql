-- name: GetAssetListByAccountUserId :many
SELECT * FROM asset WHERE account_user_id = ?;

-- name: GetAssetByEnumId :one
SELECT * FROM asset WHERE account_user_id = ? AND enum_id = ?;

-- name: UpsertAsset :execresult
INSERT INTO asset (account_user_id, enum_id, type, balance) VALUES (?, ?, ?, sqlc.arg(amount)) ON DUPLICATE KEY UPDATE balance = balance + sqlc.arg(amount);

-- name: UpsertAssetLimitMax :execresult
INSERT INTO asset (account_user_id, enum_id, type, balance) VALUES (?, ?, ?, sqlc.arg(amount)) ON DUPLICATE KEY UPDATE balance = IF(balance >= sqlc.arg(max_amount), balance, IF(balance + sqlc.arg(amount) >= sqlc.arg(max_amount), sqlc.arg(max_amount), balance + sqlc.arg(amount)));

-- name: AddAssetBalance :execresult
UPDATE asset SET balance = balance + sqlc.arg(amount) WHERE id = ?;

-- name: AddAssetBalanceByEnumId :execresult
UPDATE asset SET balance = balance + sqlc.arg(amount) WHERE account_user_id = ? AND enum_id = ?;

-- name: AddAssetBalanceByEnumType :execresult
UPDATE asset SET balance = balance + sqlc.arg(amount) WHERE account_user_id = ? AND type = ?;
