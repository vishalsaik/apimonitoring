package mongodb

import (
	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/config/logger"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDBConfig(cfg *config.Config, log *logger.Logger) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoDBConfig.MongoDBUrl)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(cfg.MongoDBConfig.MongoDBName)
	log.Info("Connected to MongoDB successfully", zap.String("database", cfg.MongoDBConfig.MongoDBName))
	return &MongoDB{
		Client:   client,
		Database: db,
	}, nil
}
