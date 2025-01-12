FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files first
COPY services/gpio-service/go.mod ./
# Copy go.sum if it exists
COPY services/gpio-service/go.sum ./

# Download dependencies
RUN go mod download

# Copy the service code
COPY services/gpio-service/ .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]