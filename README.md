# ss-markdown

ss-markdownは、Markdownファイルを多言語展開するためのGitHub Actionsです。

下記のAPIが使用できます。

- OpenAI
- DeepSeek
- Google(Gemini)

## 入力

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `file` | 翻訳するMarkdownファイルのパス | No | `README.md` |
| `openai-api-key` | OpenAI APIキー | No | - |
| `deepseek-api-key` | DeepSeek APIキー | No | - |
| `google-api-key` | Google APIキー | No | - |
| `google-model` | Google Generative AIモデル名 | No | - |
| `openai-model` | OpenAIモデル名 | No | - |
| `ss-model` | 使用するモデルプロバイダーの設定('openai' or 'deepseek' or 'google') | Yes | - |

## 使用例

```yaml
name: Translate Docs
on:
  workflow_dispatch:

jobs:
  translate:
    permissions:
      contents: write
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: n3xem/ss-markdown@v0.0.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: ${{ secrets.SS_MARKDOWN_OPENAI_GENERATIVE_MODEL }}
          ss-model: ${{ secrets.SS_MARKDOWN_MODEL }}
      - uses: EndBug/add-and-commit@v9
```
