package config

import (
	"github.com/audriusdai/eventing-api/util"
)

var (
	APP_ADDRESS                       string
	GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS int
)

func populateConfig() {
	APP_ADDRESS = util.GetEnvOrDefault("APP_ADDRESS", ":8080")
	GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS = util.GetEnvAsIntOrDefault("GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS", 5)
}
