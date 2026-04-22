package file

import (
	"os"
)

func CreateDir(dirName string) error {
	if err := os.MkdirAll(dirName, os.ModePerm); err != nil {
		return err
	}

	return nil
}
