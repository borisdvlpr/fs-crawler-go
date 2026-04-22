package main

import (
	"log/slog"
	"os"

	crawler "github.com/borisdvlpr/fs-crawler-go/cmd"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := crawler.Run(); err != nil {
		slog.Error("server exited with error", "error", err)
		os.Exit(1)
	}
}
