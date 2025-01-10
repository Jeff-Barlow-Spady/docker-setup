# Edge Device API Documentation

## GPIO Service API

### REST Endpoints

#### Setup GPIO Pin
```http
POST /gpio/{pin}/setup
Authorization: Bearer <token>

{
    "direction": "out"|"in",
    "initial": true|false  // Optional, for output pins
}
```

#### Write to GPIO Pin
```http
POST /gpio/{pin}/write
Authorization: Bearer <token>

{
    "value": true|false
}
```

#### Read from GPIO Pin
```http
GET /gpio/{pin}/read
Authorization: Bearer <token>
```

#### Metrics
```http
GET /metrics
```

### WebSocket Interface

Connect to WebSocket:
```
ws://<host>/ws/gpio
```

#### WebSocket Messages

Writing to a pin:
```json
{
    "action": "write",
    "pin": <pin_number>,
    "value": true|false
}
```

Reading from a pin:
```json
{
    "action": "read",
    "pin": <pin_number>
}
```

#### WebSocket Events

Pin state change event:
```json
{
    "event": "pin_change",
    "pin": <pin_number>,
    "value": true|false
}
```

### Metrics Available

1. `gpio_operations_total`: Counter of GPIO operations
    - Labels: operation, pin

2. `gpio_pin_state`: Current state of GPIO pins
    - Labels: pin

3. `active_websocket_connections`: Number of active WebSocket connections

## Authentication

### Create User
```http
POST /auth/register
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```

### Login
```http
POST /auth/login
Content-Type: application/json

{
    "username": "string",
    "password": "string"
}
```

Response:
```json
{
    "token": "JWT_TOKEN",
    "token_type": "bearer"
}
```

## Error Handling

All API endpoints return standard HTTP status codes:

- 200: Success
- 400: Bad Request
- 401: Unauthorized
- 403: Forbidden
- 404: Not Found
- 500: Internal Server Error

Error response format:
```json
{
    "status": "error",
    "message": "Error description"
}
```

## Metrics Service API

### Endpoints

#### System Metrics
```http
GET /metrics
```

Response includes:
- CPU usage
- Memory usage
- Disk usage
- Service uptime
- GPIO operations
- Database metrics

#### Health Check
```http
GET /health
```

Response:
```json
{
    "status": "healthy|degraded",
    "timestamp": 1234567890,
    "checks": {
        "memory": "ok|warning",
        "cpu": "ok|warning",
        "disk": "ok|warning"
    }
}
```

### Prometheus Metrics

Available metrics:
- `system_cpu_usage`
- `system_memory_usage`
- `system_disk_usage`
- `system_uptime_seconds`
- `service_uptime_seconds`
- `http_requests_total`
- `database_connections_active`
- `database_queries_total`
