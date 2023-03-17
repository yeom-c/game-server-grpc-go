package redis_battle

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Queries struct {
	db *redis.Client
}

func New(db *redis.Client) *Queries {
	return &Queries{
		db: db,
	}
}

func (q *Queries) UpdateExpire(ctx context.Context, key string, expire time.Duration) {
	q.db.Expire(ctx, key, expire)
}
