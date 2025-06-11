package config

import (
	"os"
)

type EnvConfig struct {
	SecretName string
	UrlPrefix  string
	DBName     string
}

// LoadConfig loads all configuration values from environment variables
func LoadConfig() (*EnvConfig, error) {
	return &EnvConfig{
		SecretName: os.Getenv("SecretName"),
		UrlPrefix:  os.Getenv("UrlPrefix"),
		DBName:     "gambit", // Can be replaced with os.Getenv("DB_NAME") if needed
	}, nil
}
