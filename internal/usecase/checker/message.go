package checker

import "route256-gmail-checker/internal/domain"

const (
	labelIDUnread = "UNREAD"
)

func (c *EmailChecker) ReadMessage(message *domain.Message) error {
	if message.Status == domain.StatusRead {
		return nil
	}
	if err := c.email.DeleteLabelByID(message.ID, labelIDUnread); err != nil {
		return err
	}
	message.Status = domain.StatusRead
	return nil
}

func (c *EmailChecker) SendMessage(message domain.Message) error {
	for _, chatID := range c.cfg.ChatIDs {
		if err := c.messenger.SendMessage(chatID, message); err != nil {
			return err
		}
	}
	return nil
}
