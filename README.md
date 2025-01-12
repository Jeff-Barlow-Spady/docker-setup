# Edge Device Monitoring System

A comprehensive monitoring system for edge devices, built with Go and designed for ARM architectures (Raspberry Pi).

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
│   ├── auth/       # Authentication service
│   ├── gpio/       # GPIO control service
│   └── metrics/    # System metrics service
├── docs/           # Documentation
├── caddy/         # Reverse proxy configuration
└── docker/        # Docker configurations
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

1. Build for ARM:
```bash
make docker-build PLATFORMS=linux/arm/v7
```

2. Deploy:
```bash
docker compose up -d
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