# Base Go image
FROM golang:1.23.4-alpine as builder

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x brokerApp

# Tiny docker image
FROM alpine:latest

RUN mkdir -p /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]
