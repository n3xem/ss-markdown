package model

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GoogleTranslator struct {
	client *genai.Client
	model  string
}

func NewGoogleTranslator(apiKey string, model string) (TranslationClient, error) {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	return &GoogleTranslator{
		client: client,
		model:  model,
	}, nil
}

func (t *GoogleTranslator) Translate(content, targetLang string) (string, error) {
	model := t.client.GenerativeModel(t.model)
	prompt := fmt.Sprintf("Translate the following markdown content to %s. Preserve all markdown formatting:\n\n%s", Languages[targetLang], content)
	resp, err := model.GenerateContent(context.Background(), genai.Text(prompt))
	if err != nil {
		return "", err
	}
	return parseResponse(resp), nil
}

func parseResponse(resp *genai.GenerateContentResponse) string {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if textPart, ok := part.(genai.Text); ok {
					return string(textPart)
				}
			}
		}
	}
	return ""
}
