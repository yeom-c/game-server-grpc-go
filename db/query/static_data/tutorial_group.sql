-- name: GetTutorialGroups :many
SELECT * FROM tutorial_group;

-- name: GetTutorialGroupByEnumId :one
SELECT * FROM tutorial_group WHERE enum_id = ?;
