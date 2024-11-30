FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . ./

RUN go get github.com/rs/cors

RUN ls -l /app/backend/greenapi && ls -l frontend/

CMD ["go", "run", "/app/backend/main.go"]
