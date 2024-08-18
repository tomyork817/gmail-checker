package checker

import "route256-gmail-checker/internal/domain"

type Email interface {
	GetLast10MessageIDs(searchQuery string) ([domain.MessagesListLen]string, error)
	GetMessageByID(id string) (domain.Message, error)
	DeleteLabelByID(messageID, labelID string) error
}

type Messenger interface {
	SendMessage(chatID int64, message domain.Message) error
}
