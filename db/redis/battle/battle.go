package redis_battle

import (
	"context"
	"time"
)

func (q *Queries) EnterChannel(ctx context.Context, channelId string, accountUsersId ...interface{}) error {
	err := q.db.SAdd(ctx, channelId, accountUsersId...).Err()
	q.UpdateExpire(ctx, channelId, 60*time.Second)
	return err
}

func (q *Queries) LeaveChannel(ctx context.Context, channelId string, accountUserId ...interface{}) error {
	return q.db.SRem(ctx, channelId, accountUserId).Err()
}

func (q *Queries) GetChannelAccountUsers(ctx context.Context, channelId string) ([]string, error) {
	members := q.db.SMembers(ctx, channelId)

	return members.Result()
}
