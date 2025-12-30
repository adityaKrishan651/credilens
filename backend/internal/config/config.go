package config

import (
	"credilens-backend/internal/helpers"
	"os"
)

type Config struct {
	ServerPort string
	AppEnv     string

	AzureOpenAIEndpoint string
	AzureOpenAIKey      string
	AzureOpenAIDeploy   string
}

func Load() Config {
	return Config{
		ServerPort: helpers.Ternary(os.Getenv("SERVER_PORT") != "", os.Getenv("SERVER_PORT"), "8080"),
		AppEnv:     helpers.Ternary(os.Getenv("APP_ENV") != "", os.Getenv("APP_ENV"), "development"),

		AzureOpenAIEndpoint: helpers.Ternary(os.Getenv("AZURE_OPENAI_ENDPOINT") != "", os.Getenv("AZURE_OPENAI_ENDPOINT"), ""),
		AzureOpenAIKey:      helpers.Ternary(os.Getenv("AZURE_OPENAI_KEY") != "", os.Getenv("AZURE_OPENAI_KEY"), ""),
		AzureOpenAIDeploy:   helpers.Ternary(os.Getenv("AZURE_OPENAI_DEPLOYMENT") != "", os.Getenv("AZURE_OPENAI_DEPLOYMENT"), ""),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
