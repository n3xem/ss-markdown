# S.S. Markdown

S.S. Markdown est une action GitHub pour déployer des fichiers Markdown en plusieurs langues.

Les API suivantes peuvent être utilisées :

- OpenAI
- DeepSeek
- Google (Gemini)

## Entrée

| Input                  | Description                                                 | Required | Default      |
|-----------------------|-------------------------------------------------------------|----------|--------------|
| `file`                | Chemin du fichier Markdown à traduire                      | Non      | `README.md`  |
| `openai-api-key`      | Clé API OpenAI                                             | Non      | -            |
| `deepseek-api-key`    | Clé API DeepSeek                                         | Non      | -            |
| `google-api-key`      | Clé API Google                                           | Non      | -            |
| `google-model`        | Nom du modèle d'IA générative Google                      | Non      | -            |
| `openai-model`        | Nom du modèle OpenAI                                      | Non      | -            |
| `ss-model`            | Configuration du fournisseur de modèle à utiliser ('openai' ou 'deepseek' ou 'google') | Oui      | -            |

## Exemple d'utilisation

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