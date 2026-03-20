# Estágio de Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copiamos os arquivos de módulo da RAIZ do projeto para o container
# No docker-compose, o 'context' deve ser a raiz (.)
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo o resto (incluindo a pasta /src)
COPY . .

# Compila o binário apontando para onde o main.go está (dentro de src)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./src/main.go

# Estágio Final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8000
CMD ["./main"]
