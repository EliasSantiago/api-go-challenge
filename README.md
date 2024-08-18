# api-go-challenge

## Introdução

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

2. **Inicie os serviços com Docker Compose:**

    ```sh
    docker-compose up -d
    ```

3. **Acesse a aplicação:**

    A API estará disponível em `http://localhost:8081`.

## Estrutura do Projeto

```plaintext
/api-go-challenge
|-- /cmd
|   |-- main.go
|-- /controller
|-- /db
|-- /repository
|-- /routes
|-- /usecase
|-- go.mod
|-- go.sum