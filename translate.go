package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/n3xem/ss-markdown/model"
	"github.com/n3xem/ss-markdown/util"
)

// 設定構造体を追加
type Config struct {
	filePath     string
	outputDir    string
	targetLangs  []string
	modelName    string
	openAIConfig struct {
		apiKey          string
		generativeModel string
	}
	deepseekConfig struct {
		apiKey string
	}
	googleConfig struct {
		apiKey          string
		generativeModel string
	}
}

// 設定を読み込む関数
func loadConfig() (*Config, error) {
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("Usage: program <markdown_file> [output_directory] [target_languages]\n" +
			"If output_directory is not specified, the same directory as the input file will be used\n" +
			"target_languages: Comma-separated language codes (e.g., 'en,zh,de'). If not specified, all supported languages will be used")
	}

	config := &Config{
		filePath:  os.Args[1],
		outputDir: filepath.Dir(os.Args[1]),
	}

	// 出力ディレクトリの設定
	if len(os.Args) > 2 {
		config.outputDir = os.Args[2]
	}

	// 対象言語の設定
	config.targetLangs = make([]string, 0, len(model.Languages))
	for lang := range model.Languages {
		config.targetLangs = append(config.targetLangs, lang)
	}
	if len(os.Args) > 3 {
		config.targetLangs = strings.Split(os.Args[3], ",")
	}

	// 環境変数から設定を読み込み
	config.modelName = os.Getenv("SS_MARKDOWN_MODEL")
	config.openAIConfig.apiKey = os.Getenv("SS_MARKDOWN_OPENAI_API_KEY")
	config.openAIConfig.generativeModel = os.Getenv("SS_MARKDOWN_OPENAI_GENERATIVE_MODEL")
	config.deepseekConfig.apiKey = os.Getenv("SS_MARKDOWN_DEEPSEEK_API_KEY")
	config.googleConfig.apiKey = os.Getenv("SS_MARKDOWN_GOOGLE_API_KEY")
	config.googleConfig.generativeModel = os.Getenv("SS_MARKDOWN_GOOGLE_GENERATIVE_MODEL")

	return config, nil
}

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

	// Remove content between ignore directives
	startTag := "<!-- ss-markdown-ignore start -->"
	endTag := "<!-- ss-markdown-ignore end -->"
	markdownContent = util.RemoveTaggedContent(markdownContent, startTag, endTag)

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
	config, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ファイルの存在確認
	fileInfo, err := os.Stat(config.filePath)
	if err != nil {
		fmt.Printf("Error: File not found or cannot be accessed: %v\n", err)
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		fmt.Println("Error: Specified path is a directory. Please specify a markdown file")
		os.Exit(1)
	}

	if !strings.HasSuffix(config.filePath, ".md") {
		fmt.Println("Error: Specified file is not a markdown file")
		os.Exit(1)
	}

	var translator model.TranslationClient

	// 使用するモデルに基づいてトランスレーターを初期化
	switch config.modelName {
	case "openai":
		if config.openAIConfig.apiKey == "" {
			fmt.Println("Error: SS_MARKDOWN_OPENAI_API_KEY is not set")
			os.Exit(1)
		}
		translator = model.NewOpenAITranslator(config.openAIConfig.apiKey, config.openAIConfig.generativeModel)
	case "deepseek":
		if config.deepseekConfig.apiKey == "" {
			fmt.Println("Error: SS_MARKDOWN_DEEPSEEK_API_KEY is not set")
			os.Exit(1)
		}
		translator = model.NewDeepseekTranslator(config.deepseekConfig.apiKey)
	case "google":
		if config.googleConfig.apiKey == "" {
			fmt.Println("Error: SS_MARKDOWN_GOOGLE_API_KEY is not set")
			os.Exit(1)
		}
		translator, err = model.NewGoogleTranslator(config.googleConfig.apiKey, config.googleConfig.generativeModel)
		if err != nil {
			fmt.Printf("Error: Failed to create Google translator: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Error: Invalid or missing SS_MARKDOWN_MODEL (must be 'openai' or 'deepseek' or 'google')")
		os.Exit(1)
	}

	if err := processMarkdownFile(config.filePath, config.outputDir, translator, config.targetLangs); err != nil {
		fmt.Printf("Error processing %s: %v\n", config.filePath, err)
		os.Exit(1)
	}
}
