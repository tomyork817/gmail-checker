package main

import (
	"gmail-checker/internal/app"
	"gmail-checker/pkg/config"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error while getting config: %s", err)
	}

	app.Run(cfg)
}
