package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/borisdvlpr/fs-crawler-go/internal/file"
	"github.com/borisdvlpr/fs-crawler-go/internal/log"
)

func Run() error {
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		return err
	}

	if err := log.SetupLogs(); err != nil {
		return err
	}

	fileData, err := file.ReadFile("./folders.txt")
	if err != nil {
		return err
	}

	slog.Info("file successfully read")

	paths := strings.Split(string(fileData), ";")

	var validPaths []string
	for _, p := range paths {
		if p = strings.TrimSpace(p); p != "" {
			validPaths = append(validPaths, p)
		}
	}

	if len(validPaths) == 0 {
		return fmt.Errorf("paths file is empty")
	}

	return nil
}
