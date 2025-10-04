package configUtils

import (
	"log"
	"time"
)

// Base
var (
	AppName = func() string {
		return MustGetConfigVal("APP_NAME")
	}
	AppDebug = func() bool {
		return GetConfigBool("APP_DEBUG", false)
	}
	AppTimezone = func() *time.Location {
		tz := GetConfigVal("APP_TIMEZONE", "Asia/Shanghai")
		loc, err := time.LoadLocation(tz)
		if err != nil {
			log.Fatalf("[Config] load timezone %s failed; %s", tz, err)
		}
		return loc
	}
)

// Log
var (
	AppLogLevel = func() string {
		return GetConfigVal("APP_LOG_LEVEL", "INFO")
	}
	AppLogEncoding = func() string {
		return GetConfigVal("APP_LOG_ENCODING", "json")
	}
	AppLogOutput = func() string {
		return GetConfigVal("APP_LOG_OUTPUT", "stdout")
	}
)

// db
var (
	RedisAddr = func() string {
		return GetConfigVal("APP_REDIS_ADDR", "127.0.0.1:6379")
	}
	RedisPassword = func() string {
		return GetConfigVal("APP_REDIS_PASSWORD", "")
	}
	RedisDB = func() int {
		return GetConfigInt("APP_REDIS_DB", 0)
	}
)

// Trace
var (
	TraceEndpoint = func() string {
		return GetConfigVal("APP_TRACE_ENDPOINT", "127.0.0.1:4317")
	}
)
