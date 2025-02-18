# S.S. Markdown

<!-- ss-markdown-ignore start -->
[English](README.en.md) | [简体中文](README.zh.md) | [Español](README.es.md) | [Français](README.fr.md) | [Deutsch](README.de.md) | [한국어](README.ko.md)
<!-- ss-markdown-ignore end -->

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  All aboard!
</p>
  
&nbsp;

S.S. Markdownは、Markdownファイルを多言語展開するためのGitHub Actionsです。

下記のAPIが使用できます。

- OpenAI
- DeepSeek(動作未確認)
- Google(Gemini)(動作未確認)

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
| `languages` | 翻訳する言語コード(カンマ区切り) | No | `en,zh,fr,es,de,ko` |

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
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: ${{ secrets.SS_MARKDOWN_OPENAI_GENERATIVE_MODEL }}
          ss-model: ${{ secrets.SS_MARKDOWN_MODEL }}
      - uses: EndBug/add-and-commit@v9
```

## 一部の文章を翻訳から除外する

各言語へのリンクなど、翻訳されたマークダウンに挿入したくない文章がある場合は、`ss-markdown-ignore start/end` ディレクティブで囲むことで翻訳されないようにすることができます。

```markdown
ここの文章は翻訳されます。
下記のディレクティブによって翻訳が無視されます。(翻訳されたマークダウンを読んでいる人は、原文を読んで何が起きているか確認してください)
<!-- ss-markdown-ignore start -->
ここは翻訳されません。
<!-- ss-markdown-ignore end -->
ディレクティブが終了したので、ここの文章は翻訳されます。
```
