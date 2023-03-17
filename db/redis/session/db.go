package redis_session

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Queries struct {
	db *redis.Client
}

var sessionCtxKey = "mySession"

func GetMySession(ctx context.Context) Session {
	return ctx.Value(sessionCtxKey).(Session)
}

func SetMySession(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionCtxKey, session)
}

func New(db *redis.Client) *Queries {
	return &Queries{
		db: db,
	}
}

func (q *Queries) UpdateExpire(ctx context.Context, sessionKey string, expire time.Duration) {
	q.db.Expire(ctx, sessionKey, expire)
}
