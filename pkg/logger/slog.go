package logger

import (
	"log/slog"
	"os"
)

// Set sets the default slog format.
func Set() {
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)
	slog.SetDefault(logger)
}
