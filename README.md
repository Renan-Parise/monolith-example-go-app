# Exemplo de Aplicação Web Monolítica em Go

Este é um exemplo simples de uma aplicação web em monolito escrita em Go, utilizando o framework Gin para roteamento e templates HTML para renderização.

## Como Executar

Para executar esta aplicação, siga estes passos:

1. Certifique-se de ter o Go instalado em seu sistema.

2. Clone este repositório para sua máquina local usando o Git:

    ```
    git clone <repository-url>
    ```

3. Navegue até o diretório do projeto:

    ```
    cd <project-directory>
    ```

4. Instale o framework Gin executando o seguinte comando no seu terminal:

    ```
    go get -u github.com/gin-gonic/gin
    ```

5. Execute a aplicação executando o arquivo principal Go:

    ```
    go run main.go
    ```

6. Acesse no seu navegador [127.0.0.1:7000/users](127.0.0.1:7000/users) para ver todos os usuários, ou [127.0.0.1:7000/users/id](127.0.0.1:7000/id) para ver um específico.