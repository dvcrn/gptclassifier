package openai

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func (c *OpenAIClient) Classify(content string, options []string) (string, error) {
	optionsText := strings.Join(options, ", ")

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: fmt.Sprintf("Classify the given content into one of the following options: '%s'. Only respond with classification and NOTHING ELSE. Do NOT change the casing or add punctuation, ONLY respond with one of the options given.", optionsText),
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}
