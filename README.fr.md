# S.S. Markdown

&nbsp;
<p align="center">
  <img src="https://github.com/user-attachments/assets/dab375e4-f973-41dd-bf26-1ff34231af8c"><br>
  Tous à bord !
</p>
  
&nbsp;

S.S. Markdown est une action GitHub pour déployer des fichiers Markdown dans plusieurs langues.

Les API suivantes sont disponibles.

- OpenAI
- DeepSeek (fonctionnement non confirmé)
- Google (Gemini) (fonctionnement non confirmé)

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
| `languages` | Codes des langues à traduire (séparés par des virgules) | Non | `en,zh,fr,es,de,ko` |

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
      - uses: n3xem/ss-markdown@v0.2.1
        with:
          file: "README.md"
          openai-api-key: ${{ secrets.SS_MARKDOWN_OPENAI_API_KEY }}
          openai-model: "gpt-4o-mini"
          ss-model: "openai"
      - uses: EndBug/add-and-commit@v9
```

## Exclure certaines phrases de la traduction

S'il y a des phrases que vous ne souhaitez pas insérer dans le Markdown traduit, comme des liens vers des langues, vous pouvez utiliser les directives `ss-markdown-ignore start/end` pour éviter qu'elles ne soient traduites.

```markdown
Cette phrase ici sera traduite.
La traduction sera ignorée par la directive ci-dessous. (Les lecteurs du Markdown traduit peuvent lire le texte original pour comprendre ce qui se passe)

La directive étant terminée, cette phrase ici sera traduite.
```