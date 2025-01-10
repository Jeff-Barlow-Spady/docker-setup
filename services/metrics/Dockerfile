FROM golang:1.21-alpine as builder
WORKDIR /app
COPY . /app/
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o metrics ./metrics

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/metrics .
EXPOSE 8081
CMD ["./metrics"]
