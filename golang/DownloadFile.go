package main

import (
    "io"
    "net/http"
    "os"
)

func main() {

    fileUrl := "https://golangcode.com/images/avatar.jpg"

    err := DownloadFile("avatar.jpg", fileUrl)
    if err != nil {
        panic(err)
    }

}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

    // Create the file
    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    // Get the data
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return err
    }

    return nil
}
/*
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func download(url, filename string) {
	fmt.Println("Downloading " + url + " ...")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)
}

func main() {
	pUrl := flag.String("url", "", "URL to be processed")
	flag.Parse()
	url := *pUrl
	if url == "" {
		fmt.Fprintf(os.Stderr, "Error: empty URL!\n")
		return
	}

	filename := path.Base(url)
	fmt.Println("Checking if " + filename + " exists ...")
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		download(url, filename)
		fmt.Println(filename + " saved!")
	} else {
		fmt.Println(filename + " already exists!")
	}
}*/