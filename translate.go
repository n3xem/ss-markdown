package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

var languages = map[string]string{
	"en": "English",
	"es": "Spanish",
	"fr": "French",
	"de": "German",
	"zh": "Chinese",
	"ko": "Korean",
}

// TranslationClient は翻訳サービスのインターフェース
type TranslationClient interface {
	Translate(content, targetLang string) (string, error)
}

func processMarkdownFile(filePath string, translator TranslationClient) error {
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
		translatedContent, err := translator.Translate(markdownContent, langCode)
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
	if len(os.Args) < 2 {
		fmt.Println("Error: Please specify a markdown file path")
		os.Exit(1)
	}

	filePath := os.Args[1]

	// ファイルの存在確認
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("Error: File not found or cannot be accessed: %v\n", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Println("Error: Specified path is a directory. Please specify a markdown file")
		os.Exit(1)
	}

	if !strings.HasSuffix(filePath, ".md") {
		fmt.Println("Error: Specified file is not a markdown file")
		os.Exit(1)
	}

	var translator TranslationClient

	// 環境変数からAPIキーとモデルを取得
	deepseekKey := os.Getenv("DEEPSEEK_API_KEY")
	openaiKey := os.Getenv("OPENAI_API_KEY")
	model := os.Getenv("SS_MODEL")

	// 使用するモデルに基づいてトランスレーターを初期化
	switch model {
	case "openai":
		if openaiKey == "" {
			fmt.Println("Error: OPENAI_API_KEY is not set")
			os.Exit(1)
		}
		translator = NewOpenAITranslator(openaiKey, "openai")
	case "deepseek":
		if deepseekKey == "" {
			fmt.Println("Error: DEEPSEEK_API_KEY is not set")
			os.Exit(1)
		}
		translator = NewDeepseekTranslator(deepseekKey, "deepseek")
	default:
		fmt.Println("Error: Invalid or missing SS_MODEL (must be 'openai' or 'deepseek')")
		os.Exit(1)
	}

	if err := processMarkdownFile(filePath, translator); err != nil {
		fmt.Printf("Error processing %s: %v\n", filePath, err)
		os.Exit(1)
	}
}
