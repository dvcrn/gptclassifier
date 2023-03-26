package openai

import (
	"context"
	"fmt"
	"strings"

	"github.com/dvcrn/gptclassifier/internal/utils"
	openai "github.com/sashabaranov/go-openai"
)

func (c *OpenAIClient) Classify(content string, options []string) (string, error) {
	optionsText := strings.Join(options, ", ")

	prompt := fmt.Sprintf("Classify the given content sent in the next message into one of the following options: '%s'. Only respond with classification and NOTHING ELSE. Do NOT change the casing or add punctuation, ONLY respond with one of the options given.", optionsText)
	promptTokenCount := utils.CountTokens(prompt)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: utils.SliceTokens(content, 2800-promptTokenCount),
		},
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Messages:    messages,
			Temperature: 0.5,
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}
