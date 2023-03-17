package db_game

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const createCharacterBroadcasts = `-- name: CreateCharacterBroadcasts :execresult
INSERT INTO character_broadcast (account_user_id, character_enum_id, timeline_enum_id, ` + "`" + `type` + "`" + `, on_air, complete) VALUES %s
`

func (q *Queries) CreateCharacterBroadcasts(ctx context.Context, characterBroadcasts []CreateCharacterBroadcastParams) (sql.Result, error) {
	params := []string{}
	for _, characterBroadcast := range characterBroadcasts {
		params = append(params, fmt.Sprintf("(%d, '%s', '%s', %d, %d, %d)", characterBroadcast.AccountUserID, characterBroadcast.CharacterEnumID, characterBroadcast.TimelineEnumID, characterBroadcast.Type, characterBroadcast.OnAir, characterBroadcast.Complete))
	}
	query := fmt.Sprintf(createCharacterBroadcasts, strings.Join(params, ","))

	return q.db.ExecContext(ctx, query)
}
