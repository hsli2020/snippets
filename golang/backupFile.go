package main

import (
	"os"
	"path/filepath"
	"strings"
)

func BackupFile(filename, dir string) (string, error) {
	file, err := os.Stat(filename)
	if err != nil {
		return "", err
	}
	modTime := file.ModTime().Format("20060102-150405")

	ext := filepath.Ext(filename)
	name := strings.TrimSuffix(filename, ext)

	newName := name + "-" + modTime + ext
	if dir != "" {
		newName = filepath.Join(dir, newName)
	}

	err = os.Rename(filename, newName)
	if err != nil {
		return "", err
	}

	return newName, nil
}

func ArchiveFile(filename string) (string, error) {
	return BackupFile(filename, "archive")
}

func main() {
	BackupFile("zzz-pna.log", "files")
}
