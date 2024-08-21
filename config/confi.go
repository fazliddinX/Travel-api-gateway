package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HTTP_PORT            string
	ACCESS_TOKEN         string
	USER_SERVICE_PORT    string
	CONTENT_SERVICE_PORT string
	USER_HOST            string
	CONTENT_HOST         string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := Config{}

	config.HTTP_PORT = cast.ToString(coalesce("HTTP_PORT", ":8080"))
	config.CONTENT_SERVICE_PORT = cast.ToString(coalesce("CONTENT_SERVICE_PORT", "50055"))
	config.ACCESS_TOKEN = cast.ToString(coalesce("ACCESS_TOKEN", "abcde"))
	config.USER_SERVICE_PORT = cast.ToString(coalesce("USER_SERVICE_PORT", "50056"))
	config.USER_HOST = cast.ToString(coalesce("USER_HOST", "localhost"))
	config.CONTENT_HOST = cast.ToString(coalesce("CONTENT_HOST", "localhost"))
	return config
}

func coalesce(env string, defaultValue interface{}) interface{} {
	value, exists := os.LookupEnv(env)
	if !exists {
		return defaultValue
	}
	return value
}
