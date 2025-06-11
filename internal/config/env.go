package config

import (
	"errors"
	"os"
)

type EnvConfig struct {
	SecretName string
	UrlPrefix  string
	DBName     string
}

// ValidateEnvVars ensures that required environment variables are set
func ValidateEnvVars(vars ...string) error {
	for _, v := range vars {
		if _, exists := os.LookupEnv(v); !exists {
			return errors.New("missing required env var: " + v)
		}
	}
	return nil
}

// LoadConfig loads all configuration values from environment variables
func LoadConfig() (*EnvConfig, error) {
	// You can extend this with fallback defaults or stricter checks
	return &EnvConfig{
		SecretName: os.Getenv("SecretName"),
		UrlPrefix:  os.Getenv("UrlPrefix"),
		DBName:     "gambit", // Can be replaced with os.Getenv("DB_NAME") if needed
	}, nil
}
