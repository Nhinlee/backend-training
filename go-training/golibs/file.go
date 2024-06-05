package golibs

import (
	"io/fs"
	"os"
	"path/filepath"
)

func CountFilesInFolder(folderPath string) (int, error) {
	var fileCount int

	err := filepath.Walk(folderPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileCount++
		}

		return nil
	})

	return fileCount, err
}

func ReadFile(folderPath string) (string, error) {
	content, err := os.ReadFile(folderPath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
