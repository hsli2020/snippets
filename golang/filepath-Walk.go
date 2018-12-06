package main

import (
    "fmt"
    //"html/template"
    "os"
    "path/filepath"
)

func consumer(p string, i os.FileInfo, e error) error {
    //t := template.New(p)
    //fmt.Println(t.Name())
    fmt.Println(p)
    return nil
}

func main() {
    filepath.Walk("./gotmpl", filepath.WalkFunc(consumer))
}
