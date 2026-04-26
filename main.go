package main

import (
	"log/slog"
	"os"

	runner "github.com/borisdvlpr/fs-crawler-go/cmd"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := runner.Run(); err != nil {
		slog.Error("application exited with error", "error", err)
		os.Exit(1)
	}
}
