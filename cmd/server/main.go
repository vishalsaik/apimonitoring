package main

import (
	"context"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"

	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/logger"
	"api-monitoring/src/shared/mongodb"
	"api-monitoring/src/shared/postgres"
	"api-monitoring/src/shared/rabbitmq"

	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()

	log, err := logger.NewLogger(cfg)
	if err != nil {
		stdlog.Fatalf("Failed to create logger: %v", err)
	}
	defer log.Sync()

	mongo, err := mongodb.NewMongoDBConfig(cfg, log)
	if err != nil {
		log.Error("Failed to connect to MongoDB", zap.Error(err))
		return
	}
	defer func() {
		if err := mongo.Client.Disconnect(context.Background()); err != nil {
			log.Error("Failed to disconnect MongoDB", zap.Error(err))
		} else {
			log.Info("Disconnected MongoDB successfully")
		}
	}()

	pg, err := postgres.NewPostgres(cfg, log)
	if err != nil {
		log.Error("Failed to connect to Postgres", zap.Error(err))
		return
	}
	defer func() {
		pg.Pool.Close()
		log.Info("Disconnected Postgres pool successfully")
	}()

	rabbit, err := rabbitmq.NewRabbitMQ(cfg, log)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ", zap.Error(err))
		return
	}
	defer func() {
		rabbit.Channel.Close()
		rabbit.Connection.Close()
		log.Info("Disconnected RabbitMQ successfully")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Server started — press Ctrl+C to stop")
	<-quit
	log.Info("Shutdown signal received, disconnecting...")
}
