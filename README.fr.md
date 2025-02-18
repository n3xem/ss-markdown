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
| `ss-model` | Paramètre du fournisseur de modèle à utiliser ('openai' ou 'deepseek' ou 'google') | Oui | - |
| `languages` | Codes des langues à traduire (séparés par des virgules) | Non | `en,zh,fr,es,de,ko` |

## Exemples d'utilisation

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

## Exclure certaines phrases de la traduction

Si vous avez des phrases que vous ne souhaitez pas insérer dans le Markdown traduit, comme des liens vers chaque langue, vous pouvez les entourer avec la directive `ss-markdown-ignore start/end` pour qu'elles ne soient pas traduites.

```markdown
Cette phrase sera traduite.
La traduction sera ignorée par la directive ci-dessous. (Les personnes qui lisent le Markdown traduit peuvent consulter le texte original pour voir ce qui se passe)

La directive étant terminée, cette phrase sera traduite.
```