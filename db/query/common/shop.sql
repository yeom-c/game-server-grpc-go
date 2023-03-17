-- name: GetShopByTypeAndVisible :one
SELECT * FROM shop WHERE `type` = sqlc.arg(shop_type) AND `visible` = sqlc.arg(shop_visible) ORDER BY id DESC LIMIT 1;
