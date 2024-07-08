package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bagashiz/sea-salon/internal/config"
)

// The Start function creates a new http.Server type, configures the routes, and adds middleware.
func Start(ctx context.Context, cfg *config.App, sessionManager *scs.SessionManager) {
	mux := http.NewServeMux()

	// global middleware
	var handler http.Handler = mux
	handler = sessionManager.LoadAndSave(handler)
	handler = logger(handler)

	// routes
	mux.Handle("GET /assets/", staticFiles())

	mux.Handle("GET /", notFound())
	mux.Handle("GET /{$}", index())

	mux.Handle("GET /register", register())
	mux.Handle("GET /login", login())

	server := &http.Server{
		Addr:    net.JoinHostPort(cfg.Host, cfg.Port),
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("error listening and serving", "error", err)
		}
	}()

	slog.Info("started the HTTP server", "host", cfg.Host, "port", cfg.Port)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			slog.Error("error shutting down server", "error", err)
			return
		}

		slog.Info("server shut down gracefully")
	}()

	wg.Wait()
}
