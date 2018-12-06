package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.chinanews.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	utf8reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(utf8reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
