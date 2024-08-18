package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gmail-checker/internal/domain"
)

const (
	messageFormat = `На ваш телефон пришло новое сообщение... Посмотри вдруг там что-то важное...
Тема: %s
Отправитель: %s
Фрагмент: %s`
)

func (c *Client) SendMessage(chatID int64, message domain.Message) error {
	messageText := fmt.Sprintf(messageFormat, message.Subject, message.From, message.Snippet)
	tgMessage := tgbotapi.NewMessage(chatID, messageText)
	_, err := c.botAPI.Send(tgMessage)
	return err
}
