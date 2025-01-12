#!/bin/bash

# Configuration
PI_HOST=${PI_HOST:-"pi@raspberrypi.local"}  # Change this to your Pi's address
PI_PATH=${PI_PATH:-"/home/pi/edge-device"}   # Where to deploy on the Pi
VERSION=$(git describe --tags --always --dirty)

# Build Docker images with buildx for ARM
echo "Building Docker images for ARM..."
docker buildx create --use --name=crossplatform || true
docker buildx inspect --bootstrap

# Build and push images
echo "Building and pushing images..."
docker buildx build --platform linux/arm/v7 \
    -t localhost:5000/auth:${VERSION} \
    -f services/auth/Dockerfile ./services/auth --load

docker buildx build --platform linux/arm/v7 \
    -t localhost:5000/gpio:${VERSION} \
    -f services/gpio/Dockerfile ./services/gpio --load

docker buildx build --platform linux/arm/v7 \
    -t localhost:5000/metrics:${VERSION} \
    -f services/metrics/Dockerfile ./services/metrics --load

# Create deployment directory structure
echo "Creating deployment package..."
mkdir -p deploy
cp docker-compose.yml deploy/
cp .env.example deploy/.env
cp -r caddy deploy/
cp -r data deploy/
cp -r init-scripts deploy/

# Update docker-compose.yml with version
sed -i "s/latest/${VERSION}/g" deploy/docker-compose.yml

# Create deploy script for Pi
cat > deploy/start.sh << 'EOF'
#!/bin/bash
# Load images
for img in *.tar; do
    docker load -i "$img"
done

# Start services
docker compose up -d
EOF
chmod +x deploy/start.sh

# Save images to tar files
echo "Saving images..."
docker save localhost:5000/auth:${VERSION} > deploy/auth.tar
docker save localhost:5000/gpio:${VERSION} > deploy/gpio.tar
docker save localhost:5000/metrics:${VERSION} > deploy/metrics.tar

# Deploy to Pi
echo "Deploying to Raspberry Pi..."
ssh ${PI_HOST} "mkdir -p ${PI_PATH}"
rsync -avz --progress deploy/ ${PI_HOST}:${PI_PATH}/

# Start services on Pi
echo "Starting services on Pi..."
ssh ${PI_HOST} "cd ${PI_PATH} && ./start.sh"

echo "Deployment complete!"