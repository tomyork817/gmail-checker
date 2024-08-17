package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"route256-gmail-checker/pkg/gmail"
)

type Config struct {
	GoogleAPI gmail.Config `mapstructure:"google_api" validate:"required"`
}

const (
	configName = "config"
	configType = "toml"
	configPath = "config"
)

func GetConfig() (*Config, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
