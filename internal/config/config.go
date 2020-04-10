package config

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Hostname       string
	ApiKey         string
	BasicAuth      bool
	BasicAuthCreds string
	Port           int
	LogLevel       string
}

func New() *Config {
	return &Config{
		Hostname:       getEnvStr("RADARR_HOSTNAME", "127.0.0.1"),
		ApiKey:         getEnvStr("RADARR_APIKEY", ""),
		BasicAuth:      getEnvBool("BASIC_AUTH", false),
		BasicAuthCreds: getEnvStr("BASIC_AUTH_CREDS", ""),
		Port:           getEnvInt("PORT", 9811),
		LogLevel:       strings.ToUpper(getEnvStr("LOG_LEVEL", "INFO")),
	}
}

func getEnvStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvInt(name string, defaultVal int) int {
	valueStr := getEnvStr(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvBool(name string, defaultVal bool) bool {
	valStr := getEnvStr(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
