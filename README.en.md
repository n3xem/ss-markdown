# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  All aboard!
</p>
  
&nbsp;

S.S. Markdown is a GitHub Action for multilingual deployment of Markdown files.

The following APIs are available for use:

- OpenAI
- DeepSeek (unconfirmed functionality)
- Google (Gemini) (unconfirmed functionality)

## Input

| Input | Description | Required | Default |
|-------|-------------|----------|---------|
| `file` | Path to the Markdown file to be translated | No | `README.md` |
| `openai-api-key` | OpenAI API key | No | - |
| `deepseek-api-key` | DeepSeek API key | No | - |
| `google-api-key` | Google API key | No | - |
| `google-model` | Google Generative AI model name | No | - |
| `openai-model` | OpenAI model name | No | - |
| `ss-model` | Configuration of the model provider to use ('openai' or 'deepseek' or 'google') | Yes | - |

## Usage Example

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