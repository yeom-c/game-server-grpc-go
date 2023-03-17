-- name: GetMailListByAccountUserId :many
SELECT * FROM mail WHERE account_user_id = ? ORDER BY `status` ASC, `created_at` DESC, id desc LIMIT ?;

-- name: GetMail :one
SELECT * FROM mail WHERE id = ?;

-- name: CreateMail :execresult
INSERT INTO mail (account_user_id, sender, `type`, `status`, delete_all, attachment, title, message, expired_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateMailStatus :execresult
UPDATE mail SET `status` = ? WHERE id = ?;

-- name: DeleteConfirmMail :execresult
DELETE FROM mail WHERE account_user_id = ? AND id = ? AND `status` = ?;

-- name: DeleteAllConfirmMail :execresult
DELETE FROM mail WHERE account_user_id = ? AND `status` = ? AND delete_all = ?;

-- name: DeleteAllExpiredMail :execresult
DELETE FROM mail WHERE account_user_id = ? AND expired_at IS NOT NULL AND expired_at < NOW();
