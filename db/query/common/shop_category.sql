-- name: GetShopCategoryListByShopId :many
SELECT * FROM shop_category WHERE shop_id = ? ORDER BY `order` ASC, id DESC;
