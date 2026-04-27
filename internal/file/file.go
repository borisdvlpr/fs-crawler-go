package file

import (
	"fmt"
	"io"
	"os"
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
