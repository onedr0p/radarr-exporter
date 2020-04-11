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
		Hostname:       GetEnvStr("RADARR_HOSTNAME", "http://127.0.0.1:7878"),
		ApiKey:         GetEnvStr("RADARR_APIKEY", ""),
		BasicAuth:      GetEnvBool("BASIC_AUTH", false),
		BasicAuthCreds: GetEnvStr("BASIC_AUTH_CREDS", ""),
		Port:           GetEnvInt("PORT", 9707),
		LogLevel:       strings.ToUpper(GetEnvStr("LOG_LEVEL", "INFO")),
	}
}

// GetEnvStr -
func GetEnvStr(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// GetEnvInt -
func GetEnvInt(name string, defaultVal int) int {
	valueStr := GetEnvStr(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// GetEnvBool -
func GetEnvBool(name string, defaultVal bool) bool {
	valStr := GetEnvStr(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}
