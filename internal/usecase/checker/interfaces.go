package checker

import "gmail-checker/internal/domain"

type Email interface {
	GetLastNMessageIDs(searchQuery string, count int) ([]string, error)
	GetMessageByID(id string) (domain.Message, error)
	DeleteLabelByID(messageID, labelID string) error
}

type Messenger interface {
	SendMessage(chatID int64, message domain.Message) error
}
