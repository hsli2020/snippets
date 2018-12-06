package fileutil

import (
    "os"
)

func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}
