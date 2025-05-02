# S.S. Markdown



&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  Alle an Bord!
</p>
  
&nbsp;

S.S. Markdown ist eine GitHub Action zur mehrsprachigen Bereitstellung von Markdown-Dateien.

Die folgenden APIs können genutzt werden:

- OpenAI
- DeepSeek (Funktion nicht bestätigt)
- Google (Gemini) (Funktion nicht bestätigt)

## Eingabe

| Eingabe | Beschreibung | Erforderlich | Standard |
|---------|--------------|--------------|----------|
| `file` | Pfad zur Markdown-Datei, die übersetzt werden soll | Nein | `README.md` |
| `openai-api-key` | OpenAI API-Schlüssel | Nein | - |
| `deepseek-api-key` | DeepSeek API-Schlüssel | Nein | - |
| `google-api-key` | Google API-Schlüssel | Nein | - |
| `google-model` | Name des Google Generative AI-Modells | Nein | - |
| `openai-model` | Name des OpenAI-Modells | Nein | - |
| `ss-model` | Einstellung des verwendeten Modellanbieters ('openai' oder 'deepseek' oder 'google') | Ja | - |
| `languages` | Sprachcodes zur Übersetzung (kommagetrennt) | Nein | `en,zh,fr,es,de,ko` |

## Anwendungsbeispiel

```yaml
name: Dokumente übersetzen
on:
  workflow_dispatch:

jobs:
  translate:
    permissions:
      contents: write
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: "gpt-4o-mini"
          ss-model: "openai"
      - uses: EndBug/add-and-commit@v9
```

## Teile des Textes von der Übersetzung ausschließen

Wenn es Texte gibt, die nicht in die übersetzten Markdown-Dokumente aufgenommen werden sollen, beispielsweise Links zu verschiedenen Sprachen, können Sie diese durch die Direktive `ss-markdown-ignore start/end` von der Übersetzung ausschließen.

```markdown
Dieser Text wird übersetzt.
Der folgende Direktive wird die Übersetzung ignoriert. (Leser der übersetzten Markdown sollten den Originaltext lesen, um zu verstehen, was passiert)

Die Direktive ist beendet, daher wird dieser Text wieder übersetzt.
```