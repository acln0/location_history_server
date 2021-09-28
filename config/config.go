package config

import "C"
import (
	"log"
	"os"

	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

func getEnv(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

var Conf Config

type Config struct {
	Server struct {
		Port string `env:"HISTORY_SERVER_LISTEN_ADDR"`
	}
}

func InitConfig(envPath string) {
	err := config.New().
		AddFeeder(feeder.DotEnv{Path: getEnv("LOCATION_SERVER_ENV_FILE", envPath)}).
		AddFeeder(feeder.Env{}).
		AddStruct(&Conf).
		Feed()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}
