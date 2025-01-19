FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

# Add build arguments for cross-compilation
ARG TARGETARCH
ARG TARGETOS
ARG TARGETPLATFORM
ARG BUILDPLATFORM

WORKDIR /app

# Add build tools and enable Go modules
RUN apk add --no-cache git
ENV GO111MODULE=on

# Copy only the module files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source tree
COPY . .

# Build the application with proper arch support
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -v -o /app/auth ./cmd/auth

# Create minimal production image
FROM --platform=$TARGETPLATFORM alpine:latest

WORKDIR /app
COPY --from=builder /app/auth .

# Configure container
EXPOSE 8000
CMD ["./auth"]
