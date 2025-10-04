package logger

import (
	"log"
	"strings"

	"github.com/ovinc-cn/apicenter/v2/pkg/configUtils"
	"go.uber.org/zap"
)

func LogLevel() zap.AtomicLevel {
	level := strings.ToUpper(configUtils.AppLogLevel())
	switch level {
	case "DEBUG":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "INFO":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "WARN", "WARNING":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "ERROR":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "DPANIC":
		return zap.NewAtomicLevelAt(zap.DPanicLevel)
	case "PANIC":
		return zap.NewAtomicLevelAt(zap.PanicLevel)
	case "FATAL":
		return zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		log.Printf("[Logger] unknown log level %s, use INFO level", level)
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}
