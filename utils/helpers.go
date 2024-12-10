package utils

import (
	"fmt"
	"os"
)

func ValidateFolder(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fileInfo != nil && !fileInfo.IsDir() {
		return fmt.Errorf("%s : is not a directory", path)
	}

	return nil
}
