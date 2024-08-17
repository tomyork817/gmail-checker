package config

import "route256-gmail-checker/pkg/gmail"

type Config struct {
	GoogleAPI gmail.Config `mapstructure:"google_api"`
}
