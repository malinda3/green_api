FROM golang:1.23-alpine as builder

WORKDIR /app

# Копируем все файлы в контейнер
COPY . ./

# Устанавливаем зависимости, включая cors
RUN go get github.com/rs/cors

# Проверка структуры каталогов
RUN ls -l /app/backend/greenapi

# Запуск приложения
CMD ["go", "run", "/app/backend/main.go"]
