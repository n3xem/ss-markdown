package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
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

var languages = map[string]string{
	"en": "English",
	"es": "Spanish",
	"fr": "French",
	"de": "German",
	"zh": "Chinese",
	"ko": "Korean",
}

func translateContent(content, targetLang string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}

	reqBody := OpenAIRequest{
		Model: "gpt-4",
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: fmt.Sprintf("You are a professional translator. Translate the following markdown content to %s. Preserve all markdown formatting.", languages[targetLang]),
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
	req.Header.Set("Authorization", "Bearer "+apiKey)

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

func processMarkdownFile(filePath string) error {
	// 既に翻訳ファイルの場合はスキップ
	baseName := filepath.Base(filePath)
	if strings.Contains(strings.TrimSuffix(baseName, ".md"), ".") {
		return nil
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// frontmatterとコンテンツを分離
	parts := bytes.SplitN(content, []byte("---\n"), 3)
	if len(parts) != 3 {
		return fmt.Errorf("invalid frontmatter format")
	}

	var metadata map[string]interface{}
	if err := yaml.Unmarshal(parts[1], &metadata); err != nil {
		return err
	}

	markdownContent := string(parts[2])

	// 各言語に翻訳
	for langCode := range languages {
		// 元の言語はスキップ
		if originalLang, ok := metadata["language"].(string); ok && originalLang == langCode {
			continue
		}

		dir := filepath.Dir(filePath)
		base := strings.TrimSuffix(filepath.Base(filePath), ".md")
		translatedPath := filepath.Join(dir, fmt.Sprintf("%s.%s.md", base, langCode))

		// 既存の翻訳ファイルをスキップ
		if _, err := os.Stat(translatedPath); err == nil {
			continue
		}

		// コンテンツを翻訳
		translatedContent, err := translateContent(markdownContent, langCode)
		if err != nil {
			fmt.Printf("Error translating to %s: %v\n", langCode, err)
			continue
		}

		// 新しいメタデータを作成
		newMetadata := make(map[string]interface{})
		for k, v := range metadata {
			newMetadata[k] = v
		}
		newMetadata["language"] = langCode

		// 翻訳ファイルを保存
		var buf bytes.Buffer
		buf.WriteString("---\n")
		yamlData, err := yaml.Marshal(newMetadata)
		if err != nil {
			return err
		}
		buf.Write(yamlData)
		buf.WriteString("---\n\n")
		buf.WriteString(translatedContent)

		if err := os.WriteFile(translatedPath, buf.Bytes(), 0644); err != nil {
			return err
		}

		fmt.Printf("Created translation: %s\n", translatedPath)
	}

	return nil
}

func main() {
	// リポジトリ内のすべてのMarkdownファイルを処理
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".md") && !strings.Contains(path, ".github") {
			if err := processMarkdownFile(path); err != nil {
				fmt.Printf("Error processing %s: %v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through files: %v\n", err)
		os.Exit(1)
	}
}
