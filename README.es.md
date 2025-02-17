# S.S. Markdown

S.S. Markdown es una acción de GitHub para desplegar archivos Markdown en múltiples idiomas.

Se pueden utilizar las siguientes APIs.

- OpenAI
- DeepSeek
- Google(Gemini)

## Entrada

| Entrada | Descripción | Requerido | Por defecto |
|---------|-------------|-----------|-------------|
| `file` | Ruta del archivo Markdown a traducir | No | `README.md` |
| `openai-api-key` | Clave API de OpenAI | No | - |
| `deepseek-api-key` | Clave API de DeepSeek | No | - |
| `google-api-key` | Clave API de Google | No | - |
| `google-model` | Nombre del modelo de IA generativa de Google | No | - |
| `openai-model` | Nombre del modelo de OpenAI | No | - |
| `ss-model` | Configuración del proveedor de modelo a utilizar ('openai' o 'deepseek' o 'google') | Sí | - |

## Ejemplo de uso

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