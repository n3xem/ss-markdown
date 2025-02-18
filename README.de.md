# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  Alle an Bord!
</p>
  
&nbsp;

S.S. Markdown ist ein GitHub-Action für die mehrsprachige Bereitstellung von Markdown-Dateien.

Die folgenden APIs können verwendet werden:

- OpenAI
- DeepSeek (Funktion noch nicht überprüft)
- Google (Gemini) (Funktion noch nicht überprüft)

## Eingabe

| Eingabe | Beschreibung | Erforderlich | Standard |
|---------|--------------|--------------|----------|
| `file` | Pfad zur Markdown-Datei, die übersetzt werden soll | Nein | `README.md` |
| `openai-api-key` | OpenAI API-Schlüssel | Nein | - |
| `deepseek-api-key` | DeepSeek API-Schlüssel | Nein | - |
| `google-api-key` | Google API-Schlüssel | Nein | - |
| `google-model` | Name des Google Generative AI-Modells | Nein | - |
| `openai-model` | Name des OpenAI-Modells | Nein | - |
| `ss-model` | Einstellung des zu verwendenden Modellanbieters ('openai' oder 'deepseek' oder 'google') | Ja | - |

## Anwendungsbeispiel

```yaml
name: Dokumente übersetzen
on:
  workflow_dispatch:

jobs:
  übersetzen:
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