package app

import (
	"context"

	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/config/logger"
	"api-monitoring/src/shared/config/mongodb"
	"api-monitoring/src/shared/config/postgres"
	"api-monitoring/src/shared/config/rabbitmq"
	authRepository "api-monitoring/src/shared/services/auth/repository"
	"api-monitoring/src/shared/services/auth/controller"
	"api-monitoring/src/shared/services/auth/service"
)

type App struct {
	Config         *config.Config
	Log            *logger.Logger
	Mongo          *mongodb.MongoDB
	Postgres       *postgres.Postgres
	Rabbit         *rabbitmq.RabbitMQ
	AuthController *controller.AuthController
}

func Initialize() (*App, error) {
	cfg := config.NewConfig()

	log, err := logger.NewLogger(cfg)
	if err != nil {
		return nil, err
	}

	mongo, err := mongodb.NewMongoDBConfig(cfg, log)
	if err != nil {
		log.Error("Failed to connect to MongoDB")
		return nil, err
	}

	pg, err := postgres.NewPostgres(cfg, log)
	if err != nil {
		log.Error("Failed to connect to Postgres")
		return nil, err
	}

	rabbit, err := rabbitmq.NewRabbitMQ(cfg, log)
	if err != nil {
		log.Error("Failed to connect to RabbitMQ")
		return nil, err
	}
	collection := mongo.Client.Database(cfg.MongoDBConfig.MongoDBName).Collection("users")
	userRepo := authRepository.NewMongoUserRepository(collection, log)
	authSvc := service.NewAuthService(userRepo, cfg.JwtConfig.SecretKey, cfg.JwtConfig.ExpirationTime, log)
	authCtrl := controller.NewAuthController(authSvc, cfg.CookieConfig)

	return &App{
		Config:         cfg,
		Log:            log,
		Mongo:          mongo,
		Postgres:       pg,
		Rabbit:         rabbit,
		AuthController: authCtrl,
	}, nil
}

func (a *App) Shutdown() {
	if err := a.Mongo.Client.Disconnect(context.Background()); err != nil {
		a.Log.Error("Failed to disconnect MongoDB")
	} else {
		a.Log.Info("Disconnected MongoDB successfully")
	}

	a.Postgres.Pool.Close()
	a.Log.Info("Disconnected Postgres successfully")

	a.Rabbit.Channel.Close()
	a.Rabbit.Connection.Close()
	a.Log.Info("Disconnected RabbitMQ successfully")
}
