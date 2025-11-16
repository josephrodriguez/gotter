package config

import (
	"log"
	"os"
	"strconv"
)

// GetEnv returns the environment variable value for key or fallback if not set.
func GetEnvStr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// GetEnvInt returns the environment variable parsed as int, or fallback on error/not set.
func GetEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
		log.Printf("[Gotter] invalid value for %s: %v; using default %d", key, err, fallback)
	}
	return fallback
}
