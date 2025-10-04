package redisUtils

import (
	"context"
	"fmt"
	"time"

	"github.com/ovinc-cn/apicenter/v2/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	locked bool

	ctx  context.Context
	conn redis.Cmdable

	key        string
	expiration time.Duration
}

func (l *RedisLock) Lock() error {
	result, err := SetNX(l.ctx, l.conn, l.key, time.Now().Format(time.RFC3339Nano), l.expiration)
	if err != nil {
		return err
	}
	if !result {
		return &ConcurrencyLocked{Key: l.key}
	}
	l.locked = true
	return nil
}

func (l *RedisLock) Release() {
	if l.locked {
		_, err := Del(l.ctx, l.conn, l.key)
		if err != nil {
			logger.Logger(l.ctx).Error(fmt.Sprintf("[RedisLock] release lock %s failed; %s", l.key, err.Error()))
		}
	}
}

func NewLock(ctx context.Context, conn redis.Cmdable, key string, expiration time.Duration) *RedisLock {
	return &RedisLock{ctx: ctx, conn: conn, key: key, expiration: expiration}
}
