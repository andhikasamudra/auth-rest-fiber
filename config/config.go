package config

import "os"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) TokenSecret() string {
	return os.Getenv("TOKEN_SECRET")
}
