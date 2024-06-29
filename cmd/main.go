package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/bagashiz/sea-salon/internal/server"
	"github.com/bagashiz/sea-salon/pkg/config"
	"github.com/bagashiz/sea-salon/pkg/logger"
	db "github.com/bagashiz/sea-salon/pkg/postgres"
)

// entry point of the application.
func main() {
	ctx := context.Background()
	logger.Set()

	if err := run(ctx, os.Getenv); err != nil {
		slog.Error("error running application", "error", err)
		os.Exit(1)
	}
}

/**
 * The run function sets up the server and starts it.
 * It also listens for an interrupt signal to shut down the server gracefully.
 */
func run(ctx context.Context, getEnv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config, err := config.New(getEnv)
	if err != nil {
		return err
	}

	store, err := db.NewStore(ctx, config.DB)
	if err != nil {
		return err
	}

	slog.Info("connected to database", "type", config.DB.Type)

	if err := store.Migrate(); err != nil {
		return err
	}

	httpServer := server.NewServer(config.App)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error listening and serving", "error", err)
		}
	}()

	slog.Info("started the HTTP server", "host", config.App.Host, "port", config.App.Port)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("error shutting down server", "error", err)
			return
		}

		slog.Info("server shut down gracefully")
	}()

	wg.Wait()

	return nil
}
