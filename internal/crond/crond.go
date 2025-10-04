package crond

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
)

var AsyncClient *asynq.Client
var RedisClientOption = asynq.RedisClientOpt{
	Addr:     cfg.RedisAddr(),
	Password: cfg.RedisPassword(),
	DB:       cfg.RedisDB(),
}

func init() {
	AsyncClient = asynq.NewClient(RedisClientOption)
	log.Printf("[AsyncClient] init success\n")
}
