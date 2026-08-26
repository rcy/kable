package ai

import (
	"oj/internal/config"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

var apiKey = config.MustGetenv("DEEPSEEK_API_KEY")

type AI struct {
	Client *openai.Client
}

func New() *AI {
	client := openai.NewClient(
		option.WithBaseURL("https://api.deepseek.com/v1"),
		option.WithAPIKey(os.Getenv("DEEPSEEK_API_KEY")),
	)
	return &AI{Client: &client}
}
