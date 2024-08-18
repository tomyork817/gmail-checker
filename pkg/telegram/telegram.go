package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Config struct {
	APIToken string `mapstructure:"api_token"`
}

func NewBotAPI(cfg Config) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(cfg.APIToken)
}
