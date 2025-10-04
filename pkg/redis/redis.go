package redis

import (
	"context"
	"log"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	// init client
	rdb = redis.NewClient(
		&redis.Options{
			Addr:     cfg.RedisAddr(),
			Password: cfg.RedisPassword(),
			DB:       cfg.RedisDB(),
		},
	)
	if err := Ping(context.Background(), rdb); err != nil {
		log.Fatalf("[RedisClient] failed to ping; %s", err)
	}

	// log
	log.Printf("[RedisClient] connect success\n")
}

func Client() *redis.Client {
	return rdb
}
