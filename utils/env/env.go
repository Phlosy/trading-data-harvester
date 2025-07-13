package utils

import (
	"os"
	"strings"
)

func GetEnvOrDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvListOrDefault(key string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		// Split the value by comma
		return strings.Split(value, ",")
	}
	return fallback
}
