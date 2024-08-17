package domain

type Status string

const (
	StatusRead   Status = "READ"
	StatusUnread Status = "UNREAD"
)

type Message struct {
	ID      string
	From    string
	Subject string
	Snippet string
	Status  Status
}

const (
	MessagesListLen = 10
)
