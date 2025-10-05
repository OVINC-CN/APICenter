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

type appAccountConfig struct {
	EmailDomain             string
	AuthTokenTimeout        int
	AuthTokenCookieName     string
	AuthTokenCookieDomain   string
	AuthTokenCookieSecure   bool
	AuthTokenCookieHTTPOnly bool
}

type configModel struct {
	API        apiConfig
	Cron       cronConfig
	AppAccount appAccountConfig
}
