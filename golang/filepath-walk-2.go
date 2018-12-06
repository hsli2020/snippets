package main

import (
    "os"
    "path/filepath"
    "strings"
    "fmt"
    "io/ioutil"
     "log"
)

var files []os.FileInfo

func walker(path string, info os.FileInfo, err error) error {
    if strings.HasSuffix(info.Name(), ".php") {
        files = append(files, info)
    }
    return nil
}

func method1() {
	// walk recursively
    err := filepath.Walk(".", walker)
    if err != nil {
        println("Error", err)
    } else {
        for _, f := range files {
            println(f.Name())
            // This is where we'd like to open the file
        }
    }
}

func method2() {
    files, err := filepath.Glob("*")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(files) // contains a list of all files in the current directory
}

func method3() {
    files, err := ioutil.ReadDir("./snippets")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
		fmt.Println(f.Name())
    }
}

func main() {
	//method1()
	method2()
	//method3()
}
