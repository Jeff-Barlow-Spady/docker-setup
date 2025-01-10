# Development Guide

## Development Environment Setup

### Prerequisites

- Python 3.9+
- Docker and Docker Compose
- Git
- Code editor (VS Code recommended)
- `curl` for API testing

### Local Development Setup

1. Create Python virtual environment:
```bash
python -m venv venv
source venv/bin/activate
```

2. Install development dependencies:
```bash
pip install -r services/gpio/requirements.txt
pip install pytest pytest-asyncio pytest-cov black flake8
```

## Project Structure

```
.
├── services/
│   ├── gpio/              # GPIO service
│   │   ├── main.py
│   │   └── requirements.txt
│   ├── auth/              # Authentication service
│   │   └── auth.py
│   └── metrics/           # Metrics service
│       └── metrics.py
├── docs/                  # Documentation
├── data/                  # Data directory
├── caddy/                 # Caddy configuration
└── docker-compose.yml     # Container orchestration
```

## API Development

### WebSocket API

Example WebSocket client:
```python
import websockets
import asyncio
import json

async def gpio_client():
    uri = "ws://localhost/ws/gpio"
    async with websockets.connect(uri) as websocket:
        # Write to GPIO
        await websocket.send(json.dumps({
            "action": "write",
            "pin": 18,
            "value": True
        }))
        response = await websocket.recv()
        print(response)

        # Listen for changes
        while True:
            message = await websocket.recv()
            print(f"Received: {message}")

asyncio.get_event_loop().run_until_complete(gpio_client())
```

### REST API

Example REST client:
```python
import requests

# Login
response = requests.post(
    "http://localhost/auth/login",
    json={"username": "admin", "password": "password"}
)
token = response.json()["token"]

# Setup GPIO
requests.post(
    "http://localhost/gpio/18/setup",
    headers={"Authorization": f"Bearer {token}"},
    json={"direction": "out"}
)

# Write to GPIO
requests.post(
    "http://localhost/gpio/18/write",
    headers={"Authorization": f"Bearer {token}"},
    json={"value": True}
)
```

## Testing

### Unit Tests

Run unit tests:
```bash
pytest tests/
```

With coverage:
```bash
pytest --cov=services tests/
```

### Integration Tests

```bash
pytest tests/integration/
```

## Code Style

This project follows PEP 8 style guide. Use Black for formatting:

```bash
black services/
```

Check style with flake8:
```bash
flake8 services/
```

## Debugging

### GPIO Service

Enable debug logging:
```python
import logging
logging.basicConfig(level=logging.DEBUG)
```

### WebSocket Debugging

Use websocat for terminal-based testing:
```bash
websocat ws://localhost/ws/gpio
```

## Metrics and Monitoring

Access Prometheus metrics:
```bash
curl localhost/metrics
```

Monitor specific components:
```bash
curl localhost/health
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make changes
4. Run tests
5. Submit pull request

### Commit Message Format

```
type(scope): description

[optional body]

[optional footer]
```

Types:
- feat: New feature
- fix: Bug fix
- docs: Documentation
- style: Formatting
- refactor: Code restructuring
- test: Adding tests
- chore: Maintenance

## Release Process

1. Update version in setup.py
2. Update CHANGELOG.md
3. Tag release
4. Build and push Docker images
5. Update documentation

## Deployment

### Development
```bash
docker compose -f docker-compose.dev.yml up
```

### Production
```bash
docker compose up -d
```

## Common Issues and Solutions

1. GPIO Permission Issues
    ```bash
    sudo usermod -a -G gpio $USER
    ```

2. WebSocket Connection Failed
    - Check Caddy logs
    - Verify port forwarding
    - Check authentication token

3. Database Connection Issues
    - Check PostgreSQL logs
    - Verify connection string
    - Check network connectivity
