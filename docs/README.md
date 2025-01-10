# Edge Device Services Documentation

## Overview

This is a lightweight edge device server designed for Raspberry Pi Zero 2 W, providing:
- GPIO control via WebSocket and REST APIs
- Real-time GPIO state monitoring
- PostgreSQL database with PostGIS
- Metrics and monitoring
- File serving capabilities
- Secure authentication

## System Architecture

```
┌─────────────┐     ┌──────────┐
│   Client    │────▶│  Caddy   │
│ Application │◀────│  Proxy   │
└─────────────┘     └────┬─────┘
                        │
                ┌───────┴───────┐
                │               │
        ┌───────▼──┐     ┌─────▼─────┐
        │   GPIO   │     │  Metrics  │
        │ Service  │     │  Service  │
        └───────┬──┘     └─────┬─────┘
                │             │
                │       ┌─────▼─────┐
                │       │ PostgreSQL │
                │       │   PostGIS  │
                │       └───────────┘
                │
        ┌───────▼──────┐
        │ Physical GPIO │
        └──────────────┘
```

## Quick Start

1. Clone and setup:
```bash
git clone <repository-url>
cd docker-setup
cp .env.example .env
# Edit .env with your settings
```

2. Start services:
```bash
sudo docker compose up -d
```

3. Create user and get token:
```bash
# Create user
curl -X POST http://localhost/auth/register \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"your_password"}'

# Get token
curl -X POST http://localhost/auth/login \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"your_password"}'
```

## Documentation Index

1. [Setup Guide](SETUP.md)
   - System requirements
   - Installation
   - Configuration
   - Security considerations

2. [API Documentation](API.md)
   - REST endpoints
   - WebSocket interface
   - Authentication
   - Error handling

3. [Development Guide](DEVELOPMENT.md)
   - Development environment
   - Testing
   - Contributing
   - Deployment

## Resource Usage

Designed for Raspberry Pi Zero 2 W:
- Memory: ~400MB RAM
- Storage: ~1GB
- CPU: 4 cores utilized efficiently
- Network: Minimal bandwidth requirements

## Features

1. GPIO Control
   - Real-time WebSocket interface
   - REST API
   - Event-based state monitoring
   - Pin state persistence

2. Security
   - JWT authentication
   - HTTPS support (optional)
   - Permission-based access
   - Secure password storage

3. Monitoring
   - System metrics
   - Service health checks
   - Resource usage tracking
   - GPIO operation logging

4. Database
   - PostgreSQL with PostGIS
   - Spatial data support
   - Query optimization
   - Automatic backups

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues and feature requests, please use the GitHub issue tracker.

## Contributing

See [DEVELOPMENT.md](DEVELOPMENT.md) for contribution guidelines.
