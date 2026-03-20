package main

import (
	"context"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"api-monitoring/app"
	"api-monitoring/handlers"
)

func main() {
	a, err := app.Initialize()
	if err != nil {
		stdlog.Fatalf("Failed to initialize app: %v", err)
	}
	defer a.Shutdown()
	defer a.Log.Sync()

	router := handlers.NewRouter(a.Log)

	srv := &http.Server{
		Addr:    ":" + a.Config.Server.Port,
		Handler: router,
	}

	go func() {
		a.Log.Info("Server started", zap.String("port", a.Config.Server.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Log.Error("Server error", zap.Error(err))
		}
	}()

	// ── Graceful Shutdown ─────────────────────────────────
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.Log.Info("Shutdown signal received, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.Log.Error("Forced shutdown", zap.Error(err))
	} else {
		a.Log.Info("HTTP server stopped")
	}
}
