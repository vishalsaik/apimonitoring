package main

import (
	"context"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"api-monitoring/handlers"
	"api-monitoring/src/shared/config"
	"api-monitoring/src/shared/config/logger"
	"api-monitoring/src/shared/config/mongodb"
	"api-monitoring/src/shared/config/postgres"
	"api-monitoring/src/shared/config/rabbitmq"
	"api-monitoring/src/shared/middleware"
	"api-monitoring/src/shared/utils"
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

	// ── Router ────────────────────────────────────────────
	router := gin.New()

	// Request logger middleware
	router.Use(middleware.LoggerHandler(log))

	// Error handler middleware
	router.Use(middleware.ErrorHandler(log))

	// ── Routes ────────────────────────────────────────────
	router.GET("/health", handlers.HealthCheckHandler)
	router.GET("/", handlers.RootHandler)

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.ErrorResponse("Endpoint not found", http.StatusNotFound, nil))
	})

	// ── HTTP Server ───────────────────────────────────────
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// Start server in background goroutine
	go func() {
		log.Info("Server started", zap.String("port", cfg.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("Server error", zap.Error(err))
		}
	}()

	// ── Graceful Shutdown ─────────────────────────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutdown signal received, shutting down gracefully...")

	// Give in-flight requests 10s to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Forced shutdown", zap.Error(err))
	} else {
		log.Info("HTTP server stopped")
	}
}
