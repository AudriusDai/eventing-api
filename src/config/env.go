package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables from `.env` file.
func load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("was unable to load `.env` file - %q", err.Error())
	}
}
