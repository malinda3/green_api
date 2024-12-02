# Этап сборки
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY backend/go.mod backend/go.sum ./

# Устанавливаем зависимости
RUN go mod download

# Копируем код приложения
COPY backend/ .

# Сборка бинарного файла
RUN go build -o main .

# Минимальный финальный образ
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированный файл
COPY --from=builder /app/main .

# Добавляем зависимости (например, сертификаты для HTTPS)
RUN apk --no-cache add ca-certificates

# Указываем команду для запуска
CMD ["./main"]
