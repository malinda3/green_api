# Стадия сборки
FROM golang:1.23-alpine as builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# Сборка бинарного файла
RUN go build -o /app/server /app/main.go

# Стадия запуска
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/server /app/server

# Установить права на выполнение
RUN chmod +x /app/server

# Указать открытые порты
EXPOSE 8881

CMD ["/app/server"]
