package main

import (
	"log"

	"github.com/audriusdai/eventing-api/config"
	"github.com/audriusdai/eventing-api/web"
)

func main() {
	config.SetupConfig()

	if err := web.SetupWeb(); err != nil {
		log.Fatal(err)
	}
}
