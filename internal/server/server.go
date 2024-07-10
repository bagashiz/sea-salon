package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bagashiz/sea-salon/internal/config"
	"golang.org/x/sync/errgroup"
)

// Server wraps the http.Server type for extending functionality.
type Server struct {
	*http.Server
}

// New creates a new http.Server type, configures the routes, and adds middleware.
func New(cfg *config.App, sessionManager *scs.SessionManager) *Server {
	mux := http.NewServeMux()

	addRoutes(mux)

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	handler := sessionManager.LoadAndSave(mux)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &Server{server}
}

// Start starts the HTTP server in a separate goroutine and listens for
// the context cancellation signal to shut down the server gracefully.
func (s *Server) Start(ctx context.Context) error {
	errs, ctx := errgroup.WithContext(ctx)

	errs.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	errs.Go(func() error {
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			return err
		}

		return nil
	})

	return errs.Wait()
}
