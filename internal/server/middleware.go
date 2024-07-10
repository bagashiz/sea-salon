package server

import (
	"log/slog"
	"net/http"
	"time"
)

// handlerFunc is a function that handles an HTTP request and returns an error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// handlerError is an error that contains an HTTP status code and message.
type handlerError struct {
	statusCode int
	message    string
}

// Error returns the error message for the handlerError type.
func (h handlerError) Error() string {
	return h.message
}

// responseWriter is a wrapper around http.ResponseWriter that stores the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the WriteHeader method to store the status code.
func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// handle wraps a handlerFunc as an http.Handler, handles errors, and logs requests.
func handle(h handlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		if err := h(writer, r); err != nil {
			if err, ok := err.(handlerError); ok {
				http.Error(writer, err.message, err.statusCode)
				slog.Error("request failed",
					slog.Int("status", err.statusCode),
					slog.String("error", err.Error()),
					slog.String("method", r.Method),
					slog.String("url", r.URL.Path),
					slog.Duration("duration", time.Since(start)),
				)
				return
			}
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			slog.Error("request failed",
				slog.Int("status", http.StatusInternalServerError),
				slog.String("error", err.Error()),
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.Duration("duration", time.Since(start)),
			)
			return
		}

		slog.Info("request served",
			slog.Int("status", writer.statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", time.Since(start)),
		)
	})
}
