// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: character_get_dialogue.sql

package db_static_data

import (
	"context"
)

const getCharacterGetDialogueByEnumId = `-- name: GetCharacterGetDialogueByEnumId :one
SELECT id, enum_id, character_enum_id, animation, day_dialogue, night_dialogue, day_voice, night_voice, created_at FROM character_get_dialogue WHERE enum_id = ?
`

func (q *Queries) GetCharacterGetDialogueByEnumId(ctx context.Context, enumID string) (CharacterGetDialogue, error) {
	row := q.db.QueryRowContext(ctx, getCharacterGetDialogueByEnumId, enumID)
	var i CharacterGetDialogue
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.CharacterEnumID,
		&i.Animation,
		&i.DayDialogue,
		&i.NightDialogue,
		&i.DayVoice,
		&i.NightVoice,
		&i.CreatedAt,
	)
	return i, err
}

const getCharacterGetDialogues = `-- name: GetCharacterGetDialogues :many
SELECT id, enum_id, character_enum_id, animation, day_dialogue, night_dialogue, day_voice, night_voice, created_at FROM character_get_dialogue
`

func (q *Queries) GetCharacterGetDialogues(ctx context.Context) ([]CharacterGetDialogue, error) {
	rows, err := q.db.QueryContext(ctx, getCharacterGetDialogues)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterGetDialogue{}
	for rows.Next() {
		var i CharacterGetDialogue
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.CharacterEnumID,
			&i.Animation,
			&i.DayDialogue,
			&i.NightDialogue,
			&i.DayVoice,
			&i.NightVoice,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
