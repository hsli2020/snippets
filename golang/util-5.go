package sailor

import (
	"reflect"
	"unsafe"
)

// b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// s2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}



package sailor

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Exec(bin string, args []string) {
	binPath, err := exec.LookPath(bin)
	if err != nil {
		return
	}

	cmd := exec.Command(binPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return
	}

	err = cmd.Wait()
	if err != nil {
		log.Println("exec failed, err: ", err)

		if exitError, ok := err.(*exec.ExitError); ok {
			os.Exit(exitError.ExitCode())
		}
	}
}

func ExecCommand(bin string, args []string) (string, error) {
	binPath, err := exec.LookPath(bin)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(binPath, args...)

	var out bytes.Buffer

	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.Trim(out.String(), "\n"), nil
}



package sailor

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	FilePerm600 os.FileMode = 0600 // For secret files.
	FilePerm644 os.FileMode = 0644 // For normal files.
	FilePerm755 os.FileMode = 0755 // For directory or execute files.
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func EnsureFolderExists(folder string) {
	if Exists(folder) {
		return
	}

	err := os.MkdirAll(folder, FilePerm755)
	if err != nil {
		log.Fatal("directory not exists, err: ", err)
	}
}

func EnsureFileExists(path string) {
	if Exists(path) {
		return
	}

	EnsureFolderExists(filepath.Dir(path))

	err := ioutil.WriteFile(path, nil, FilePerm644)
	if err != nil {
		log.Fatal("file not exists, err: ", err)
	}
}

func WriteFile(path string, content string) error {
	EnsureFileExists(path)

	err := ioutil.WriteFile(path, StringToBytes(content), FilePerm644)
	if err != nil {
		return err
	}

	return nil
}
