package server

import (
	"log/slog"
	"net/http"
	"time"
)

// handlerFunc is a function that handles an HTTP request and returns an error.
type handlerFunc func(http.ResponseWriter, *http.Request) error

// handlerError is a custom error for HTTP handlers that includes a status code.
type handlerError struct {
	message    string
	statusCode int
}

// Error returns the error message for the handlerError type.
func (h handlerError) Error() string {
	return h.message
}

// responseWriter extends the http.ResponseWriter type to store the status code.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the WriteHeader method to store the status code.
func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// handle wraps a handlerFunc type as an http.Handler, switch to custom ResponseWriter,
// handles errors, and logs request information.
func handle(h handlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &responseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		if err := h(writer, r); err != nil {
			statusCode := http.StatusInternalServerError
			errMsg := http.StatusText(statusCode)

			if err, ok := err.(handlerError); ok {
				statusCode = err.statusCode
				errMsg = err.Error()
			}

			http.Error(writer, errMsg, statusCode)
			logRequest(r, statusCode, time.Since(start), errMsg)
			return
		}

		logRequest(r, writer.statusCode, time.Since(start), "")
	})
}

// logRequest logs request information based on the status code and error message.
func logRequest(r *http.Request, statusCode int, duration time.Duration, errMsg string) {
	switch {
	case statusCode < 400:
		slog.Info("success",
			slog.Int("status", statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", duration),
		)
	case statusCode >= 400 && statusCode < 500:
		slog.Warn("client error",
			slog.Int("status", statusCode),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", duration),
		)
	case statusCode >= 500 || errMsg != "":
		slog.Error("server error",
			slog.Int("status", statusCode),
			slog.String("error", errMsg),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
			slog.Duration("duration", duration),
		)
	}
}
