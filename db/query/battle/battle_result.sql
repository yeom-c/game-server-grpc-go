-- name: GetUnconfirmedBattleResults :many
SELECT * FROM battle_result WHERE account_user_id = ? AND confirmed_at IS NULL;

-- name: GetBattleResultByChannelId :one
SELECT * FROM battle_result WHERE account_user_id = ? AND channel_id = ?;

-- name: CreateBattleResult :execresult
INSERT INTO battle_result (account_user_id, match_account_user_id, channel_id, deck_id, battle_start_at) VALUES (?, ?, ?, ?, ?);

-- name: UpdateBattleResultResultByChannelId :execresult
UPDATE battle_result SET result = ? WHERE account_user_id = ? AND channel_id = ?;

-- name: UpdateBattleResultConfirmedAt :execresult
UPDATE battle_result SET confirmed_at = ? WHERE id = ?;
