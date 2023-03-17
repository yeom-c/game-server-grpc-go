-- name: GetRecipes :many
SELECT * FROM recipe;

-- name: GetRecipeByEnumId :one
SELECT * FROM recipe WHERE enum_id = ?;
