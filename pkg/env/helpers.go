package env

import (
	"os"
)

func GetEnvValueWithFallback(envName string, fallback string) string {
	envValue := os.Getenv(envName)
	if envValue == "" {
		return fallback
	}
	return envValue
}
