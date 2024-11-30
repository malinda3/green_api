FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .  # Копируем весь код

# Собираем приложение
RUN go build -o /app/main /app/backend/main.go

# Запускаем приложение
CMD ["/app/main"]
