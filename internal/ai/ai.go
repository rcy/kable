package ai

import (
	"oj/internal/config"

	openai "github.com/sashabaranov/go-openai"
)

type AI struct {
	Client *openai.Client
	Model  string
}

func New() *AI {
	apiKey := config.Getenv("OPENAI_API_KEY", "")
	baseURL := config.Getenv("OPENAI_BASE_URL", "https://api.openai.com/v1")
	model := config.Getenv("OPENAI_MODEL", "deepseek-chat")

	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = baseURL

	return &AI{
		Client: openai.NewClientWithConfig(cfg),
		Model:  model,
	}
}
