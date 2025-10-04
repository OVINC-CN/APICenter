package config

import (
	"time"

	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
)

var Config configModel

func init() {
	Config = configModel{
		API: apiConfig{
			Addr: configUtils.GetConfigVal("APP_API_ADDR", "0.0.0.0:8000"),
		},
		Cron: cronConfig{
			WorkerConcurrency:     configUtils.GetConfigInt("APP_CRON_WORKER_CONCURRENCY", 1),
			WorkerQueues:          configUtils.GetConfigStruct("APP_CRON_WORKER_QUEUES", "{}", make(map[string]int)),
			WorkerShutDownTimeout: time.Duration(configUtils.GetConfigInt64("APP_CRON_WORKER_SHUTDOWN_TIMEOUT", 60)) * time.Second,
		},
	}
}
