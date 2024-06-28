package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/bagashiz/sea-salon/internal/server"
	"github.com/joho/godotenv"
)

// init loads the environment variables from the .env file for local development.
func init() {
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}
}

// entry point of the application.
func main() {
	ctx := context.Background()

	if err := run(ctx, os.Getenv); err != nil {
		fmt.Printf("error: %v", err)
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

	config := map[string]string{
		"APP_HOST":          getEnv("APP_HOST"),
		"APP_PORT":          getEnv("APP_PORT"),
		"POSTGRES_USER":     getEnv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": getEnv("POSTGRES_PASSWORD"),
		"POSTGRES_DB":       getEnv("POSTGRES_DB"),
	}

	httpServer := server.NewServer(config)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("error listening and serving: %s\n", err)
		}
	}()

	fmt.Printf("listening on %s\n", httpServer.Addr)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("error shutting down server: %s\n", err)
			return
		}

		fmt.Println("server shut down gracefully")
	}()

	wg.Wait()

	return nil
}
