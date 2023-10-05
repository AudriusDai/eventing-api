package main

import (
	"log"

	"github.com/audriusdai/eventing-api/config"
	"github.com/audriusdai/eventing-api/db"
	"github.com/audriusdai/eventing-api/web"
)

func main() {
	config.SetupConfig()
	if err := db.SetupDb(); err != nil {
		log.Fatal(err)
	}

	if err := web.SetupWeb(); err != nil {
		log.Fatal(err)
	}
}
