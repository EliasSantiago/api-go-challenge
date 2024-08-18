# Etapa de construção
FROM golang:1.23-alpine AS builder

# Instalar git para baixar dependências
RUN apk add --no-cache git

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copiar o go.mod e go.sum para baixar as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copiar o código-fonte para o contêiner
COPY . .

# Compilar a aplicação
RUN go build -o main cmd/main.go

# Etapa final
FROM alpine:latest

# Instalar bash
RUN apk add --no-cache bash

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copiar o binário da etapa de construção
COPY --from=builder /go/src/app/main .

# Copiar o script wait-for-it.sh
COPY wait-for-it.sh .

# Copiar o arquivo .env
COPY .env .

# Expor a porta 8081
EXPOSE 8081

# Definir o comando de entrada para executar a aplicação
CMD ["./wait-for-it.sh", "db:5432", "--", "./main"]