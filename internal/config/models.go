package config

import "time"

type apiConfig struct {
	Addr           string
	CookieDomain   string
	CookieSecure   bool
	CookieHTTPOnly bool
}

type cronConfig struct {
	WorkerConcurrency     int
	WorkerQueues          map[string]int
	WorkerShutDownTimeout time.Duration
}

type appAccountConfig struct {
	EmailDomain         string
	AuthTokenTimeout    int
	AuthTokenCookieName string
}

type configModel struct {
	API        apiConfig
	Cron       cronConfig
	AppAccount appAccountConfig
}
