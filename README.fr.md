# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  Tous à bord !
</p>

&nbsp;

S.S. Markdown est une action GitHub pour déployer des fichiers Markdown en plusieurs langues.

Les API suivantes peuvent être utilisées.

- OpenAI
- DeepSeek (fonctionnement non vérifié)
- Google (Gemini) (fonctionnement non vérifié)

## Entrée

| Input | Description | Requis | Par défaut |
|-------|-------------|--------|------------|
| `file` | Chemin du fichier Markdown à traduire | Non | `README.md` |
| `openai-api-key` | Clé API OpenAI | Non | - |
| `deepseek-api-key` | Clé API DeepSeek | Non | - |
| `google-api-key` | Clé API Google | Non | - |
| `google-model` | Nom du modèle d'IA générative de Google | Non | - |
| `openai-model` | Nom du modèle OpenAI | Non | - |
| `ss-model` | Configuration du fournisseur de modèle à utiliser ('openai' ou 'deepseek' ou 'google') | Oui | - |

## Exemple d'utilisation

```yaml
name: Traduire les docs
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