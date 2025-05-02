# S.S. Markdown



&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  모두 탑승!
</p>
  
&nbsp;

S.S. Markdown은 Markdown 파일을 다국어로 전개하기 위한 GitHub Actions입니다.

다음 API를 사용할 수 있습니다.

- OpenAI
- DeepSeek(동작 미확인)
- Google(Gemini)(동작 미확인)

## 입력

| 입력 | 설명 | 필수 | 기본값 |
|-------|-------------|----------|---------|
| `file` | 번역할 Markdown 파일의 경로 | 아니오 | `README.md` |
| `openai-api-key` | OpenAI API 키 | 아니오 | - |
| `deepseek-api-key` | DeepSeek API 키 | 아니오 | - |
| `google-api-key` | Google API 키 | 아니오 | - |
| `google-model` | Google 생성 AI 모델 이름 | 아니오 | - |
| `openai-model` | OpenAI 모델 이름 | 아니오 | - |
| `ss-model` | 사용할 모델 공급자 설정('openai' 또는 'deepseek' 또는 'google') | 예 | - |
| `languages` | 번역할 언어 코드(쉼표로 구분) | 아니오 | `en,zh,fr,es,de,ko` |

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
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: "gpt-4o-mini"
          ss-model: "openai"
      - uses: EndBug/add-and-commit@v9
```

## 일부 문장을 번역에서 제외하기

각 언어로의 링크 등 번역된 Markdown에 삽입하고 싶지 않은 문장이 있는 경우, `ss-markdown-ignore start/end` 지시문으로 감싸서 번역되지 않도록 할 수 있습니다.

```markdown
여기 문장은 번역됩니다.
아래의 지시문에 의해 번역이 무시됩니다. (번역된 Markdown을 읽고 있는 사람은 원문을 읽어 상황을 확인하시기 바랍니다)

지시문이 종료되었으므로, 여기 문장은 번역됩니다.
```