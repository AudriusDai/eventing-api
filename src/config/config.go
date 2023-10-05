package config

import (
	"github.com/audriusdai/eventing-api/util"
)

var (
	// base
	APP_ADDRESS                       string
	GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS int

	// db
	DB_HOSTNAME string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_SCHEMA   string
)

func populateConfig() {
	// base
	APP_ADDRESS = util.GetEnvOrDefault("APP_ADDRESS", ":8080")
	GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS = util.GetEnvAsIntOrDefault("GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS", 5)

	// db
	DB_HOSTNAME = util.GetEnvOrDefault("DB_HOSTNAME", "127.0.0.1")
	DB_NAME = util.GetEnvOrDefault("DB_NAME", "postgres")
	DB_SCHEMA = util.GetEnvOrDefault("DB_SCHEMA", "public")
	DB_PORT = util.GetEnvOrDefault("DB_PORT", "5432")
	DB_USERNAME = util.GetEnvOrDefault("DB_USERNAME", "postgres")
	DB_PASSWORD = util.GetEnvOrDefault("DB_PASSWORD", "")
}
