# Authentication Service

Handles user authentication, authorization, and token management for the Edge Device Monitoring System.

## Architecture

The authentication service follows standard Go project layout:

```
.
├── cmd/
│   └── auth/          # Main application
│       └── main.go    # Entry point
├── internal/
│   ├── auth/          # Core authentication logic
│   │   ├── service.go # Authentication service
│   │   └── models.go  # Data models
│   └── server/        # HTTP server implementation
│       └── server.go  # HTTP handlers
└── pkg/               # Public libraries
    └── common/        # Shared utilities
```

## Features

- JWT-based authentication
- Role-based access control (RBAC)
- Secure password hashing
- Token refresh mechanism
- Rate limiting
- Audit logging

## API Endpoints

- `POST /auth/login` - User login
- `POST /auth/register` - User registration
- `POST /auth/refresh` - Refresh access token
- `GET /auth/verify` - Verify token
- `GET /auth/user` - Get user information

## Development

1. Setup:
```bash
cd services/auth
go mod download
```

2. Run Tests:
```bash
go test -v ./...
go test -race ./...
go test -cover ./...
```

3. Local Development:
```bash
# Run with hot reload
air -c .air.toml

# Or standard way
go run cmd/auth/main.go
```

## Building

### For ARM64 (Raspberry Pi):
```bash
# Using buildx
docker buildx build --platform linux/arm64 -t auth-service:latest --load .

# Using Go directly
GOOS=linux GOARCH=arm64 go build -o auth ./cmd/auth
```

### Multi-arch Build:
```bash
docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 \
-t myregistry/auth-service:latest --push .
```

## Configuration

Environment variables:
- `PORT` - Server port (default: 8080)
- `DB_DSN` - Database connection string
- `JWT_SECRET` - JWT signing key
- `TOKEN_EXPIRY` - Token expiration time
- `ENABLE_METRICS` - Enable Prometheus metrics
- `LOG_LEVEL` - Logging level (debug/info/warn/error)

## Testing

1. Unit Tests:
```bash
go test ./internal/auth -v
```

2. Integration Tests:
```bash
go test ./internal/server -v
```

3. Coverage Report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Monitoring

The service exposes metrics at `/metrics` in Prometheus format:
- Request latencies
- Error rates
- Active sessions
- Authentication attempts

## Security Considerations

- All passwords are hashed using bcrypt
- Tokens are signed with HMAC-SHA256
- Rate limiting prevents brute force attacks
- Environment variables for sensitive configuration
- CORS and security headers configured

## Contributing

1. Create a feature branch
2. Run tests and linters
3. Update documentation
4. Submit pull request

