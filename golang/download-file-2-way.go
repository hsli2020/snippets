// https://github.com/akshaykhairmode/learning-go/blob/main/file_download/download.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var (
	uri  = "https://file-examples-com.github.io/uploads/2017/10/file-example_PDF_1MB.pdf"
	base = filepath.Base(uri)
)

func main() {

	downloadFileMoreMemory()
	fileDetails("downloadFileMoreMemory")
	downloadFileLessMemory()
	fileDetails("downloadFileLessMemory")
}

//downloadFileLessMemory will not load all data in memory
func downloadFileLessMemory() {

	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fileHandle, err := os.OpenFile(base, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	_, err = io.Copy(fileHandle, resp.Body)
	if err != nil {
		panic(err)
	}

}

//downloadFileMoreMemory will load all the file data in memory. In case of big file it will be problem
func downloadFileMoreMemory() {

	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(base, data, 0644); err != nil {
		panic(err)
	}

}

func fileDetails(fn string) {

	defer os.Remove(base)

	info, err := os.Stat(base)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File Size for function %v is : %v\n", fn, info.Size())
}