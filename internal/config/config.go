package config

import (
	"time"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
)

var Config configModel

func init() {
	Config = configModel{
		API: apiConfig{
			Addr: cfg.GetConfigVal("APP_API_ADDR", "0.0.0.0:8000"),
		},
		Cron: cronConfig{
			WorkerConcurrency:     cfg.GetConfigInt("APP_CRON_WORKER_CONCURRENCY", 1),
			WorkerQueues:          cfg.GetConfigStruct("APP_CRON_WORKER_QUEUES", "{}", make(map[string]int)),
			WorkerShutDownTimeout: time.Duration(cfg.GetConfigInt64("APP_CRON_WORKER_SHUTDOWN_TIMEOUT", 60)) * time.Second,
		},
	}
}
