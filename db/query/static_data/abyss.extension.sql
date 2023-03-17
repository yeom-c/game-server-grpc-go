-- name: GetAbyssByLevel :one
SELECT * FROM abyss WHERE `level` = ? LIMIT 1;
