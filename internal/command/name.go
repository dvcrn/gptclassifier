package command

import (
	"github.com/dvcrn/gptclassifier/pkg/openai"
)

func Name(content string, client *openai.OpenAIClient, extension string, amount int, example string) ([]string, error) {
	return client.Name(content, extension, amount, example)
}
