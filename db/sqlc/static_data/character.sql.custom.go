package db_static_data

import (
	"context"
	"fmt"
	"strings"
)

const getCharactersByEnumIds = `-- name: getCharactersByEnumIds :many
SELECT id, enum_id, character_root, ce_character_species, crew, ce_character_class, ce_character_property, mbti, hobby, partner, ce_character_grade, base_stat, level_stat, resource_list, skill_set, item_preference, signature_weapon, costume_bundle, second_name, library, profile, unit_size, basic, active, created_at FROM ` + "`" + `character` + "`" + ` WHERE enum_id IN ("%s")
`

func (q *Queries) GetCharactersByEnumIds(ctx context.Context, enumIDs []string) ([]Character, error) {
	query := fmt.Sprintf(getCharactersByEnumIds, strings.Join(enumIDs, "\",\""))
	rows, err := q.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Character{}
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.EnumID,
			&i.CharacterRoot,
			&i.CeCharacterSpecies,
			&i.Crew,
			&i.CeCharacterClass,
			&i.CeCharacterProperty,
			&i.Mbti,
			&i.Hobby,
			&i.Partner,
			&i.CeCharacterGrade,
			&i.BaseStat,
			&i.LevelStat,
			&i.ResourceList,
			&i.SkillSet,
			&i.ItemPreference,
			&i.SignatureWeapon,
			&i.CostumeBundle,
			&i.SecondName,
			&i.Library,
			&i.Profile,
			&i.UnitSize,
			&i.Basic,
			&i.Active,
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
