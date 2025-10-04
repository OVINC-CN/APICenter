package config

import "time"

type apiConfig struct {
	Addr string
}

type cronConfig struct {
	WorkerConcurrency     int
	WorkerQueues          map[string]int
	WorkerShutDownTimeout time.Duration
}

type configModel struct {
	API  apiConfig
	Cron cronConfig
}
