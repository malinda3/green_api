FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY backend/greenapi/go.mod ./backend/greenapi/

RUN go mod tidy

RUN mkdir -p /app/backend/greenapi

WORKDIR /app/backend/greenapi
RUN go mod tidy

COPY . ./

WORKDIR /app/backend
RUN ls

RUN go run main.go
