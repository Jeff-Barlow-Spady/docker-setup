# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

# Build the binary with specific tags for GPIO support
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 \
    go build -o gpio -tags=embedded ./gpio

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/gpio .

EXPOSE 8000
CMD ["./gpio"]
