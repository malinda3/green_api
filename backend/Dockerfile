FROM golang:1.23-alpine as builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

RUN go mod tidy

# Копируем весь проект в контейнер
COPY . .  # Убедитесь, что эта строка работает правильно

RUN go build -o /app/main /app/backend/main.go

CMD ["/app/main"]
