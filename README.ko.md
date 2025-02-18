# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  모두 탑승하세요!
</p>

&nbsp;

S.S. Markdown은 Markdown 파일을 다국어로 전개하기 위한 GitHub Actions입니다.

아래의 API를 사용할 수 있습니다.

- OpenAI
- DeepSeek(작동 미확인)
- Google(Gemini)(작동 미확인)

## 입력

| Input | 설명 | 필수 | 기본값 |
|-------|-------------|----------|---------|
| `file` | 번역할 Markdown 파일의 경로 | 아니요 | `README.md` |
| `openai-api-key` | OpenAI API 키 | 아니요 | - |
| `deepseek-api-key` | DeepSeek API 키 | 아니요 | - |
| `google-api-key` | Google API 키 | 아니요 | - |
| `google-model` | Google Generative AI 모델 이름 | 아니요 | - |
| `openai-model` | OpenAI 모델 이름 | 아니요 | - |
| `ss-model` | 사용할 모델 제공자의 설정('openai' 또는 'deepseek' 또는 'google') | 예 | - |
| `languages` | 번역할 언어 코드(쉼표로 구분) | 아니요 | `en,zh,fr,es,de,ko` |

## 사용 예시

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

## 일부 문장을 번역에서 제외하기

각 언어로의 링크 등 번역된 Markdown에 삽입하고 싶지 않은 문장이 있는 경우 `ss-markdown-ignore start/end` 지시어로 둘러싸서 번역되지 않도록 할 수 있습니다.

```markdown
여기 문장은 번역됩니다.
아래의 지시어에 의해 번역이 무시됩니다. (번역된 Markdown을 읽고 있는 사람은 원문을 읽고 무슨 일이 일어나고 있는지 확인하세요)

지시어가 종료되어서, 여기 문장은 번역됩니다.
```