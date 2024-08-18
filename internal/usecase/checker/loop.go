package checker

import (
	"go.uber.org/zap"
	"route256-gmail-checker/internal/domain"
	"strings"
	"time"
)

func (c *EmailChecker) Start() error {
	c.logger.Info("starting email checker loop")
	for range time.Tick(c.cfg.Interval) {
		c.logger.Info("new iteration of loop")

		messageList, err := c.email.GetLastNMessageIDs(c.cfg.Search, c.cfg.MessagesCount)
		if err != nil {
			return err
		}

		for _, id := range messageList {
			message, err := c.email.GetMessageByID(id)
			if err != nil {
				return err
			}

			c.logger.Info("reading message",
				zap.String("ID", message.ID),
				zap.String("Subject", message.Subject))

			if message.Status == domain.StatusRead {
				continue
			}

			c.logger.Info("found new unread message",
				zap.String("ID", message.ID),
				zap.String("Subject", message.Subject))

			if c.CheckMatches(message) {
				c.logger.Info("new message with fragments!",
					zap.String("ID", message.ID),
					zap.String("From", message.From),
					zap.String("Subject", message.Subject),
					zap.String("Snippet", message.Snippet))
				if err = c.SendMessage(message); err != nil {
					return err
				}
				c.logger.Info("sent new message to telegram",
					zap.String("ID", message.ID),
					zap.String("Subject", message.Subject))
			}

			if err = c.ReadMessage(&message); err != nil {
				return err
			}
			c.logger.Info("marked message as read",
				zap.String("ID", message.ID),
				zap.String("Subject", message.Subject))
		}
	}
	return nil
}

func (c *EmailChecker) CheckMatches(message domain.Message) bool {
	for _, fragment := range c.cfg.SubjectFragments {
		if strings.Contains(message.Subject, fragment) {
			return true
		}
	}

	for _, fragment := range c.cfg.FromFragments {
		if strings.Contains(message.From, fragment) {
			return true
		}
	}

	for _, fragment := range c.cfg.SnippetFragments {
		if strings.Contains(message.Snippet, fragment) {
			return true
		}
	}

	return false
}
