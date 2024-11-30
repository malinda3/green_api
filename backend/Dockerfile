FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY backend/greenapi/go.mod ./backend/greenapi/

RUN go mod tidy

WORKDIR /app/backend/greenapi/backend/greenapi
RUN go mod tidy

COPY . ./

RUN ls -l /app/backend/greenapi

RUN go get github.com/rs/cors

RUN go run /app/backend/greenapi/backendmain.go
