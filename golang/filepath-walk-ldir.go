package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			/*
				type FileInfo interface {
					Name() string       // base name of the file
					Size() int64        // length in bytes for regular files; system-dependent for others
					Mode() FileMode     // file mode bits
					ModTime() time.Time // modification time
					IsDir() bool        // abbreviation for Mode().IsDir()
					Sys() interface{}   // underlying data source (can return nil)
				}
			*/
			if !info.IsDir() {
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
