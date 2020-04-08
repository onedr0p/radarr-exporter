package config

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
)

// Config is the exporter CLI configuration.
type Config struct {
	Hostname     string        `config:"radarr_hostname"`
	ApiKey       string        `config:"radarr_apikey"`
	Port         string        `config:"port"`
	AuthType     string        `config:"auth_type"`
	Interval     time.Duration `config:"interval"`
	StartupDelay time.Duration `config:"startup_delay"`
}

func getDefaultConfig() *Config {
	return &Config{
		Hostname:     "http://127.0.0.1:7878",
		ApiKey:       "",
		Port:         "9811",
		Interval:     10 * time.Minute,
		StartupDelay: 0 * time.Second,
	}
}

// Load method loads the configuration by using both flag or environment variables.
func Load() *Config {
	loaders := []backend.Backend{
		env.NewBackend(),
		flags.NewBackend(),
	}

	loader := confita.NewLoader(loaders...)

	cfg := getDefaultConfig()
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		panic(err)
	}

	cfg.show()

	return cfg
}

func (c Config) show() {
	val := reflect.ValueOf(&c).Elem()
	log.Println("-------------------------------------")
	log.Println("-   Radarr exporter configuration   -")
	log.Println("-------------------------------------")
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		log.Println(fmt.Sprintf("%s : %v", typeField.Name, valueField.Interface()))
	}
	log.Println("------------------------------------")
}
