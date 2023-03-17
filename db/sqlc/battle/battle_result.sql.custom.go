package db_battle

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const updateBattleResultsConfirmedAt = `-- name: UpdateBattleResultsConfirmedAt :execresult
UPDATE battle_result SET confirmed_at = ? WHERE id IN (%s)
`

func (q *Queries) UpdateBattleResultsConfirmedAt(ctx context.Context, confirmedAt sql.NullTime, battleResultsId []int32) (sql.Result, error) {
	ids := []string{}
	for _, id := range battleResultsId {
		ids = append(ids, fmt.Sprintf("%d", id))
	}

	query := fmt.Sprintf(updateBattleResultsConfirmedAt, strings.Join(ids, ","))

	return q.db.ExecContext(ctx, query, confirmedAt)
}
