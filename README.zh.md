# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  全员上船！
</p>
  
&nbsp;

S.S. Markdown 是一个用于将 Markdown 文件进行多语言展现的 GitHub Actions。

下列 API 可供使用。

- OpenAI
- DeepSeek（功能未确认）
- Google（Gemini）（功能未确认）

## 输入

| 输入 | 描述 | 必需 | 默认 |
|-------|-------------|----------|---------|
| `file` | 要翻译的 Markdown 文件路径 | 否 | `README.md` |
| `openai-api-key` | OpenAI API 密钥 | 否 | - |
| `deepseek-api-key` | DeepSeek API 密钥 | 否 | - |
| `google-api-key` | Google API 密钥 | 否 | - |
| `google-model` | Google 生成 AI 模型名称 | 否 | - |
| `openai-model` | OpenAI 模型名称 | 否 | - |
| `ss-model` | 要使用的模型提供者的设置（'openai' 或 'deepseek' 或 'google'） | 是 | - |

## 使用示例

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