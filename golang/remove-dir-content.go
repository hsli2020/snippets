// https://stackoverflow.com/questions/33450980/how-to-remove-all-contents-of-a-directory-using-golang
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func RemoveContents(dir string) error {
    d, err := os.Open(dir)
    if err != nil {
        return err
    }
    defer d.Close()
    names, err := d.Readdirnames(-1)
    if err != nil {
        return err
    }
    for _, name := range names {
        err = os.RemoveAll(filepath.Join(dir, name))
        if err != nil {
            return err
        }
    }
    return nil
}

func main() {
    dir := strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
    dir = filepath.Join(os.TempDir(), dir)
    dirs := filepath.Join(dir, `tmpdir`)
    err := os.MkdirAll(dirs, 0777)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    file := filepath.Join(dir, `tmpfile`)
    f, err := os.Create(file)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    f.Close()
    file = filepath.Join(dirs, `tmpfile`)
    f, err = os.Create(file)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    f.Close()

    err = RemoveContents(dir)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}