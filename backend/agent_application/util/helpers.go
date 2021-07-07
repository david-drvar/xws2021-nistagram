package util

import (
	"os"
	"path/filepath"
)

const (
	filesystemLocation = "content"
)

func GetContentLocation(filename string) string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, filesystemLocation, filename)
}
