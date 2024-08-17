# Usar uma imagem base do Go
FROM golang:1.23-alpine

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /go/src/app

# Copiar o código-fonte para o contêiner
COPY . .

# Expor a porta 8081
EXPOSE 8081

# Compilar a aplicação
RUN go build -o main cmd/main.go

# Definir o comando de entrada para executar a aplicação
CMD ["./main"]