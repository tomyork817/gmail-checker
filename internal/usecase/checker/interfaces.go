package checker

import "route256-gmail-checker/internal/domain"

type Mail interface {
	GetLast10MessageIDs(searchQuery string) ([domain.MessagesListLen]string, error)
	GetMessageByID(id string) (domain.Message, error)
	DeleteLabelByID(messageID, labelID string) error
}
