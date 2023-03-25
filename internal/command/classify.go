package command

import (
	"github.com/dvcrn/gptclassifier/pkg/openai"
)

func Classify(content string, client *openai.OpenAIClient, options []string) (string, error) {
	return client.Classify(content, options)
}
