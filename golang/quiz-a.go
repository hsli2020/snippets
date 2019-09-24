package main

import (
	"os"
	"io"
	"strings"
	"os/exec"
)

func main() {
	r := strings.NewReader("foo\n")
	cmd := exec.Command("go", "version")
	cmd.Stdin = r
	cmd.Run()
	io.Copy(os.Stdout, r)
}
