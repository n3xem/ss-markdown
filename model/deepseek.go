package model

import (
	"context"
	"fmt"
	"os"

	"github.com/cohesion-org/deepseek-go"
	"github.com/cohesion-org/deepseek-go/constants"
)

type DeepseekTranslator struct {
	apiKey string
}

func NewDeepseekTranslator(apiKey string) TranslationClient {
	return &DeepseekTranslator{
		apiKey: apiKey,
	}
}

func (t *DeepseekTranslator) Translate(content, targetLang string) (string, error) {
	// Set up the Deepseek client
	client := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))

	// Create a chat completion request
	request := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: constants.ChatMessageRoleSystem, Content: fmt.Sprintf("You are a professional translator. Translate the following markdown content to %s. Preserve all markdown formatting.", Languages[targetLang])},
			{Role: constants.ChatMessageRoleUser, Content: content},
		},
	}

	// Send the request and handle the response
	ctx := context.Background()
	response, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
}
