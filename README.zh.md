# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  欢迎登船!
</p>
  
&nbsp;

S.S. Markdown是一个用于在GitHub Actions中多语言展开Markdown文件的工具。

以下API可供使用。

- OpenAI
- DeepSeek（未确认操作）
- Google（Gemini）（未确认操作）

## 输入

| 输入 | 描述 | 必需 | 默认 |
|------|------|------|------|
| `file` | 要翻译的Markdown文件路径 | 否 | `README.md` |
| `openai-api-key` | OpenAI API密钥 | 否 | - |
| `deepseek-api-key` | DeepSeek API密钥 | 否 | - |
| `google-api-key` | Google API密钥 | 否 | - |
| `google-model` | Google生成性AI模型名称 | 否 | - |
| `openai-model` | OpenAI模型名称 | 否 | - |
| `ss-model` | 使用的模型提供者设置（'openai'或'deepseek'或'google'） | 是 | - |

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