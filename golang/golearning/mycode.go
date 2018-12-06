package main

import (
    "fmt"
    "log"
    "path/filepath"
    "os"
    "time"
    "strings"
    "reflect"
    "errors"
)

func pr(v ...interface{}) {
    //fmt.Printf("%q\n", v...)
    //fmt.Printf("%+v\n", v...)
    fmt.Printf("%#v\n", v...)
}

// archiveFiles("e:/amazon/*.xml")
func archiveFiles(pattern string) {
    files, err := filepath.Glob(pattern)
    checkError(err)
    //fmt.Println(files)

    for _, fname := range files {
        dir, file := filepath.Split(fname)
        fileTime, _ := getFileTime(fname)

        newDir := dir + "archive\\" + fileTime.Format("2006-01-02");
        os.MkdirAll(newDir, 0777)

        newFile := newDir + "\\" + file

        if time.Now().Sub(fileTime) > 15*24*time.Hour {
            //os.Rename(fname, newFile)
            fmt.Println(fname, newFile)
        }
    }
}

func checkError(err error) {
    if err != nil {
        log.Fatal(err)
        //panic(err)
    }
}

func checkErr(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func getFileTime(name string) (mtime time.Time, err error) {
    fi, err := os.Stat(name)
    if err != nil {
        return
    }
    mtime = fi.ModTime()
    return
}

// fmt.Println(getPartNum("DH-1234-ABC"))
func getPartNum(sku string) string {
    const sep = "-"

    arr := strings.Split(sku, sep);
    if (len(arr) > 1) {
        arr = arr[1:]
    }
    return strings.Join(arr, sep)
}
