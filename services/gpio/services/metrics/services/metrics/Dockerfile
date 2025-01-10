# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

# Build the binary with ARM support
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 \
    go build -o metrics ./metrics

# Final stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/metrics .

EXPOSE 8000
CMD ["./metrics"]
