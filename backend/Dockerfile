FROM golang:1.23-alpine as builder

WORKDIR /app

# Копируем go.mod и go.sum из корня проекта
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

# Теперь копируем весь код в контейнер
COPY . .

# Строим проект
RUN go build -o /app/main /app/backend/main.go

CMD ["/app/main"]
