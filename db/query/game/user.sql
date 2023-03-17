-- name: GetUser :one
SELECT * FROM `user` WHERE id = ?;

-- name: GetUserByAccountUserId :one
SELECT * FROM `user` WHERE account_user_id = ? LIMIT 1;

-- name: GetUserTutorialInfoByAccountUserId :one
SELECT tutorial_info FROM `user` WHERE account_user_id = ?;

-- name: GetUserShopInfoByAccountUserId :one
SELECT shop_info FROM `user` WHERE account_user_id = ?;

-- name: CreateUser :execresult
INSERT INTO `user` (`account_user_id`) VALUES (?);

-- name: UpdateUserTutorialInfoByAccountUserId :execresult
UPDATE `user` SET tutorial_info = ? WHERE account_user_id = ?;

-- name: UpdateUserShopInfoByAccountUserId :execresult
UPDATE `user` SET shop_info = ? WHERE account_user_id = ?;

-- name: UpdateUserDailyResetAtByAccountUserId :execresult
UPDATE `user` SET daily_reset_at = ? WHERE account_user_id = ?;

-- name: UpdateUserBroadcastResetAtByAccountUserId :execresult
UPDATE `user` SET broadcast_reset_at = ? WHERE account_user_id = ?;

-- name: UpdateUserStoryIndexByAccountUserId :execresult
UPDATE `user` SET story_index = ? WHERE account_user_id = ?;
