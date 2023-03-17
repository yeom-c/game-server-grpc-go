-- name: GetDialogues :many
SELECT * FROM dialogue;

-- name: GetDialogueByEnumId :one
SELECT * FROM dialogue WHERE enum_id = ?;
