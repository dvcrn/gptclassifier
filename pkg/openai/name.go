package openai

import (
	"context"
	"fmt"
	"strings"

	"github.com/dvcrn/gptclassifier/internal/utils"
	openai "github.com/sashabaranov/go-openai"
)

func (c *OpenAIClient) Name(content string, extension string, amount int, example []string) ([]string, error) {
	examples := ""
	if len(example) > 0 {
		examples = fmt.Sprintf("Consider this list of example file names and follow similar naming: \n- %s", strings.Join(example, "\n- "))
	}

	messageContent := fmt.Sprintf("Given the following content sent in the next message, generate %d (not more and not less) concise but descriptive file names that could fit this content. The file type is %s. Output the names comma-separated, in one line, and nothing else. DO NOT OUTPUT BULLETPOINTS OR A LIST. Spaces are fine to include in the file name. %s", amount, extension, examples)

	promptTokenCount := utils.CountTokens(messageContent)

	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: messageContent,
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
		return nil, fmt.Errorf("ChatCompletion error: %v", err)
	}

	choices := strings.Split(resp.Choices[0].Message.Content, ",")
	for i, v := range choices {
		choices[i] = strings.TrimSpace(v)
	}
	return choices, nil
}
