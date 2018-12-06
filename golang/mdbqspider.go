package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"io/ioutil"
	"os"
	"io"
	"bytes"
	"strings"
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func Star() {
	doc, err := goquery.NewDocument("http://md.itlun.cn/")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	doc.Find(".pic li a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		if href == "" {
			return
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			err := fetchPage(href)
			if err != nil {
				fmt.Println(err.Error(), href)
			}
		}()
	})
}

func fetchPage(href string) error {
	doc, err := goquery.NewDocument("http://md.itlun.cn" + href)
	if err != nil {
		return err
	}

	imgEl := doc.Find("#imgString a")
	imgLink, _ := imgEl.Find("img").Attr("src")
	if imgLink == "" {
		return nil
	}

	err = downloadImg(imgLink)
	if err != nil {
		fmt.Println(err.Error())
	}

	nextLink, _ := imgEl.Attr("href")
	if nextLink == "" {
		return nil
	}

	return fetchPage(nextLink)
}

func downloadImg(imgPath string) error {
	if strings.HasPrefix(imgPath, "//") {
		imgPath = "http:" + imgPath
	}

	filename := imgPath[strings.LastIndex(imgPath, "/")+1:]

	resp, err := http.Get(imgPath)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	out, err := os.Create("img/" + filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, bytes.NewReader(body))
	return err
}

func main() {
	os.Mkdir("img", 0777)
	Star()
	wg.Wait()
}
