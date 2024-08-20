# api-go-challenge

## Descrição

Este projeto é uma API desenvolvida em Go que utiliza PostgreSQL como banco de dados. O Docker Compose é utilizado para facilitar a configuração e inicialização dos serviços.

## Pré-requisitos

- [Go](https://golang.org/doc/install) (versão 1.16 ou superior)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Como Rodar a Aplicação

1. **Clone o repositório:**

    ```sh
        git clone https://github.com/EliasSantiago/api-go-challenge.git
        cd api-go-challenge
    ```

2. **Crie um arquivo `.env` com as seguintes configurações:**

    ```properties
        POSTGRES_USER=postgres
        POSTGRES_PASSWORD=postgres
        POSTGRES_DB=postgres
        DB_HOST=db
        DB_PORT=5432
        DB_USER=postgres
        DB_PASSWORD=postgres
        DB_NAME=postgres
    ```

3. **Inicie os serviços com Docker Compose:**

    ```sh
    docker-compose up -d
    ```

4. **Acesse a aplicação:**

    A API estará disponível em `http://localhost:8081`.

## Configuração do Banco de Dados

1. **Execute o script de criação das tabelas:**

    O script de criação das tabelas está localizado em `db/migrations/create_tables.sql`.

## Testes

1. **Para rodar os testes, use o seguinte comando:**

    ```sh
    go test ./...
    ```

## Collection do Postman

1. **Para facilitar o teste das rotas da API, uma collection do Postman está disponível no projeto.**
```plaintext
|-- /postman/api-go-challenge.postman_collection.json
```

## Estrutura do Projeto

```plaintext
/api-go-challenge
|-- /cmd
|   |-- main.go
|-- /controller
|-- /db
|   |-- migrations
|-- /postman
|-- /repository
|-- /routes
|-- /usecase
|-- go.mod
|-- go.sum