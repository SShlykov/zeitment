package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type HTTPServer struct {
	Address                  string        `yaml:"address" env:"ADDRESS" env-default:":8080" env-required:"true"`
	Timeout                  time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"30s" env-required:"true"`
	IddleTimeout             time.Duration `yaml:"iddle_timeout" env:"IDDLE_TIMEOUT" env-default:"30s" env-required:"true"`
	RequestLimit             int           `yaml:"request_limit" env:"REQUEST_LIMIT" env-default:"100" env-required:"true"`
	MinRequests              int           `yaml:"min_requests" env:"MIN_REQUESTS" env-default:"10" env-required:"true"`
	ErrorThresholdPercentage float64       `yaml:"error_threshold_percentage" env-default:"0.6" env-required:"true"`
	IntervalDuration         time.Duration `yaml:"interval_duration" env:"INTERVAL_DURATION" env-default:"10s" env-required:"true"`
	OpenStateTimeout         time.Duration `yaml:"open_state_timeout" env:"OPEN_STATE_TIMEOUT" env-default:"10s" env-required:"true"`
	SwaggerEnabled           bool          `yaml:"swagger_enabled" env:"ENABLED" env-default:"false" env-required:"true"`
	CorsEnabled              bool          `yaml:"cors_enabled" env:"CORS_ENABLED" env-default:"false" env-required:"true"`
}

func LoadServerConfig(configPath string) (*HTTPServer, error) {
	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	configPath = configPath + "/server.yml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New(fmt.Sprintf("config file not found: %s", configPath))
	}

	var cfg HTTPServer

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.New(fmt.Sprintf("config file read error: %v", err))
	}

	return &cfg, nil
}
