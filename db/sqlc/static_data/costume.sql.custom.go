package db_static_data

import (
	"context"
	"fmt"
	"strings"
)

const getCostumesByConditionAndCharacterEnumIds = `-- name: GetCostumesByConditionAndCharacterEnumIds :many
SELECT id, enum_id, skin_name, ` + "`" + `character` + "`" + `, ce_costume_condition, condition_value, illust_reference, portrait_reference, voice_appear, created_at FROM ` + "`" + `costume` + "`" + ` WHERE ce_costume_condition = "%s" AND ` + "`" + `character` + "`" + ` IN ("%s")
`

func (q *Queries) GetCostumesByConditionAndCharacterEnumIds(ctx context.Context, condition string, characterEnumIDs []string) ([]Costume, error) {
	query := fmt.Sprintf(getCostumesByConditionAndCharacterEnumIds, condition, strings.Join(characterEnumIDs, "\",\""))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Costume{}
	for rows.Next() {
		var i Costume
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.SkinName,
			&i.Character,
			&i.CeCostumeCondition,
			&i.ConditionValue,
			&i.IllustReference,
			&i.PortraitReference,
			&i.VoiceAppear,
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
