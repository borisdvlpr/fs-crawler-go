package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"sync"

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

	raw := strings.TrimSpace(string(fileData))
	paths := strings.Split(raw, ";")

	var validPaths []string
	for _, p := range paths {
		if p = strings.TrimSpace(p); p != "" {
			validPaths = append(validPaths, p)
		}
	}

	if len(validPaths) == 0 {
		return fmt.Errorf("paths file is empty")
	}

	slog.Info("starting crawl", "path_count", len(paths))

	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()
			defer func() {
				if r := recover(); r != nil {
					slog.Error("crawler panicked", "path", p, "err", r)
				}
			}()

			crawler.StartCrawler(p)
		}(path)
	}

	wg.Wait()

	slog.Info("crawl finished", "paths_processed", len(paths))

	return nil
}
