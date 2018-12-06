// Golang: File Tree Traversal (filepath.Walk)
// 
// In this article we'll see how to walk the file system with golang. We'll see:
// 
//     a simple example of filepath.Walk
//     how to pass arguments to filepath.WalkFunc
//     how to find file duplicates
//     a du implementation
// 
// filepath.Walk
// 
// As a simple example of filepath.Walk we'll list all the files under a directory 
// recursively (simple.go):

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func printFile(path string, info os.FileInfo, err error) error {
    if err != nil {
        log.Print(err)
        return nil
    }
    fmt.Println(path)
    return nil
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1]
    err := filepath.Walk(dir, printFile)
    if err != nil {
        log.Fatal(err)
    }
}

// We set the log flags to Lshortfile to better spot errors when they happen. Everything
// else is explained very well in the go docs.
// 
// Run it with:
// 
// $ go run simple.go .
// 
// Passing arguments
// 
// We can pass arguments to filepath.WalkFunc trought a closure. printFile now doesn't
// process the files directly but returns a closure that does the work. The closure can
// access the arguments we pass to printFile as if they were local variables. For example 
// to ignore some directories we can pass a list of said directories to printFile and
// whenever the closure finds a directory whose name is inside the list it will skip it
// by returning os.SkipDir (ignore.go):

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func printFile(ignoreDirs []string) filepath.WalkFunc {
    return func(path string, info os.FileInfo, err error) error {
        if err != nil {
            log.Print(err)
            return nil
        }
        if info.IsDir() {
            dir := filepath.Base(path)
            for _, d := range ignoreDirs {
                if d == dir {
                    return filepath.SkipDir
                }
            }
        }
        fmt.Println(path)
        return nil
    }
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1]
    ignoreDirs := []string{".bzr", ".hg", ".git"}
    err := filepath.Walk(dir, printFile(ignoreDirs))
    if err != nil {
        log.Fatal(err)
    }
}

// Find file duplicates
// 
// For a more realistic application we'll write a program that will find all the file 
// duplicates under a directory. For each file we'll store its crypto/sha512 digest
// inside a map. If the digest was already present, the file is a duplicate, otherwise
// we store its path using the digest as a key (fdup.go):

package main

import (
    "crypto/sha512"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

var files = make(map[[sha512.Size]byte]string)

func checkDuplicate(path string, info os.FileInfo, err error) error {
    if err != nil {
        log.Print(err)
        return nil
    }
    if info.IsDir() {
        return nil
    }

    data, err := ioutil.ReadFile(path)
    if err != nil {
        log.Print(err)
        return nil
    }
    digest := sha512.Sum512(data)
    if v, ok := files[digest]; ok {
        fmt.Printf("%q is a duplicate of %q\n", path, v)
    } else {
        files[digest] = path
    }

    return nil
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1]
    err := filepath.Walk(dir, checkDuplicate)
    if err != nil {
        log.Fatal(err)
    }
}

// Let's try it
// 
// In a terminal run:
// 
// $ mkdir test
// $ cd test
// $ echo 'run free, run GNU' > gnu
// $ echo 'from outer space' > plan9
// $ cp gnu free
// $ cp plan9 outer
// $ ls
// free  gnu  outer  plan9
// $ cd ..
// $ go run fdup.go test
// "test/gnu" is a duplicate of "test/free"
// "test/plan9" is a duplicate of "test/outer"

// du
// 
// Despite filepath.Walk's usefulness it can not model all type of programs, one such program
// is du. Starting with one directory, du reports the cumulative size of the given directory
// and all its subdirectories recursively. The entries of a directory are read with os.Readdir 
// (du.go):

package main

import (
    "fmt"
    "log"
    "os"
)

func du(currentPath string, info os.FileInfo) int64 {
    size := info.Size()
    if !info.IsDir() {
        return size
    }

    dir, err := os.Open(currentPath)
    if err != nil {
        log.Print(err)
        return size
    }
    defer dir.Close()

    fis, err := dir.Readdir(-1)
    if err != nil {
        log.Fatal(err)
    }
    for _, fi := range fis {
        if fi.Name() == "." || fi.Name() == ".." {
            continue
        }
        size += du(currentPath+"/"+fi.Name(), fi)
    }

    fmt.Printf("%d %s\n", size, currentPath)

    return size
}

func main() {
    log.SetFlags(log.Lshortfile)
    dir := os.Args[1]
    info, err := os.Lstat(dir)
    if err != nil {
        log.Fatal(err)
    }
    du(dir, info)
}
