package main

import (
	"io/ioutil"
	"net/http"
	"os/user"
	"strings"

	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const baseUrl = "http://www.mzitu.com/"

var (
	rwg      sync.WaitGroup
	basePath string
	cFlag    int
	cPage    int

	// 频道
	flagLabel = []string{
		"",
		"hot",
		"best",
		"zhuanti",
		"xinggan",
		"japan",
		"taiwan",
		"mm",
		"zipai",
	}
	flagDesc = []string{
		"最新",
		"最热",
		"推荐",
		"专题",
		"性感",
		"日本",
		"台湾",
		"清纯",
		"自拍",
	}

	// 浏览器代理
	userAgents = []string{
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
	}

	userAgentsLen = len(userAgents)

	flagLen = len(flagDesc)

	// 失败重试最大次数
	maxRetry = 2
)

type Task struct {
	Url     string
	Referer string
	Retry   int
	Folder  string
}

func main() {
	fmt.Println("\n欢迎使用《司机之友》，本程序由 Golang 强力驱动。")
	fmt.Println("\n\n请输入图片保存目录路径和内容抓取类别，默认为首页推荐 ")
	fmt.Print("\r\n\r\n")

	for i := 0; i < flagLen; i++ {
		fmt.Println(i, " - ", flagDesc[i])
	}

	fmt.Print("\r\n\r\n")

	for {
		fmt.Println("请输入合法的图片保存目录路径（如 /usr/image、E:\\image）：")
		fmt.Scanln(&basePath)
		basePath = changePathToHomeDir(strings.TrimSpace(basePath))
		if basePath != "" {
			if _, err := os.Stat(basePath); err == nil {
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

	fmt.Println("请输入抓取内容类别编码（数字，默认为 0）：")
	fmt.Scanln(&cFlag)
	fmt.Println("请输入起始抓取页码（数字，默认为 1）：")
	fmt.Scanln(&cPage)
	if cFlag < 0 || cFlag > flagLen-1 {
		cFlag = 0
	}

	p := flagLabel[cFlag]
	url := baseUrl
	if p != "" {
		url += p + "/"
	}
	if cPage > 0 {
		url += "page/" + strconv.Itoa(cPage)
	}

	log.Println("Task start!")
	err := execute(&Task{
		Url:    url,
		Folder: filterFilename(flagDesc[cFlag]),
	})
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Task shutdown!")
	log.Println("Press <Enter> key to exit...")
	fmt.Scan()
	os.Exit(0)
}

func execute(task *Task) error {
	resp, err := request(task)
	if err != nil {
		if err = retry(task); err != nil {
			return fmt.Errorf("request %s failed, %v", task.Url, err)
		}
		return nil
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		if err = retry(task); err != nil {
			return fmt.Errorf("parse page failed, SKIP! %s, %v", task.Url, err.Error())
		}
		return nil
	}

	var currentNode *goquery.Selection
	if currentNode = doc.Find(".main-image"); currentNode.Size() > 0 {
		imgUrl, _ := currentNode.Find("p img").Attr("src")
		desc, _ := currentNode.Find("p img").Attr("alt")
		if imgUrl != "" {
			saveImage(task.Url, imgUrl, filepath.Join(task.Folder, filterFilename(desc)))
		}

		var nextUrl string
		selection := doc.Find(".pagenavi").Find("a").Last()
		text := selection.Find("span").Text()
		if strings.TrimSpace(text) == "下一页»" {
			nextUrl, _ = selection.Attr("href")
		}
		if nextUrl == "" {
			return nil
		}

		time.Sleep(time.Duration(1) * time.Second)
		err = execute(&Task{
			Url:    nextUrl,
			Folder: task.Folder,
		})
	} else if currentNode = doc.Find(".postlist"); currentNode.Size() > 0 {
		if tags := currentNode.Find("dl.tags"); tags.Size() > 0 {
			curTitle := ""
			tags.Children().Each(func(i int, s *goquery.Selection) {
				if s.Is("dt") {
					curTitle = s.Text()
					return
				}
				a := s.Find("a")
				tagLink, _ := a.Attr("href")
				tagAlt := a.Text()
				if tagLink != "" {
					e := execute(&Task{
						Url:    tagLink,
						Folder: filepath.Join(task.Folder, filterFilename(curTitle), filterFilename(tagAlt)),
					})
					if e != nil {
						fmt.Println(e.Error())
					}
				}
			})
			return nil
		}

		var index = 0
		var nextPage string
		// 使用此循环方式来避免无法直接获取到 nav 元素
		currentNode.Children().Each(func(i int, s *goquery.Selection) {
			if index == 0 {
				s.Find("li>a").Each(func(idx int, sel *goquery.Selection) {
					detailUrl, exists := sel.Attr("href")
					if !exists || detailUrl == "" {
						return
					}

					detailUrl = strings.TrimPrefix(detailUrl, "/")
					rwg.Add(1)
					go func() {
						defer rwg.Done()
						// 缓冲执行协程，防止过快一起执行
						time.Sleep(time.Duration(i) * time.Second)
						e := execute(&Task{
							Url:    detailUrl,
							Folder: task.Folder,
						})
						if e != nil {
							fmt.Println(e.Error())
						}
					}()
				})
			} else if index == 1 {
				nextPage, _ = s.Find(".next").Attr("href")
			}
			index++
		})

		if nextPage == "" {
			log.Println("Page end!", task.Url)
			return nil
		}

		rwg.Wait()
		log.Println("Next page ", nextPage)
		err = execute(&Task{
			Url:    nextPage,
			Folder: task.Folder,
		})
	} else {
		return retry(task)
	}

	return err
}

// 请求指定的 task
func request(task *Task) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", task.Url, nil)
	if err != nil {
		return nil, err
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	req.Header.Set("User-Agent", userAgents[r.Intn(userAgentsLen)])
	if task.Referer != "" {
		req.Header.Set("Accept", "image/webp,image/*,*/*;q=0.8")
		req.Header.Set("Referer", task.Referer)
	} else {
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Referer", baseUrl)
	}

	return client.Do(req)
}

// 保存图片资源，分目录层次
func saveImage(pageUrl string, imgUrl string, folder string) {
	fragment := strings.TrimPrefix(pageUrl, baseUrl)
	fragments := strings.Split(fragment, "/")
	idx := "1"
	fragLen := len(fragments)
	if fragLen == 0 {
		return
	}
	if fragLen > 1 {
		idx = fragments[1]
	}

	folder = strings.TrimSpace(folder) + "_" + fragments[0]
	p := filepath.Join(basePath, folder)
	err := os.MkdirAll(p, 0777)
	if err != nil {
		log.Println("Create directory failed! ", p, err.Error())
		return
	}

	ext := filepath.Ext(imgUrl)
	p = filepath.Join(p, idx+ext)
	resp, err := request(&Task{
		Url:     imgUrl,
		Referer: pageUrl,
	})
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read response faield! ", err.Error())
		return
	}

	out, err := os.Create(p)
	if err != nil {
		log.Println("Create file failed! ", err.Error())
		return
	}

	bs, err := io.Copy(out, bytes.NewReader(body))
	if err != nil {
		log.Println("Save file failed! ", err.Error())
	} else {
		log.Printf("Saved file [%d bytes]: %s", bs, p)
	}
}

// 失败重试判断
func retry(task *Task) error {
	if task.Retry >= maxRetry {
		return fmt.Errorf("invalid page %s", task.Url)
	}
	task.Retry++
	time.Sleep(time.Duration(2) * time.Second)
	return execute(task)
}

// 过滤非法文件名
func filterFilename(name string) string {
	name = strings.TrimSpace(name)
	reg, err := regexp.Compile("[\\\\/:*?\"<>|]")
	if err != nil {
		return name
	}
	return reg.ReplaceAllString(name, "_")
}

// 把用户输入的保存路径中的 ~ 符号替换为当前用户目录
func changePathToHomeDir(path string) string {
	if strings.Index(path, "~") != 0 {
		return path
	}
	u, err := user.Current()
	if err != nil {
		return path
	}
	return u.HomeDir + path[1:]
}
