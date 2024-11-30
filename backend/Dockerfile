FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . ./

RUN ls -l /app/backend/greenapi

RUN go get github.com/rs/cors

RUN go run /app/backend/greenapi/backend/main.go
