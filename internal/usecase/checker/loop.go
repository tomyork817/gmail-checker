package checker

import (
	"go.uber.org/zap"
	"route256-gmail-checker/internal/domain"
	"strings"
	"time"
)

func (c *EmailChecker) Start() error {
	for range time.Tick(c.cfg.Interval) {
		messageList, err := c.email.GetLast10MessageIDs(c.cfg.Search)
		if err != nil {
			return err
		}

		for _, id := range messageList {
			message, err := c.email.GetMessageByID(id)
			if err != nil {
				return err
			}

			if message.Status == domain.StatusRead {
				continue
			}

			if c.CheckMatches(message) {
				c.logger.Info("new message!",
					zap.String("From", message.From),
					zap.String("Subject", message.Subject),
					zap.String("Snippet", message.Snippet))
				if err = c.SendMessage(message); err != nil {
					return err
				}
			}

			if err = c.ReadMessage(&message); err != nil {
				return err
			}
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
