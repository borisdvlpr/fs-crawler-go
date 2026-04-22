package cmd

import (
	"github.com/borisdvlpr/fs-crawler-go/internal/file"
)

func Run() error {
	if err := file.CreateDir("output"); err != nil {
		return err
	}

	if err := file.CreateDir("logs"); err != nil {
		return err
	}

	return nil
}
