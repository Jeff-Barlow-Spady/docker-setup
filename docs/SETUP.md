# Edge Device Setup Guide

## System Requirements

- Raspberry Pi Zero 2 W or compatible
- Debian-based OS
- At least 512MB RAM
- 8GB+ SD card
- GPIO pins accessible

## Initial Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd docker-setup
```

2. Create necessary directories:
```bash
mkdir -p data/files services/{gpio,auth,metrics} docs
```

3. Set up environment variables:
```bash
cat << 'EOL' > .env
POSTGRES_DB=gisdb
POSTGRES_USER=gisuser
POSTGRES_PASSWORD=your_secure_password
JWT_SECRET_KEY=your_jwt_secret
TOKEN_EXPIRY_HOURS=24
EOL
```

4. Build and start services:
```bash
sudo docker compose up -d
```

## Service Configuration

### GPIO Service

1. Install Python dependencies:
```bash
cd services/gpio
pip install -r requirements.txt
```

2. Start the GPIO service:
```bash
python main.py
```

### Authentication

1. Create first user:
```bash
curl -X POST http://localhost/auth/register \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"your_password"}'
```

2. Get authentication token:
```bash
curl -X POST http://localhost/auth/login \
    -H "Content-Type: application/json" \
    -d '{"username":"admin","password":"your_password"}'
```

## Monitoring Setup

1. Access metrics:
```bash
curl http://localhost/metrics
```

2. Monitor system health:
```bash
curl http://localhost/health
```

## Security Considerations

1. Change default passwords
2. Update JWT secret key
3. Enable HTTPS in production
4. Restrict access to sensitive endpoints
5. Regular security updates

## Backup and Maintenance

1. Database backup:
```bash
./manage.sh backup
```

2. View logs:
```bash
./manage.sh logs
```

3. Update services:
```bash
git pull
sudo docker compose pull
sudo docker compose up -d
```

## Troubleshooting

1. Check service status:
```bash
sudo docker compose ps
```

2. View service logs:
```bash
sudo docker compose logs -f
```

3. Restart services:
```bash
sudo docker compose restart
```

4. Check GPIO permissions:
```bash
ls -l /dev/gpio*
```

## Development

For development setup and contribution guidelines, see [DEVELOPMENT.md](DEVELOPMENT.md)
