package logger

import (
	"fmt"
	"log/slog"
	"os"
)

func InitLogger(p string, debugMode bool) error {
	if p == "" {
		return fmt.Errorf("log path cannot be empty")
	}

	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	logOptions := &slog.HandlerOptions{AddSource: true}

	if debugMode {
		logOptions.Level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(f, logOptions)
	slog.SetDefault(slog.New(handler))

	return nil
}
