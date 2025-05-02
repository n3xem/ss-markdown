# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  All aboard!
</p>
  
&nbsp;

S.S. Markdown is a GitHub Actions tool for multilingual deployment of Markdown files.

The following APIs can be used:

- OpenAI
- DeepSeek (unverified operation)
- Google (Gemini) (unverified operation)

## Input

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `file` | Path to the Markdown file to be translated | No | `README.md` |
| `openai-api-key` | OpenAI API key | No | - |
| `deepseek-api-key` | DeepSeek API key | No | - |
| `google-api-key` | Google API key | No | - |
| `google-model` | Google Generative AI model name | No | - |
| `openai-model` | OpenAI model name | No | - |
| `ss-model` | Settings for the model provider to use ('openai' or 'deepseek' or 'google') | Yes | - |
| `languages` | Language codes for translation (comma-separated) | No | `en,zh,fr,es,de,ko` |

## Example Usage

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
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: "gpt-4o-mini"
          ss-model: "openai"
      - uses: EndBug/add-and-commit@v9
```

## Excluding Certain Texts from Translation

If there are texts that you do not want to be included in the translated Markdown, such as links to each language, you can surround them with the `ss-markdown-ignore start/end` directive to exclude them from translation.

```markdown
This text will be translated.
The following directive will cause the translation to be ignored. (Readers of the translated Markdown should refer to the original text to see what is happening)

Since the directive has ended, this text will be translated.
```