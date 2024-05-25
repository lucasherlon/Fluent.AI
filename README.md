## Cliente da API do Google implementado em Go usado para traduzir textos
- requer a versão 1.22.3 da linguegem Go
- Para criar o executável do programa basta abrir o terminal na pasta do programa e digitar o seguinte comando:

  ```
  go build
  ```
- utiliza o modelo Gemini 1.5
- necessita de uma [chave do Google AI](https://aistudio.google.com/app/apikey) salva em uma variável de ambiente chamada GEMINI_KEY
- traduz para o português brasileiro por padrão, mas pode receber outro idioma de saída passado como argumento de linha de comando:
  
    ```
    ./fluent_ai francês
    ```
