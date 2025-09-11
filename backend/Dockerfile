FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o main ./cmd/api

FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=builder /app/main /main

EXPOSE 8081

USER nonroot:nonroot

ENTRYPOINT ["/main"]