# S.S. Markdown

S.S. Markdown은 Markdown 파일을 다국어로 배포하기 위한 GitHub Actions입니다.

다음 API를 사용할 수 있습니다.

- OpenAI
- DeepSeek
- Google(Gemini)

## 입력

| 입력 | 설명 | 필수 | 기본값 |
|-------|-------------|----------|---------|
| `file` | 번역할 Markdown 파일의 경로 | 아니오 | `README.md` |
| `openai-api-key` | OpenAI API 키 | 아니오 | - |
| `deepseek-api-key` | DeepSeek API 키 | 아니오 | - |
| `google-api-key` | Google API 키 | 아니오 | - |
| `google-model` | Google Generative AI 모델 이름 | 아니오 | - |
| `openai-model` | OpenAI 모델 이름 | 아니오 | - |
| `ss-model` | 사용할 모델 제공자의 설정('openai' 또는 'deepseek' 또는 'google') | 예 | - |

## 사용 예

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