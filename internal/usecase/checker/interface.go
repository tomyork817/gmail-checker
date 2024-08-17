package checker

import "route256-gmail-checker/internal/domain"

type Gmail interface {
	ListMessages() []domain.Message
}
