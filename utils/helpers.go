package utils

import "os"

func FolderExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	if fileInfo.IsDir() {
		return true
	}

	return false
}
