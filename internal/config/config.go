package config

import (
	"log"
	"os"
)

func MustGetenv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s not set", key)
	}
	return val
}

func Getenv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
