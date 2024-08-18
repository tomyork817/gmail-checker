package app

import (
	"context"
	gmail2 "gmail-checker/internal/infrastructure/gmail"
	telegram2 "gmail-checker/internal/infrastructure/telegram"
	"gmail-checker/internal/usecase/checker"
	"gmail-checker/pkg/config"
	"gmail-checker/pkg/gmail"
	logger2 "gmail-checker/pkg/logger"
	"gmail-checker/pkg/telegram"
	"go.uber.org/zap"
	"log"
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
	wg.Add(1)
	go func() {
		err = emailChecker.Start()
		logger.Error("stopped fetch cycle", zap.Error(err))
		wg.Done()
	}()

	wg.Wait()
	logger.Info("stopping application")
}
