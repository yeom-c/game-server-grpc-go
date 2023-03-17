-- name: GetStoryStages :many
SELECT * FROM story_stage;

-- name: GetStoryStageByEnumId :one
SELECT * FROM story_stage WHERE enum_id = ?;
