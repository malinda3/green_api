FROM golang:1.23-alpine as builder

WORKDIR /app

# Копируем go.mod и go.sum в контейнер
COPY go.mod go.sum ./

# Копируем go.mod и go.sum из backend/greenapi (если они есть)
COPY backend/greenapi/go.mod ./backend/greenapi/

# Выполняем go mod tidy для основного проекта
RUN go mod tidy

# Создаем директорию для backend/greenapi
RUN mkdir -p /app/backend/greenapi

# Выполняем go mod tidy для зависимостей в backend/greenapi
WORKDIR /app/backend/greenapi
RUN go mod tidy

# Копируем весь код в контейнер
COPY . .

# Делаем вывод содержимого директории перед запуском go run
WORKDIR /app/backend
RUN ls -l

# Теперь выполняем go run main.go
RUN go run main.go
