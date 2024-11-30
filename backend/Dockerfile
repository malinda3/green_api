FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY backend/greenapi/go.mod  ./backend/greenapi
RUN go mod tidy

WORKDIR /app/backend/greenapi
RUN go mod tidy

COPY . .  

RUN go build -o /app/main /app/backend/main.go

CMD ["/app/main"]
