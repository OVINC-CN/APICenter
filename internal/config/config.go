package config

import (
	"time"

	"github.com/ovinc-cn/apicenter/v2/pkg/cfg"
)

var Config configModel

func init() {
	Config = configModel{
		API: apiConfig{
			Addr:           cfg.GetConfigVal("APP_API_ADDR", "0.0.0.0:8000"),
			CookieDomain:   cfg.GetConfigVal("APP_API_COOKIE_DOMAIN", ""),
			CookieSecure:   cfg.GetConfigBool("APP_API_COOKIE_SECURE", true),
			CookieHTTPOnly: cfg.GetConfigBool("APP_API_COOKIE_HTTP_ONLY", true),
		},
		Cron: cronConfig{
			WorkerConcurrency:     cfg.GetConfigInt("APP_CRON_WORKER_CONCURRENCY", 1),
			WorkerQueues:          cfg.GetConfigStruct("APP_CRON_WORKER_QUEUES", `{"default": 1}`, make(map[string]int)),
			WorkerShutDownTimeout: time.Duration(cfg.GetConfigInt64("APP_CRON_WORKER_SHUTDOWN_TIMEOUT", 60)) * time.Second,
		},
		AppAccount: appAccountConfig{
			EmailDomain:         cfg.GetConfigVal("APP_ACCOUNT_EMAIL_DOMAIN", "example.com"),
			AuthTokenTimeout:    cfg.GetConfigInt("APP_ACCOUNT_AUTH_TOKEN_TIMEOUT", 7*24*60*60),
			AuthTokenCookieName: cfg.GetConfigVal("APP_ACCOUNT_AUTH_TOKEN_COOKIE_NAME", "auth-token"),
		},
	}
}
