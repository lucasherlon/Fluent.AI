## Cliente da API do Google implementado em Go usado para traduzir textos
- Requer a versão 1.22.3 da linguagem Go
- Para criar o executável do programa basta abrir o terminal na pasta do programa e digitar o seguinte comando:

  ```
  go build
  ```
- Utiliza o modelo Gemini 1.5
- Necessita de uma [chave do Google AI](https://aistudio.google.com/app/apikey) salva em uma variável de ambiente chamada GEMINI_KEY guardada em arquivo .env na pasta principal do projeto
- Traduz para o português brasileiro por padrão, mas pode receber outro idioma de saída passado como argumento de linha de comando:

    ```
    ./fluent_ai francês
    ```
## Interface
- Essa aplicação utiliza os frameworks [Bubble Tea](https://github.com/charmbracelet/bubbletea) e [Lipgloss](https://github.com/charmbracelet/lipgloss) para estilizar a interface de linha de comando.
