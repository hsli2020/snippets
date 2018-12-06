package main

import (
    "fmt"
    "os"
    "error"
)

func main() {
    // To check if a file doesn't exist, 

    if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
      // path/to/whatever does not exist
    }

    // In the above example we are not checking if err != nil because os.IsNotExist(nil) == false.

    // To check if a file exists

    if _, err := os.Stat("/path/to/whatever"); err == nil {
      // path/to/whatever exists
    }
}

// Exists reports whether the named file or directory exists.

// this code returns true, even if the file does not exist, for example 
// when Stat() returns permission denied.
func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return true
}

// good
func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

// better
func Exists(name string) (bool, error) {
    err := os.Stat(name)
    if os.IsNotExist(err) {
        return false, nil
    }
    return err != nil, err
}
