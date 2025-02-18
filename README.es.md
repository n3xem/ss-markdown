# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  ¡Todos a bordo!
</p>
  
&nbsp;

S.S. Markdown es una acción de GitHub para desplegar archivos Markdown en múltiples idiomas.

Se pueden usar las siguientes API:

- OpenAI
- DeepSeek (no verificado en funcionamiento)
- Google (Gemini) (no verificado en funcionamiento)

## Entrada

| Input | Descripción | Requerido | Predeterminado |
|-------|-------------|-----------|----------------|
| `file` | Ruta del archivo Markdown a traducir | No | `README.md` |
| `openai-api-key` | Clave API de OpenAI | No | - |
| `deepseek-api-key` | Clave API de DeepSeek | No | - |
| `google-api-key` | Clave API de Google | No | - |
| `google-model` | Nombre del modelo de IA generativa de Google | No | - |
| `openai-model` | Nombre del modelo de OpenAI | No | - |
| `ss-model` | Configuración del proveedor de modelo a utilizar ('openai' o 'deepseek' o 'google') | Sí | - |
| `languages` | Código de idioma para traducir (separados por comas) | No | `en,zh,fr,es,de,ko` |

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

## Excluir algunas frases de la traducción

Si hay texto que no desea insertar en el Markdown traducido, como enlaces a otros idiomas, puede rodearlo con las directivas `ss-markdown-ignore start/end` para que no sea traducido.

```markdown
Este texto será traducido.
El siguiente texto será ignorado por las directivas.(Quien esté leyendo el Markdown traducido, por favor, lea el original para entender lo que está ocurriendo)

Como la directiva ha terminado, este texto será traducido.
```