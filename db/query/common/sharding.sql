-- name: GetShardings :many
SELECT * FROM sharding;

-- name: GetGameDb :one
SELECT * FROM sharding ORDER BY `count` ASC LIMIT 1;

-- name: UpdateShardingCountByGameDb :execresult
UPDATE sharding SET `count` = `count` + ? WHERE game_db = ?;
