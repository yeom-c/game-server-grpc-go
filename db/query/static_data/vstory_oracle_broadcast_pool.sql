-- name: GetVstoryOracleBroadcastPool :many
SELECT * FROM vstory_oracle_broadcast_pool;

-- name: GetVstoryOracleBroadcastPoolByEnumId :one
SELECT * FROM vstory_oracle_broadcast_pool WHERE enum_id = ?;
