package checker

import (
	"go.uber.org/zap"
	"time"
)

type EmailChecker struct {
	email     Email
	messenger Messenger
	cfg       Config
	logger    *zap.Logger
}

type Config struct {
	Interval         time.Duration `mapstructure:"interval" validate:"required" default:"1m"`
	MessagesCount    int           `mapstructure:"messages_count" validate:"required" default:"5"`
	Search           string        `mapstructure:"search" validate:"required" default:"-"`
	SubjectFragments []string      `mapstructure:"subject_fragments" validate:"required" default:"[]"`
	FromFragments    []string      `mapstructure:"from_fragments" validate:"required" default:"[]"`
	SnippetFragments []string      `mapstructure:"snippet_fragments" validate:"required" default:"[]"`
	ChatIDs          []int64       `mapstructure:"chat_ids" validate:"required"`
}

func NewEmailChecker(email Email, messenger Messenger, cfg Config, logger *zap.Logger) *EmailChecker {
	return &EmailChecker{email: email, messenger: messenger, cfg: cfg, logger: logger}
}
