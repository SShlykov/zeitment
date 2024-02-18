package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	CorsEnabled     bool          `yaml:"cors_enabled" env:"CORS_ENABLED" env-default:"false" env-required:"true"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT" env-default:"10s"`
	HTTPServer      `yaml:"http_server"`
	Logger          `yaml:"logger"`
	Swagger         `yaml:"swagger"`
}

type Swagger struct {
	SwaggerEnabled bool `yaml:"enabled" env:"ENABLED" env-default:"false" env-required:"true"`
}

type HTTPServer struct {
	Address      string        `yaml:"address" env:"ADDRESS" env-default:":8080" env-required:"true"`
	Timeout      time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"30s" env-required:"true"`
	IddleTimeout time.Duration `yaml:"iddle_timeout" env:"IDDLE_TIMEOUT" env-default:"30s" env-required:"true"`
}

type Logger struct {
	Level string `yaml:"level" env:"LEVEL" env-default:"debug" env-required:"true"`
}

func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("config file not found: %s", configPath))
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.New(fmt.Sprintf("config file read error: %v", err))
	}

	return &cfg, nil
}
