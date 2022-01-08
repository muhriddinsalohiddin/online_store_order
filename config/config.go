package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment       string
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	RPCPort           string
	ReviewServiceHost string
	ReviewServicePort int
}

func Load() (cfg Config) {
	cfg.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "book_order"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "muhriddin"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))

	cfg.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	cfg.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9006"))

	return
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
