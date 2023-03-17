-- name: GetShopGoodsListByShopCategoryId :many
SELECT * FROM shop_goods WHERE shop_category_id = ?;

-- name: GetShopGoods :one
SELECT * FROM shop_goods WHERE id = ?;
