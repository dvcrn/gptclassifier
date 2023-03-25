package openai

import (
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func (c *OpenAIClient) Name(content string, extension string, amount int, example string) ([]string, error) {
	if example != "" {
		example = fmt.Sprintf("An example of how the names should look like is %s", example)
	}

	messageContent := fmt.Sprintf("Given the following content, generate %d (not more and not less) concise but descriptive file names that could fit this content. The file type is %s. Output the names comma-separated, in one line, and nothing else. DO NOT OUTPUT BULLETPOINTS OR A LIST. %s", amount, extension, example)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: messageContent,
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
		return nil, fmt.Errorf("ChatCompletion error: %v", err)
	}

	choices := strings.Split(resp.Choices[0].Message.Content, ",")
	for i, v := range choices {
		choices[i] = strings.TrimSpace(v)
	}
	return choices, nil
}
