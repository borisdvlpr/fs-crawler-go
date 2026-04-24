package cmd

import (
	"os"
)

func Run() error {
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		return err
	}

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		return err
	}

	return nil
}
