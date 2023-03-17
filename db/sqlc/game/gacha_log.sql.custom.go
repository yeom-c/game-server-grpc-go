package db_game

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const createGachaLogs = `-- name: CreateGachaLogs :execresult
INSERT INTO gacha_log (account_user_id, enum_id, character_enum_id) VALUES %s
`

func (q *Queries) CreateGachaLogs(ctx context.Context, gachaLogs []CreateGachaLogParams) (sql.Result, error) {
	params := []string{}
	for _, gachaLog := range gachaLogs {
		params = append(params, fmt.Sprintf("(%d, '%s', '%s')", gachaLog.AccountUserID, gachaLog.EnumID, gachaLog.CharacterEnumID))
	}
	query := fmt.Sprintf(createGachaLogs, strings.Join(params, ","))

	return q.db.ExecContext(ctx, query)
}
