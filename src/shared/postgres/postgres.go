package postgres

import (
	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/logger"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgres(cfg *config.Config, log *logger.Logger) (*Postgres, error) {
	// connString := "postgres://postgres:root@localhost:5432/apimonitoring"
	connString := "postgres://" + cfg.PostgresConfig.PostgresUser + ":" + cfg.PostgresConfig.PostgresPassword + "@" + cfg.PostgresConfig.PostgresHost + ":" + cfg.PostgresConfig.PostgresPort + "/" + cfg.PostgresConfig.PostgresDBName + "?sslmode=disable"
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Error("Failed to parse Postgres connection string", zap.Error(err))
		return nil, err
	}
	poolConfig.MaxConns = 10
	pgxPool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Error("Failed to connect to Postgres", zap.Error(err))
		return nil, err
	}
	err = pgxPool.Ping(context.Background())
	if err != nil {
		log.Error("Failed to ping Postgres", zap.Error(err))
		return nil, err
	}
	log.Info("Connected to Postgres successfully")
	return &Postgres{Pool: pgxPool}, nil
}
