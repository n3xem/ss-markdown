# S.S. Markdown

S.S. Markdown ist eine GitHub Action zur mehrsprachigen Bereitstellung von Markdown-Dateien.

Die folgenden APIs können verwendet werden:

- OpenAI
- DeepSeek
- Google (Gemini)

## Eingabe

| Eingabe | Beschreibung | Erforderlich | Standard |
|---------|--------------|--------------|----------|
| `file` | Pfad zur Markdown-Datei, die übersetzt werden soll | Nein | `README.md` |
| `openai-api-key` | OpenAI API-Schlüssel | Nein | - |
| `deepseek-api-key` | DeepSeek API-Schlüssel | Nein | - |
| `google-api-key` | Google API-Schlüssel | Nein | - |
| `google-model` | Name des Google Generative AI-Modells | Nein | - |
| `openai-model` | Name des OpenAI-Modells | Nein | - |
| `ss-model` | Konfiguration des zu verwendenden Modellanbieters ('openai' oder 'deepseek' oder 'google') | Ja | - |

## Anwendungsbeispiel

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