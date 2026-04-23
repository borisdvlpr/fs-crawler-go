package log

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	slogmulti "github.com/samber/slog-multi"
)

func SetupLogs() error {
	timestamp := time.Now().Unix()
	logFilePath := fmt.Sprintf("./logs/log_%d.log", timestamp)

	logFile, err := os.Create(logFilePath)
	if err != nil {
		return fmt.Errorf("failed to create log file: %w", err)
	}

	opts := &slog.HandlerOptions{Level: slog.LevelInfo}

	slog.SetDefault(slog.New(
		slogmulti.Fanout(
			slog.NewTextHandler(os.Stdout, opts), // human-readable to terminal
			slog.NewJSONHandler(logFile, opts),   // JSON to file
		),
	))

	return nil
}
