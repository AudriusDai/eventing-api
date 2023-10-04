package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Load environment variables from `.env` and `.env.local` files.
func load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("WARN: was unable to load `.env` file - %q", err.Error())
	}

	if err := godotenv.Overload(".env.local"); err != nil {
		log.Printf("INFO: was unable to load `.env.local` file - %q", err.Error())
	}
}
