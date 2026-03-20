package config

import (
	"os"
	"strconv"
)

type Config struct {
	Server          serverConfig
	MongoDBConfig   MongoDBConfig
	RabbitMQConfig  rabbitMQConfig
	PostgresConfig  PostgresConfig
	JwtConfig       jwtConfig
	CookieConfig    CookieConfig
	RatelimitConfig ratelimitConfig
	CorsConfig      corsConfig
}

type corsConfig struct {
	AllowedOrigin string
}

type CookieConfig struct {
	MaxAge   int
	HttpOnly bool
	Secure   bool
}
type ratelimitConfig struct {
	WindowMilliseconds int
	MaxRequests        int
}
type jwtConfig struct {
	SecretKey      string
	ExpirationTime int
}
type serverConfig struct {
	Environment string
	Port        string
}
type rabbitMQConfig struct {
	RabbitMQUrl                 string
	RabbitMQQueueName           string
	RabbitMqpublishConfirmation bool
	RabbitMqretryCount          int
	RabbitMqretryDelay          int
}
type PostgresConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDBName   string
	PostgresUser     string
	PostgresPassword string
}
type MongoDBConfig struct {
	MongoDBUrl  string
	MongoDBName string
}

func getEnvInt(key string, fallback int) int {
	if val, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return val
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if val, err := strconv.ParseBool(os.Getenv(key)); err == nil {
		return val
	}
	return fallback
}

func NewConfig() *Config {
	return &Config{
		Server: serverConfig{
			Environment: os.Getenv("APP_ENV"),
			Port:        os.Getenv("APP_PORT"),
		},
		PostgresConfig: PostgresConfig{
			PostgresHost:     os.Getenv("POSTGRES_HOST"),
			PostgresPort:     os.Getenv("POSTGRES_PORT"),
			PostgresUser:     os.Getenv("POSTGRES_USER"),
			PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
			PostgresDBName:   os.Getenv("POSTGRES_DB"),
		},
		MongoDBConfig: MongoDBConfig{
			MongoDBUrl:  os.Getenv("MONGO_URL"),
			MongoDBName: os.Getenv("MONGO_DB"),
		},
		RabbitMQConfig: rabbitMQConfig{
			RabbitMQUrl:                 os.Getenv("RABBITMQ_URL"),
			RabbitMQQueueName:           os.Getenv("RABBITMQ_QUEUE_NAME"),
			RabbitMqpublishConfirmation: getEnvBool("RABBITMQ_PUBLISH_CONFIRMATION", true),
			RabbitMqretryCount:          getEnvInt("RABBITMQ_RETRY_COUNT", 3),
			RabbitMqretryDelay:          getEnvInt("RABBITMQ_RETRY_DELAY", 500),
		},
		JwtConfig: jwtConfig{
			SecretKey:      os.Getenv("JWT_SECRET_KEY"),
			ExpirationTime: getEnvInt("JWT_EXPIRATION_TIME", 3600),
		},
		CookieConfig: CookieConfig{
			MaxAge:   getEnvInt("COOKIE_MAX_AGE", 3600),
			HttpOnly: getEnvBool("COOKIE_HTTP_ONLY", true),
			Secure:   getEnvBool("COOKIE_SECURE", false),
		},
		RatelimitConfig: ratelimitConfig{
			WindowMilliseconds: getEnvInt("RATE_LIMIT_WINDOW_MS", 60000),
			MaxRequests:        getEnvInt("RATE_LIMIT_MAX_REQUESTS", 100),
		},
		CorsConfig: corsConfig{
			AllowedOrigin: os.Getenv("ALLOWED_ORIGIN"),
		},
	}
}
