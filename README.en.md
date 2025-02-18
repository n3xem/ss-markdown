# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  All aboard!
</p>
  
&nbsp;

S.S. Markdown is a GitHub Actions tool for multi-language deployment of Markdown files.

The following APIs are available:

- OpenAI
- DeepSeek (Functionality not verified)
- Google (Gemini) (Functionality not verified)

## Input

| Input                   | Description                                           | Required | Default               |
|------------------------|-------------------------------------------------------|----------|-----------------------|
| `file`                 | Path to the Markdown file to be translated            | No       | `README.md`           |
| `openai-api-key`       | OpenAI API key                                       | No       | -                     |
| `deepseek-api-key`     | DeepSeek API key                                     | No       | -                     |
| `google-api-key`       | Google API key                                       | No       | -                     |
| `google-model`         | Google Generative AI model name                       | No       | -                     |
| `openai-model`         | OpenAI model name                                    | No       | -                     |
| `ss-model`             | Configuration for the model provider to use ('openai' or 'deepseek' or 'google') | Yes      | -                     |
| `languages`            | Language codes to translate into (comma-separated)   | No       | `en,zh,fr,es,de,ko`   |

## Examples

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

## Excluding certain sentences from translation

If there are sentences that you do not want to include in the translated Markdown, such as links to each language, you can prevent them from being translated by surrounding them with the `ss-markdown-ignore start/end` directive.

```markdown
This sentence will be translated.
The following directive will ignore the translation. (Readers of the translated Markdown should read the original text to see what is happening.)

The directive has ended, so this sentence will be translated.
```