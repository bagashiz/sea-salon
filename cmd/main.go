package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bagashiz/sea-salon/internal/app/user"
	"github.com/bagashiz/sea-salon/internal/config"
	"github.com/bagashiz/sea-salon/internal/postgres"
	"github.com/bagashiz/sea-salon/internal/postgres/repository"
	"github.com/bagashiz/sea-salon/internal/server"
	"github.com/bagashiz/sea-salon/internal/session"
)

// entry point of the application.
func main() {
	ctx := context.Background()

	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)
	slog.SetDefault(logger)

	if err := run(ctx, os.Getenv); err != nil {
		slog.Error("error running application", "error", err)
		os.Exit(1)
	}
}

// run sets up dependencies and starts the application.
func run(ctx context.Context, getEnv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	config, err := config.New(getEnv)
	if err != nil {
		return err
	}

	db, err := postgres.NewDB(ctx, config.DB)
	if err != nil {
		return err
	}
	defer db.Close()

	slog.Info("connected to the database", "type", config.DB.Type)

	if err := db.Migrate(config.DB.Type); err != nil {
		return err
	}

	sessionManager, err := session.New(config.App, db.Pool)
	if err != nil {
		return err
	}

	postgresRepo := repository.New(db)
	userService := user.NewService(postgresRepo)

	httpServer := server.New(config.App, sessionManager, userService)

	slog.Info("starting the http server", "addr", httpServer.Addr)

	if err := httpServer.Start(ctx); err != nil {
		return err
	}

	return nil
}
