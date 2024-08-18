package app

import (
	"context"
	"go.uber.org/zap"
	"log"
	gmail2 "route256-gmail-checker/internal/infrastructure/gmail"
	telegram2 "route256-gmail-checker/internal/infrastructure/telegram"
	"route256-gmail-checker/internal/usecase/checker"
	"route256-gmail-checker/pkg/config"
	"route256-gmail-checker/pkg/gmail"
	logger2 "route256-gmail-checker/pkg/logger"
	"route256-gmail-checker/pkg/telegram"
	"sync"
)

func Run(cfg *config.Config) {
	ctx := context.Background()

	logger, err := logger2.NewZapLogger()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("starting application")

	gmailService, err := gmail.NewGmailService(ctx, cfg.GoogleAPI)
	if err != nil {
		logger.Error("unable to create gmail service", zap.Error(err))
	}

	botAPI, err := telegram.NewBotAPI(cfg.Telegram)
	if err != nil {
		logger.Error("unable to create bot api", zap.Error(err))
	}

	gmailClient := gmail2.NewClient(gmailService)
	telegramClient := telegram2.NewClient(botAPI)

	emailChecker := checker.NewEmailChecker(gmailClient, telegramClient, cfg.Checker, logger)

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		err = emailChecker.Start()
		logger.Error("stopped fetch cycle", zap.Error(err))
	}()

	wg.Wait()
	logger.Info("stopping application")
}
