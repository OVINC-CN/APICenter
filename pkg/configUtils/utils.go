package configUtils

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// load from env file
	if err := godotenv.Load(); err != nil {
		log.Printf("[Config] error reading config: %s", err)
	}
}

func GetConfigVal(key string, defaultVal string) string {
	// load from env
	if val := os.Getenv(key); val != "" {
		return val
	}
	// default val
	return defaultVal
}

func GetConfigBool(key string, defaultVal bool) bool {
	val := GetConfigVal(key, "")
	if val == "" {
		return defaultVal
	}
	return strings.ToLower(val) == "true"
}

func GetConfigInt(key string, defaultVal int) int {
	val := GetConfigVal(key, "")
	if val == "" {
		return defaultVal
	}
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("[Config] parse %s of %s to int failed; %s", key, val, err)
	}
	return intVal
}

func GetConfigInt64(key string, defaultVal int64) int64 {
	val := GetConfigVal(key, "")
	if val == "" {
		return defaultVal
	}
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Fatalf("[Config] parse %s of %s to int64 failed; %s", key, val, err)
	}
	return intVal
}

func GetConfigUint(key string, defaultVal uint) uint {
	val := GetConfigVal(key, "")
	if val == "" {
		return defaultVal
	}
	uintVal, err := strconv.ParseUint(val, 10, strconv.IntSize)
	if err != nil {
		log.Fatalf("[Config] parse %s of %s to uint failed; %s", key, val, err)
	}
	return uint(uintVal)
}

func GetConfigUint64(key string, defaultVal uint64) uint64 {
	val := GetConfigVal(key, "")
	if val == "" {
		return defaultVal
	}
	uintVal, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		log.Fatalf("[Config] parse %s of %s to uint64 failed; %s", key, val, err)
	}
	return uintVal
}

func GetConfigFloat64(key string, defaultVal float64) float64 {
	val := GetConfigVal(key, "")
	if val == "" {
		return defaultVal
	}
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Fatalf("[Config] parse %s of %s to float64 failed; %s", key, val, err)
	}
	return floatVal
}

func GetConfigStruct[T interface{}](key string, defaultVal string, out T) T {
	val := GetConfigVal(key, defaultVal)
	if err := json.Unmarshal([]byte(val), &out); err != nil {
		log.Fatalf("[Config] parse %s of `%s` to struct failed; %s", key, val, err)
	}
	return out
}

func MustGetConfigVal(key string) string {
	if val := GetConfigVal(key, ""); val != "" {
		return val
	}
	log.Fatalf("[Config] %s is empty", key)
	return ""
}
