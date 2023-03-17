package db_game

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/yeom-c/game-server-grpc-go/helper"
)

const getCostumeListByEnumId = `-- name: GetCostumeListByEnumId :many
SELECT id, account_user_id, enum_id, character_enum_id, state, created_at FROM costume WHERE account_user_id = %d AND enum_id IN ("%s")
`

func (q *Queries) GetCostumeListByEnumId(ctx context.Context, accountUserId int32, enumIdList []string) ([]Costume, error) {
	query := fmt.Sprintf(getCostumeListByEnumId, accountUserId, strings.Join(helper.RemoveEmpty(enumIdList), "\",\""))
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
			&i.AccountUserID,
			&i.EnumID,
			&i.CharacterEnumID,
			&i.State,
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

const createCostumes = `-- name: CreateCostumes :execresult
INSERT IGNORE INTO costume (account_user_id, enum_id, character_enum_id, ` + "`" + `state` + "`" + `) VALUES %s
`

func (q *Queries) CreateCostumes(ctx context.Context, costumeParams []CreateCostumeParams) (sql.Result, error) {
	params := []string{}
	for _, costumeParam := range costumeParams {
		params = append(params, fmt.Sprintf("(%d, '%s', '%s', %d)", costumeParam.AccountUserID, costumeParam.EnumID, costumeParam.CharacterEnumID, costumeParam.State))
	}
	query := fmt.Sprintf(createCostumes, strings.Join(params, ","))

	return q.db.ExecContext(ctx, query)
}
