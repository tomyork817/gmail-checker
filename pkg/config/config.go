package config

import (
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"route256-gmail-checker/internal/usecase/checker"
	"route256-gmail-checker/pkg/gmail"
	"route256-gmail-checker/pkg/telegram"
)

type Config struct {
	GoogleAPI gmail.Config    `mapstructure:"google_api" validate:"required"`
	Checker   checker.Config  `mapstructure:"checker" validate:"required"`
	Telegram  telegram.Config `mapstructure:"telegram" validate:"required"`
}

const (
	configName = "config"
	configType = "toml"
	configPath = "config"
)

func GetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.BindEnv("telegram.api_token", "TG_TOKEN"); err != nil {
		return nil, err
	}

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := defaults.Set(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
