// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: tutorial_group.sql

package db_static_data

import (
	"context"
)

const getTutorialGroupByEnumId = `-- name: GetTutorialGroupByEnumId :one
SELECT id, enum_id, next_group, group_reward, group_contents, created_at FROM tutorial_group WHERE enum_id = ?
`

func (q *Queries) GetTutorialGroupByEnumId(ctx context.Context, enumID string) (TutorialGroup, error) {
	row := q.db.QueryRowContext(ctx, getTutorialGroupByEnumId, enumID)
	var i TutorialGroup
	err := row.Scan(
		&i.ID,
		&i.EnumID,
		&i.NextGroup,
		&i.GroupReward,
		&i.GroupContents,
		&i.CreatedAt,
	)
	return i, err
}

const getTutorialGroups = `-- name: GetTutorialGroups :many
SELECT id, enum_id, next_group, group_reward, group_contents, created_at FROM tutorial_group
`

func (q *Queries) GetTutorialGroups(ctx context.Context) ([]TutorialGroup, error) {
	rows, err := q.db.QueryContext(ctx, getTutorialGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TutorialGroup{}
	for rows.Next() {
		var i TutorialGroup
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.NextGroup,
			&i.GroupReward,
			&i.GroupContents,
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
