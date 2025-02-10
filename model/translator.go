package model

var Languages = map[string]string{
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
