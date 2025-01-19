# Edge Device Monitoring System

A comprehensive monitoring system for edge devices, built with Go and optimized for ARM64 architectures (Raspberry Pi). The system follows standard Go project layout and provides robust monitoring capabilities for edge devices.

## Architecture Overview

The system is built using a microservices architecture with three core services:
- Authentication Service: Handles user authentication and authorization
- GPIO Service: Manages hardware interaction and monitoring
- Metrics Service: Collects and exposes system metrics

Each service is independently deployable and follows the same structure:
```
services/
├── cmd/            # Main applications
├── internal/       # Private application code
├── pkg/            # Public libraries
└── api/            # API definitions
```

## Features

- Real-time GPIO monitoring and control
- System metrics collection (CPU, Memory, Disk)
- Service health monitoring
- JWT-based authentication
- Prometheus metrics integration
- Cross-platform support (x86_64, ARM)

## Prerequisites

- Docker and Docker Compose
- Go 1.22 or later
- Make
- Git

## Quick Start

1. Clone the repository:
```bash
git clone https://github.com/your-username/edge-device-monitoring.git
cd edge-device-monitoring
```

2. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

3. Build and run:
```bash
make dev-setup  # Install development tools
make all       # Build, test, and verify
make run       # Start services
```

## Development

### Project Structure

```
.
├── services/
│   ├── auth/              # Authentication service
│   │   ├── cmd/
│   │   │   └── auth/      # Main application
│   │   ├── internal/      # Private implementation
│   │   │   ├── auth/      # Auth logic
│   │   │   └── server/    # HTTP server
│   │   └── pkg/           # Public packages
│   ├── gpio/              # GPIO control service
│   │   ├── cmd/gpio/      # Main application
│   │   └── internal/      # Private implementation
│   └── metrics/           # System metrics service
│       ├── cmd/metrics/   # Main application
│       └── internal/      # Private implementation
├── docs/                  # Documentation
├── caddy/                 # Reverse proxy configuration
└── docker/                # Docker configurations
```

### Building

```bash
make build              # Build all services
make docker-build       # Build Docker images
```

### Testing

```bash
make test              # Run all tests
make coverage          # Generate coverage reports
```

### Code Quality

```bash
make lint              # Run linters
make fmt              # Format code
```

## Deployment

### Production

1. Configure Builder (one-time setup):
```bash
docker buildx create --name arm64builder --driver docker-container --platform linux/amd64,linux/arm64,linux/arm/v7 --use
docker buildx inspect arm64builder --bootstrap
```

2. Build for ARM64:
```bash
# Build all services
make docker-build PLATFORMS=linux/arm64

# Build specific service
docker buildx build --platform linux/arm64 -t myregistry/auth-service:latest --push ./services/auth
```

3. Deploy:
```bash
# Using docker compose
docker compose up -d

# Or individual services
docker run -d --name auth-service -p 8080:8080 myregistry/auth-service:latest
```

4. Verify Deployment:
```bash
# Check service status
docker ps

# View logs
docker logs -f auth-service
```

### Development

```bash
make run
```

## Service Architecture

### Authentication Service
- JWT-based authentication
- User management
- Role-based access control

### GPIO Service
- Real-time GPIO monitoring
- WebSocket support for live updates
- Hardware abstraction layer

### Metrics Service
- System metrics collection
- Prometheus integration
- Health check endpoints

## API Documentation

### REST Endpoints

- `POST /auth/login` - Authenticate user
- `GET /metrics` - Retrieve system metrics
- `GET /health` - Service health status
- `POST /gpio/:pin/write` - Control GPIO pins

### WebSocket API

- `ws://host/ws/gpio` - Real-time GPIO updates

## Contributing

1. Fork the repository
2. Create your feature branch
3. Run tests and linters
4. Submit a pull request

## Best Practices

### Code Style
- Follow Go standard project layout
- Use golangci-lint for code quality
- Write comprehensive tests
- Document public APIs

### Security
- No hardcoded credentials
- Use environment variables for configuration
- Regular dependency updates
- Proper error handling

### Operations
- Graceful shutdown support
- Proper logging
- Health checks
- Resource cleanup

## License

MIT License - see LICENSE file

## Support

For support, please open an issue in the GitHub repository.