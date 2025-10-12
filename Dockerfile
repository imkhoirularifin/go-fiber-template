ARG GO_VERSION=1.24.2-alpine3.20
ARG ALPINE_VERSION=3.21.3

# Build stage
FROM golang:${GO_VERSION} AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o main .

# Final stage
FROM alpine:${ALPINE_VERSION}
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs/swagger.json ./docs/swagger.json

RUN apk --no-cache add curl ca-certificates

EXPOSE 3000
CMD ["./main"]