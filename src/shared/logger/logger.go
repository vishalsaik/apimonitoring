package logger

import (
	"api-monitoring/src/shared/config"

	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger(cfg *config.Config) (*Logger, error) {
	if cfg.Server.Environment == "production" {
		prodLogger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}
		prodLogger = prodLogger.With(zap.String("service", "api-monitoring"))
		return &Logger{prodLogger}, nil

	} else {
		devLogger, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
		devLogger = devLogger.With(zap.String("service", "api-monitoring"))
		return &Logger{devLogger}, nil
	}

}
