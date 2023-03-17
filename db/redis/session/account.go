package redis_session

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (q *Queries) GetAccountSession(ctx context.Context, accountId int32) (Session, error) {
	var session Session
	sessionKey := fmt.Sprintf("ss_%d", accountId)

	result := q.db.HGetAll(ctx, sessionKey)
	err := result.Err()
	if err != nil {
		return session, err
	}
	if err = result.Scan(&session); err != nil {
		return session, err
	}

	return session, nil
}

func (q *Queries) CreateAccountSession(ctx context.Context, session *Session) error {
	sessionKey := fmt.Sprintf("ss_%d", session.AccountId)
	sessionId := uuid.NewString()
	session.SessionId = sessionId

	err := q.db.HSet(ctx, sessionKey, map[string]interface{}{
		"session_id":      session.SessionId,
		"world_id":        session.WorldId,
		"account_id":      session.AccountId,
		"account_user_id": session.AccountUserId,
		"game_db":         session.GameDb,
		"nickname":        session.Nickname,
		"signed_in_at":    session.SignedInAt,
	}).Err()
	if err != nil {
		return err
	}

	q.UpdateExpire(ctx, sessionKey, 3600*time.Second)

	return nil
}

func (q *Queries) UpdateAccountSession(ctx context.Context, session Session) error {
	sessionKey := fmt.Sprintf("ss_%d", session.AccountId)

	err := q.db.HSet(ctx, sessionKey, map[string]interface{}{
		"session_id":      session.SessionId,
		"world_id":        session.WorldId,
		"account_id":      session.AccountId,
		"account_user_id": session.AccountUserId,
		"game_db":         session.GameDb,
		"nickname":        session.Nickname,
		"signed_in_at":    session.SignedInAt,
	}).Err()
	if err != nil {
		return err
	}

	q.UpdateExpire(ctx, sessionKey, 3600*time.Second)

	return nil
}
