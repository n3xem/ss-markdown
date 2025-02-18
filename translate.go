package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/n3xem/ss-markdown/model"
)

func processMarkdownFile(filePath string, outputDir string, translator model.TranslationClient, targetLangs []string) error {
	// 既に翻訳ファイルの場合はスキップ
	baseName := filepath.Base(filePath)
	if strings.Contains(strings.TrimSuffix(baseName, ".md"), ".") {
		return nil
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	markdownContent := string(content)

	// 指定された言語のみに翻訳
	for _, langCode := range targetLangs {
		if _, exists := model.Languages[langCode]; !exists {
			fmt.Printf("Warning: Unsupported language code '%s' - skipping\n", langCode)
			continue
		}

		base := strings.TrimSuffix(filepath.Base(filePath), ".md")
		translatedPath := filepath.Join(outputDir, fmt.Sprintf("%s.%s.md", base, langCode))

		// 出力ディレクトリが存在しない場合は作成
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}

		// コンテンツを翻訳
		translatedContent, err := translator.Translate(markdownContent, langCode)
		if err != nil {
			if strings.Contains(err.Error(), "429") {
				return fmt.Errorf("Rate limit exceeded. Please try again later. %v", err)
			}
			fmt.Printf("Error translating to %s: %v\n", langCode, err)
			continue
		}

		// 翻訳ファイルを保存
		var buf bytes.Buffer
		buf.WriteString(translatedContent)

		if err := os.WriteFile(translatedPath, buf.Bytes(), 0644); err != nil {
			return err
		}

		fmt.Printf("Created translation: %s\n", translatedPath)
		time.Sleep(5 * time.Second)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <markdown_file> [output_directory] [target_languages]")
		fmt.Println("If output_directory is not specified, the same directory as the input file will be used")
		fmt.Println("target_languages: Comma-separated language codes (e.g., 'en,zh,de'). If not specified, all supported languages will be used")
		os.Exit(1)
	}

	filePath := os.Args[1]
	outputDir := filepath.Dir(filePath)

	// デフォルトは全言語
	targetLangs := make([]string, 0, len(model.Languages))
	for lang := range model.Languages {
		targetLangs = append(targetLangs, lang)
	}

	if len(os.Args) > 2 {
		outputDir = os.Args[2]
	}

	if len(os.Args) > 3 {
		// カンマ区切りの言語コードを配列に分割
		targetLangs = strings.Split(os.Args[3], ",")
	}

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

	var translator model.TranslationClient

	// 環境変数からAPIキーとモデルを取得
	deepseekKey := os.Getenv("SS_MARKDOWN_DEEPSEEK_API_KEY")
	openaiKey := os.Getenv("SS_MARKDOWN_OPENAI_API_KEY")
	googleKey := os.Getenv("SS_MARKDOWN_GOOGLE_API_KEY")
	googleGenerativeModel := os.Getenv("SS_MARKDOWN_GOOGLE_GENERATIVE_MODEL")
	openaiGenerativeModel := os.Getenv("SS_MARKDOWN_OPENAI_GENERATIVE_MODEL")
	modelName := os.Getenv("SS_MARKDOWN_MODEL")

	// 使用するモデルに基づいてトランスレーターを初期化
	switch modelName {
	case "openai":
		if openaiKey == "" {
			fmt.Println("Error: SS_MARKDOWN_OPENAI_API_KEY is not set")
			os.Exit(1)
		}
		translator = model.NewOpenAITranslator(openaiKey, openaiGenerativeModel)
	case "deepseek":
		if deepseekKey == "" {
			fmt.Println("Error: SS_MARKDOWN_DEEPSEEK_API_KEY is not set")
			os.Exit(1)
		}
		translator = model.NewDeepseekTranslator(deepseekKey)
	case "google":
		if googleKey == "" {
			fmt.Println("Error: SS_MARKDOWN_GOOGLE_API_KEY is not set")
			os.Exit(1)
		}
		translator, err = model.NewGoogleTranslator(googleKey, googleGenerativeModel)
		if err != nil {
			fmt.Printf("Error: Failed to create Google translator: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Error: Invalid or missing SS_MARKDOWN_MODEL (must be 'openai' or 'deepseek' or 'google')")
		os.Exit(1)
	}

	if err := processMarkdownFile(filePath, outputDir, translator, targetLangs); err != nil {
		fmt.Printf("Error processing %s: %v\n", filePath, err)
		os.Exit(1)
	}
}
