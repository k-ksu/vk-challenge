package config

import (
	"fmt"
	"io"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		Cron `yaml:"cron"`
		HTTP `yaml:"http"`
		PG   `yaml:"postgres" envPrefix:"PG_"`
	}

	// App -.
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	Cron struct {
		Name              string `yaml:"name"`
		Version           string `yaml:"version"`
		Interval          string `yaml:"interval"`
		TechApiClientHost string `yaml:"tech_api_client_host"`
	}

	// HTTP -.
	HTTP struct {
		Port string `yaml:"port"`
	}

	// PG -.
	PG struct {
		PoolMax int    `yaml:"pool_max"`
		URL     string `env:"URL"`
	}
)

// New returns app config.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cfg.readConfigs(); err != nil {
		return nil, fmt.Errorf("cfg.readConfigs: %w", err)
	}

	if err := cfg.readSecrets(); err != nil {
		return nil, fmt.Errorf("cfg.readSecrets: %w", err)
	}

	return cfg, nil
}

func (cfg *Config) readConfigs() error {
	file, err := os.Open("./config/config.yaml")
	if err != nil {
		return fmt.Errorf("os.Open: %w", err)
	}

	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("io.ReadAll: %w", err)
	}

	if err = yaml.Unmarshal(b, cfg); err != nil {
		return fmt.Errorf("yaml.Unmarshal: %w", err)
	}

	return nil
}

func (cfg *Config) readSecrets() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("godotenv.Load: %w", err)
	}

	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("env.Parse: %w", err)
	}

	return nil
}
