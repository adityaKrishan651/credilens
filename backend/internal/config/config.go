package config

import "os"

type AzureConfig struct {
	OpenAIKey   string
	SearchKey   string
	VisionKey  string
	Endpoint   string
}

func LoadAzureConfig() AzureConfig {
	return AzureConfig{
		OpenAIKey: os.Getenv("AZURE_OPENAI_KEY"),
		SearchKey: os.Getenv("AZURE_AI_SEARCH_KEY"),
		VisionKey: os.Getenv("AZURE_AI_VISION_KEY"),
		Endpoint:  os.Getenv("AZURE_ENDPOINT"),
	}
}
