package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
    "time"
)

func DeleteOldFile(fileExt string, cutoff time.Duration) func (string, os.FileInfo, error) error {
    now := time.Now()

    return func (path string, info os.FileInfo, err error) error {
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
            if filepath.Ext(path) != fileExt {
                return nil
            }
            if diff := now.Sub(info.ModTime()); diff > cutoff {
                fmt.Printf("Deleting %s\n", path)
                os.Remove(path)
            }
        }
        return nil
    }
}

func main() {
    var cutoff = 1 * time.Minute // time.Hour

	err := filepath.Walk(".", DeleteOldFile(".csv", cutoff))
	if err != nil {
		log.Println(err)
	}
}

/*
	rm -rf tmp
	mkdir tmp tmp/A tmp/B tmp/C tmp/D tmp/E tmp/F tmp/G

	for n in {1..100}; do touch "tmp/A/$n.csv"; done
	for n in {1..100}; do touch "tmp/B/$n.csv"; done
	for n in {1..100}; do touch "tmp/C/$n.csv"; done
	for n in {1..100}; do touch "tmp/D/$n.csv"; done
	for n in {1..100}; do touch "tmp/E/$n.csv"; done
	for n in {1..100}; do touch "tmp/F/$n.csv"; done
	for n in {1..100}; do touch "tmp/G/$n.csv"; done
*/