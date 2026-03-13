package rabbitmq

import (
	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/logger"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQ(cfg *config.Config, log *logger.Logger) (*RabbitMQ, error) {
	conn, err := amqp.Dial(cfg.RabbitMQConfig.RabbitMQUrl)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ", zap.Error(err))
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Failed to open a channel", zap.Error(err))
		return nil, err
	}
	log.Info("Connected to RabbitMQ successfully", zap.String("url", cfg.RabbitMQConfig.RabbitMQUrl))
	return &RabbitMQ{Connection: conn, Channel: ch}, nil
}
