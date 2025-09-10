FROM golang:1.17-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

FROM alpine:latest

RUN apk add --no-cache curl

WORKDIR /root/

COPY --from=builder /app/main .

RUN adduser -D -s /bin/sh appuser

RUN chown appuser:appuser main

USER appuser

EXPOSE 8081

CMD ["./main"]