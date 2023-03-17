-- name: GetRankerList :many
SELECT * FROM `user` WHERE match_point > 0 ORDER BY match_point DESC, id ASC LIMIT ?, ?;

-- name: GetUserByAccountUserId :one
SELECT * FROM `user` WHERE account_user_id = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO `user` (`account_user_id`) VALUES (?);

-- name: UpdateUserMatchResultByAccountUserId :execresult
UPDATE `user` SET match_point = match_point + sqlc.arg(add_match_point), match_win = match_win + sqlc.arg(add_match_win), match_lose = match_lose + sqlc.arg(add_match_lose) WHERE account_user_id = ?;
