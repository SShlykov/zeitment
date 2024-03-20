package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" env-default:"10s"`
	Logger          `yaml:"logger"`
	Server          `yaml:"server"`
}

type Logger struct {
	Level string `yaml:"level" env:"LEVEL" env-default:"debug" env-required:"true"`
}

type Server struct {
	Port int `yaml:"port" env:"PORT" env-default:"8080"`
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	configPath = configPath + "/default.yml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("config file not found: %s", configPath))
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.New(fmt.Sprintf("config file read error: %v", err))
	}

	return &cfg, nil
}
