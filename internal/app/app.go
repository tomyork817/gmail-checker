package app

import (
	"context"
	"go.uber.org/zap"
	"log"
	gmail2 "route256-gmail-checker/internal/infrastructure/gmail"
	"route256-gmail-checker/internal/usecase/checker"
	"route256-gmail-checker/pkg/config"
	"route256-gmail-checker/pkg/gmail"
	logger2 "route256-gmail-checker/pkg/logger"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	logger, err := logger2.NewZapLogger()
	if err != nil {
		log.Fatal(err)
	}

	gmailService, err := gmail.NewGmailService(ctx, cfg.GoogleAPI)
	if err != nil {
		logger.Error("unable to create gmail service", zap.Error(err))
	}

	gmailClient := gmail2.NewClient(gmailService)
	telegramClient := // ...

	emailChecker := checker.NewEmailChecker(gmailClient, telegramClient, cfg.Checker, logger)

	go emailChecker.Start()
}
