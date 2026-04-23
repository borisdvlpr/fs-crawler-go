package file

import (
	"fmt"
	"io"
	"os"
)

func CreateDir(dirName string) error {
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return err
	}

	return nil
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
