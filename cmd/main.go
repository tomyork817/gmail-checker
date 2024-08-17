package main

import (
	"fmt"
	"log"
	"route256-gmail-checker/internal/app"
	"route256-gmail-checker/pkg/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("error while getting config: %s", err)
	}

	fmt.Println(*cfg)

	app.Run(cfg)
}
