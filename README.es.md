# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  ¡Todos a bordo!
</p>
  
&nbsp;

S.S. Markdown es un GitHub Actions para desplegar archivos Markdown en múltiples idiomas.

Se pueden utilizar las siguientes API:

- OpenAI
- DeepSeek (sin verificación de funcionamiento)
- Google (Gemini) (sin verificación de funcionamiento)

## Entrada

| Entrada | Descripción | Requerido | Predeterminado |
|---------|-------------|-----------|----------------|
| `file` | Ruta del archivo Markdown a traducir | No | `README.md` |
| `openai-api-key` | Clave de API de OpenAI | No | - |
| `deepseek-api-key` | Clave de API de DeepSeek | No | - |
| `google-api-key` | Clave de API de Google | No | - |
| `google-model` | Nombre del modelo de IA Generativa de Google | No | - |
| `openai-model` | Nombre del modelo de OpenAI | No | - |
| `ss-model` | Configuración del proveedor de modelos a utilizar ('openai' o 'deepseek' o 'google') | Sí | - |
| `languages` | Códigos de los idiomas a traducir (separados por comas) | No | `en,zh,fr,es,de,ko` |

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
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: "gpt-4o-mini"
          ss-model: "openai"
      - uses: EndBug/add-and-commit@v9
```

## Excluir partes del texto de la traducción

Si hay oraciones que no deseas incluir en el Markdown traducido, como enlaces a diferentes idiomas, puedes rodearlas con la directiva `ss-markdown-ignore start/end` para que no sean traducidas.

```markdown
Este texto será traducido.
Las oraciones que se rodeen con la siguiente directiva serán ignoradas en la traducción. (Los que leen el Markdown traducido, por favor, lean el texto original para entender qué está ocurriendo)

La directiva ha terminado, así que este texto será traducido.
```