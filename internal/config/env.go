package config

import (
	"fmt"
	"os"
)

func ValidateEnvVars(requiredVars ...string) error {
	for _, key := range requiredVars {
		if _, exists := os.LookupEnv(key); !exists {
			return fmt.Errorf("missing required environment variable: %s", key)
		}
	}
	return nil
}
