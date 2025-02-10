package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAIRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// OpenAITranslator はOpenAIを使用する翻訳クライアント
type OpenAITranslator struct {
	apiKey string
	model  string
}

// NewOpenAITranslator creates a new OpenAI translator instance
func NewOpenAITranslator(apiKey string, model string) *OpenAITranslator {
	return &OpenAITranslator{
		apiKey: apiKey,
		model:  model,
	}
}

// Translate implements TranslationClient interface
func (t *OpenAITranslator) Translate(content, targetLang string) (string, error) {
	reqBody := OpenAIRequest{
		Model: t.model,
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: fmt.Sprintf("You are a professional translator. Translate the following markdown content to %s. Preserve all markdown formatting.", Languages[targetLang]),
			},
			{
				Role:    "user",
				Content: content,
			},
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+t.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) == 0 {
		return "", fmt.Errorf("no translation result")
	}

	return result.Choices[0].Message.Content, nil
}
