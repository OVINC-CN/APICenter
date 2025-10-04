package crond

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
)

var AsyncClient *asynq.Client
var RedisClientOption = asynq.RedisClientOpt{
	Addr:     configUtils.RedisAddr(),
	Password: configUtils.RedisPassword(),
	DB:       configUtils.RedisDB(),
}

func init() {
	AsyncClient = asynq.NewClient(RedisClientOption)
	log.Printf("[AsyncClinet] init success\n")
}
