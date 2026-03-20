# Local dev config — run: source config.sh
# Values match docker-compose.yaml infra services
# Hosts are localhost because Docker exposes ports to your machine

# Server
export APP_ENV=development
export APP_PORT=5000

# PostgreSQL
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5433
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=root
export POSTGRES_DB=apimonitoring

# MongoDB
export MONGO_URL=mongodb://localhost:27017
export MONGO_DB=apimonitoring

# RabbitMQ
export RABBITMQ_URL=amqp://api_user:api_password@localhost:5672/api_vhost
export RABBITMQ_QUEUE_NAME=api_hits
export RABBITMQ_PUBLISH_CONFIRMATION=true
export RABBITMQ_RETRY_COUNT=3
export RABBITMQ_RETRY_DELAY=500

# JWT
export JWT_SECRET_KEY=your-local-dev-secret-key
export JWT_EXPIRATION_TIME=3600

# Cookie
export COOKIE_MAX_AGE=3600
export COOKIE_HTTP_ONLY=true
export COOKIE_SECURE=false

# Rate limiting
export RATE_LIMIT_WINDOW_MS=60000
export RATE_LIMIT_MAX_REQUESTS=100

# CORS
export ALLOWED_ORIGIN=http://localhost:3000

#password validation
export PASSWORD_MIN_LENGTH=8
export PASSWORD_REQUIRE_UPPERCASE=true
export PASSWORD_REQUIRE_LOWERCASE=true
export PASSWORD_REQUIRE_NUMBER=true     
export PASSWORD_REQUIRE_SPECIAL=true