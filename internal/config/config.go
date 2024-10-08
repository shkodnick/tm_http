package config

import (
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env        string     `json:"env" env-default:"local"`
	Database   Database   `json:"database"`
	AppName    string     `json:"app_name"`
	HTTPServer HTTPServer `json:"http_server"`
}

type HTTPServer struct {
	Port int `json:"port" env:"HTTP_SERVER_PORT"`
}

type Database struct {
	Host     string `json:"host" env:"POSTGRES_HOST"`
	Port     int    `json:"port" env:"POSTGRES_PORT"`
	Database string `json:"database" env:"POSTGRES_DBNAME"`
	User     string `json:"user" env:"POSTGRES_USER"`
	Pass     string `json:"pass" env:"POSTGRES_PASSWORD"`
}

func CreateConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if strings.TrimSpace(configPath) == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Errorf("config file does not exist - `%s`", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
