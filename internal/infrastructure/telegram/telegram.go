package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Client struct {
	botAPI *tgbotapi.BotAPI
}

func NewClient(botAPI *tgbotapi.BotAPI) *Client {
	return &Client{botAPI: botAPI}
}
