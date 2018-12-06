package main

import (
    "fmt"
    "os"
    "log"
    "io/ioutil"
    "path/filepath"
)

func main() {
    var files []string

    root := "/some/folder/to/scan"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == ".dat" {
			return nil
		}

        files = append(files, path)
		//files = append(files, info.Name())

        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}

// method 2
func visit(files *[]string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Fatal(err)
        }
        *files = append(*files, path)
        return nil
    }
}

func method2() {
    var files []string

    root := "/some/folder/to/scan"
    err := filepath.Walk(root, visit(&files))
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}


// method 3
func method3() {
    files, err := ioutil.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
}

// method 4
func method4() {
    dirname := "."

    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
}
