# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  上船吧！
</p>

&nbsp;

S.S. Markdown 是一个用于在 GitHub Actions 中多语言展开 Markdown 文件的工具。

可以使用以下 API。

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
| `google-model` | Google 生成性 AI 模型名称 | 否 | - |
| `openai-model` | OpenAI 模型名称 | 否 | - |
| `ss-model` | 要使用的模型提供者设置（'openai' 或 'deepseek' 或 'google'） | 是 | - |
| `languages` | 要翻译的语言代码（以逗号分隔） | 否 | `en,zh,fr,es,de,ko` |

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
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: "gpt-4o-mini"
          ss-model: "openai"
      - uses: EndBug/add-and-commit@v9
```

## 从翻译中排除某些文本

如果有一些不想插入到翻译后的 Markdown 中的文本，如各语言的链接，可以通过 `ss-markdown-ignore start/end` 指令将其包围，以避免翻译。

```markdown
这里的文本将会被翻译。
以下指令将会忽略翻译。（阅读翻译后的 Markdown 的人请阅读原文以确认发生了什么）

指令已结束，所以这里的文本将会被翻译。
```