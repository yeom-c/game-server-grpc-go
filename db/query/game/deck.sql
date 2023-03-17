-- name: GetDeckListByAccountUserId :many
SELECT * FROM deck WHERE account_user_id = ?;

-- name: GetDeckCount :one
SELECT COUNT(*) FROM deck WHERE account_user_id = ?;

-- name: GetDeck :one
SELECT * FROM deck WHERE id = ?;

-- name: CreateDeck :execresult
INSERT INTO deck (account_user_id, `index`, `name`, character_id_0, character_id_1, character_id_2, character_id_3, character_id_4) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateDeck :execresult
UPDATE deck SET `name`= ?, character_id_0 = ?, character_id_1 = ?, character_id_2 = ?, character_id_3 = ?, character_id_4 = ? WHERE id = ?;
