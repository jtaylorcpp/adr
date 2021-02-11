package utils

import (
	"os"
)

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
