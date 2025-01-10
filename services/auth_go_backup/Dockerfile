FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . /app/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./auth

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/auth .
EXPOSE 8080
CMD ["./auth"]
