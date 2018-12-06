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
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
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
	//fmt.Printf("%s\n",all)
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {

	//正则匹配网址http://www.zhenai.com/zhenghun/xiamen" data-v-4e064b2c>厦门</a>
	//[^>]代表以>结尾，*>代表到达之前>之前的东西，可能有换行符。
	re := regexp.MustCompile(`http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)

	matches := re.FindAll(contents, -1)

	for _, m := range matches {
		fmt.Printf("%s\n", m)
	}

	fmt.Printf("Matches found: %d\n", len(matches))
}
