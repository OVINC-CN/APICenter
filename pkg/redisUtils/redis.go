package redisUtils

import (
	"context"
	"log"

	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	// init client
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     configUtils.RedisAddr(),
			Password: configUtils.RedisPassword(),
			DB:       configUtils.RedisDB(),
		},
	)
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("[RedisClient] failed to ping; %s", err)
	}

	// log
	log.Printf("[RedisClient] connect success\n")
}

func Client() *redis.Client {
	return rdb
}
