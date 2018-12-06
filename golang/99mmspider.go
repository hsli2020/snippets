package main

import (
	"fmt"
	"strings"
	"regexp"
	"net/http"
	"io/ioutil"

	"github.com/PuerkitoBio/goquery"
	"strconv"
	"os"
	"io"
	"bytes"
	"sync"
	"log"
)

const baseUrl = "http://www.99mm.me/"

var (
	rwg sync.WaitGroup
	basePath string
	cFlag int
	flagLabel []string = []string{"hot", "meitui", "xinggan", "qingchun"}
)

func main() {
	fmt.Println("\n欢迎使用本工具，此工具仅限个人使用。")
	fmt.Println("\n\n请输入图片保存目录路径和内容抓取类别 ")
	fmt.Println("\n\n类别1 - hot")
	fmt.Println("类别2 - meitui")
	fmt.Println("类别3 - xinggan")
	fmt.Println("类别4 - qingchun")
	fmt.Print("\n\n")

	for {
		fmt.Println("请输入合法的图片保存目录路径（如 /usr/image、E:\\image）：")
		fmt.Scanln(&basePath)

		if basePath != "" {
			_, err := os.Stat(basePath)
			
			if err == nil {
				break
			}

			if os.IsNotExist(err) {
				if err = os.MkdirAll(basePath, 0777); err == nil {
					break
				} 
			}

			fmt.Println(err.Error())
		}
	}

	fmt.Println("请输入抓取内容类别编码（数字）：")
	fmt.Scanln(&cFlag)

	cFlag -= 1

	if cFlag < 0 || cFlag > 3 {
		cFlag = 0
	}

	log.Println("MM Spider start!")
	request(baseUrl + flagLabel[cFlag], "")
	log.Println("MM Spider shutdown!")
}

func request(url string, referer string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 Safari/537.36")

	if strings.Contains(url, "php") {
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Referer", referer)
	} else {
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Set("Cache-Control", "max-age=0")
	}

	resp, err := client.Do(req)

	if err != nil {
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		log.Println("Parse page failed, SKIP! ", err.Error())
		return
	}

	if strings.Contains(url, "php") {
		rawText, _ := doc.Html()

		if rawText == "" {
			return
		}

		r, _ := regexp.Compile("src=\"(.+?)\"")
		imgs := r.FindAllStringSubmatch(rawText, 2)

		if imgs == nil {
			return
		}

		if !strings.Contains(referer, "url=") {
			saveImage(imgs[0][1], genFolder(url))
		}

		if len(imgs) >= 2 {
			saveImage(imgs[1][1], genFolder(url))
		}

		if referer == "" || len(imgs) < 2 {
			return
		}

		newUrl := ""

		if strings.Contains(referer, "url=") {
			pr, _ := regexp.Compile("url=(\\d+)")
			pstr := pr.FindStringSubmatch(referer)

			if pstr == nil || len(pstr) <= 0 {
				return
			}

			pold := pstr[1]
			pint, err := strconv.Atoi(pold)

			if err != nil {
				return
			}

			p := strconv.Itoa(pint + 1)
			newUrl = strings.Replace(referer, "url="+pold, "url="+p, 1)
		} else {
			newUrl = referer + "?url=2"
		}

		if newUrl != "" {
			request(url, newUrl)
		}
	} else if doc.Find("#piclist").Size() <= 0 {
		src, exists := doc.Find(".picdata").Find("script").Attr("src")

		if !exists || src == "" {
			return
		}

		src = strings.TrimPrefix(src, "/")
		request(baseUrl+src, url)
	} else {
		doc.Find("#piclist li dt a").Each(func(i int, s *goquery.Selection) {
			detailUrl, exists := s.Attr("href")

			if !exists || detailUrl == ""{
				return
			}

			detailUrl = strings.TrimPrefix(detailUrl, "/")

			rwg.Add(1)
			go func() {
				defer rwg.Done()
				request(baseUrl+detailUrl, "")
			}()
		})

		nextPage, exists := doc.Find(".page a.next").Attr("href")

		if !exists || nextPage == "" {
			return
		}

		if !strings.Contains(nextPage, baseUrl) {
			if strings.Index(nextPage, "/") == 0 {
				nextPage = baseUrl + strings.TrimPrefix(nextPage, "/")
			} else {
				nextPage = url[0:strings.LastIndex(url, "/")+1] + nextPage
			}
		}

		log.Println("Next page!")
		rwg.Wait()
		request(nextPage, "")
	}
}

func saveImage(imgUrl string, folder string) {
	path := basePath + "/" + folder
	err := os.MkdirAll(path, 0777)

	if err != nil {
		log.Println("Create dicretory failed! ", err.Error())
		return
	}

	path += imgUrl[strings.LastIndex(imgUrl, "/"):]
	resp, err := http.Get(imgUrl)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	out, err := os.Create(path)

	if err != nil {
		log.Println("Create file failed! ", err.Error())
		return
	}

	_, err = io.Copy(out, bytes.NewReader(body))

	if err != nil {
		log.Println("Save file failed! ", err.Error())
	} else {
		log.Println("Saved file: ", path)
	}
}

func genFolder(url string) string {
	detailPagePattern, _ := regexp.Compile("id=(\\d+)")

	if strings.Contains(url, "php") {
		m := detailPagePattern.FindStringSubmatch(url)

		if m != nil {
			return m[1]
		}

		return ""
	}

	token := ".me/"
	folder := url[strings.Index(url, token)+len(token):]
	folder = folder[0:strings.Index(folder, ".html")]

	return strings.Split(folder, "/")[1]
}