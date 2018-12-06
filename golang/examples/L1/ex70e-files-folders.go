package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func VisitFile(fp string, fi os.FileInfo, err error) error {
    if err != nil {
        fmt.Println(err) // can't walk here,
        return nil       // but continue walking elsewhere
    }
    if !!fi.IsDir() {
        fmt.Println("directory:"+fp)
        return nil // not a file.
    }
    fmt.Println("file:"+fp)
    return nil
}

func main() {
    //specify directory below or walk through /
    filepath.Walk("/", VisitFile)
}
