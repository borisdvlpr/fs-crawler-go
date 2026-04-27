package file

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type FileEntry struct {
	Path string `json:"path"`
	Size int64  `json:"size"`
	LMod int64  `json:"lmod"`
}

type FilesList struct {
	Files []FileEntry `json:"files"`
}

func ReadFile(filePath string) ([]byte, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return data, nil
}

func StartCrawler(folderPath string) {
	slog.Info("Crawling", "path", folderPath)

	fileList := &FilesList{Files: []FileEntry{}}

	if err := readFolder(folderPath, fileList); err != nil {
		slog.Error("Error reading folder", "path", folderPath, "err", err)
		return
	}

	slog.Info("Folder read", "path", folderPath, "files", len(fileList.Files))

	if err := saveOutput(fileList); err != nil {
		slog.Error("Error saving output", "path", folderPath, "err", err)
	}
}

func readFolder(folderPath string, fileList *FilesList) error {
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("failed to read dir %s: %w", folderPath, err)
	}

	for _, entry := range entries {
		fullPath := filepath.Join(folderPath, entry.Name())

		if entry.IsDir() {
			if err := readFolder(fullPath, fileList); err != nil {
				slog.Error("error reading subfolder", "path", fullPath, "err", err)
			}
			continue
		}

		info, err := entry.Info()
		if err != nil {
			slog.Error("error reading file info", "path", fullPath, "err", err)
			continue
		}

		fileList.Files = append(fileList.Files, FileEntry{
			Path: fullPath,
			Size: info.Size(),
			LMod: info.ModTime().Unix(),
		})
	}

	return nil
}

func saveOutput(fileList *FilesList) error {
	jsonData, err := json.MarshalIndent(fileList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal output: %w", err)
	}

	timestamp := time.Now().UnixMicro()
	outputPath := fmt.Sprintf("./output/%d_output.json", timestamp)
	slog.Info("Saving file", "path", outputPath)

	if err := os.WriteFile(outputPath, jsonData, 0o444); err != nil {
		return fmt.Errorf("failed to write output file %s: %w", outputPath, err)
	}

	return nil
}
