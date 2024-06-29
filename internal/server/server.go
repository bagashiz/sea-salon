package server

import (
	"net"
	"net/http"

	"github.com/bagashiz/sea-salon/internal/config"
)

// The NewServer function creates a new http.Server type, configures the routes, and adds middleware.
func NewServer(cfg *config.App) *http.Server {
	mux := http.NewServeMux()
	addRoutes(mux)

	var handler http.Handler = mux
	handler = logger(handler)

	addr := net.JoinHostPort(cfg.Host, cfg.Port)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return server
}

// The addRoutes function loads the routes with their respective handlers.
func addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", http.NotFoundHandler())
	mux.Handle("GET /{$}", index())
	mux.Handle("GET /assets/", staticFiles())
}
