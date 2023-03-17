-- name: GetCharacterGetDialogues :many
SELECT * FROM character_get_dialogue;

-- name: GetCharacterGetDialogueByEnumId :one
SELECT * FROM character_get_dialogue WHERE enum_id = ?;
