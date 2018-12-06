package main

import (
    "fmt"
    "os/exec"
    "io/ioutil"
)

func main() {
    arg := []string{"-l", "-a"}
    fmt.Println(string(pipe("dir", arg, "")))
}

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func pipe(bin string, arg []string, src string) []byte {
    cmd := exec.Command(bin, arg...)
    in, err := cmd.StdinPipe();         check(err)
    out, err := cmd.StdoutPipe();       check(err)
    err = cmd.Start();                  check(err)
    _, err = in.Write([]byte(src));     check(err)
    err = in.Close();                   check(err)
    bytes, err := ioutil.ReadAll(out);  check(err)
    err = cmd.Wait();                   check(err)
    return bytes
}
