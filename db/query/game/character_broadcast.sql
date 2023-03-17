-- name: GetAllCompletedCharacterBroadcastList :many
SELECT * FROM character_broadcast WHERE account_user_id = ? AND complete = 1;

-- name: GetCompletedCharacterBroadcastList :many
SELECT * FROM character_broadcast WHERE account_user_id = ? AND character_enum_id = ? AND complete = 1;

-- name: GetOnAirCharacterBroadcastList :many
SELECT * FROM character_broadcast WHERE account_user_id = ? AND on_air = 1;

-- name: GetOnAirCharacterBroadcast :one
SELECT * FROM character_broadcast WHERE id = ?;

-- name: CreateCharacterBroadcast :execresult
INSERT INTO character_broadcast (account_user_id, character_enum_id, timeline_enum_id, `type`, on_air, complete) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpsertCharacterBroadcast :execresult
INSERT INTO character_broadcast (account_user_id, character_enum_id, timeline_enum_id, `type`, on_air, complete) VALUES (?, ?, ?, ?, sqlc.arg(on_air), sqlc.arg(complete)) ON DUPLICATE KEY UPDATE on_air = sqlc.arg(on_air), complete = sqlc.arg(complete), broadcasted_at = NOW();

-- name: UpdateCharacterBroadcastComplete :execresult
UPDATE character_broadcast SET complete = ? WHERE id = ?;

-- name: UpdateCharacterBroadcastOnAirByAccountUserId :execresult
UPDATE character_broadcast SET on_air = ? WHERE account_user_id = ?;
