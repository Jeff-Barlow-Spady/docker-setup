FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files first
COPY go.mod ./
# Copy go.sum if it exists
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]