
FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app .

FROM alpine:latest

COPY --from=builder /app/app /app

EXPOSE 8881

CMD ["/app"]
