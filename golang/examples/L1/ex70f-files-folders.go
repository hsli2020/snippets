package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func dir(thepath string) []string {

  var files []string

  filepath.Walk(thepath, func(path string, _ os.FileInfo, _ error) error {
    //fmt.Println(path)
    files = append(files, path)
    return nil
  })

  return files
}

func cwd() {
    pwd, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(pwd)
}

func main() {
  path := "/"
  fmt.Println(dir(path))
}
