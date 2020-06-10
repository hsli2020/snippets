package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GrabArticleList(blogid int) {
	page := 1
	totalPages := 0

	pageUrl := GetUrl(blogid, page)
	fmt.Printf("%s\n\n", pageUrl)

	for {
		res, err := http.Get(pageUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Extract the article list
		doc.Find(".articleList .articleCell").Each(func(i int, s *goquery.Selection) {
			// For each article found, get the title and time and href
			title := s.Find(".atc_title").Text()
			href, _ := s.Find(".atc_title a").Attr("href")
			time := s.Find(".atc_tm").Text()
			href = "https://blog.wenxuecity.com" + href
			fmt.Printf("    %s - %s %s\n", strings.TrimSpace(title), strings.TrimSpace(time), href)
		})

		// Get total pages
		if totalPages == 0 {
			doc.Find(".paging .page:last-of-type").Each(func(i int, s *goquery.Selection) {
				href, _ := s.Find("a").Attr("href")
				totalPages = GetTotalPages(href)
			})
		}

		fmt.Printf("P#%d\n", page)
		if page >= totalPages {
			break
		}
		page += 1
		pageUrl = GetUrl(blogid, page)
	}
}

func GetUrl(blogid, pageno int) string {
	url := fmt.Sprintf("https://blog.wenxuecity.com/myblog/%d/all.html", blogid)
	if pageno > 1 {
		url = fmt.Sprintf("https://blog.wenxuecity.com/blog/frontend.php?page=%d&act=articleList&blogId=%d", pageno-1, blogid)
	}
	return url
}

func GetFilename(blogid, pageno int) string {
	return fmt.Sprintf("%d-p%d.html", blogid, pageno)
}

func UrlToFilename(url string) string {
	// https://blog.wenxuecity.com/myblog/48977/201911/3432.html
	filename := strings.TrimPrefix(url, "https://blog.wenxuecity.com/myblog/")
	filename = strings.Replace(filename, "/", "-", -1)
	return filename
}

func GetTotalPages(addr string) int {
	// https://blog.wenxuecity.com/blog/frontend.php?page=8&act=articleList&blogId=48977
	u, err := url.Parse(addr)
	if err != nil {
		return 0
	}

	params := u.Query()
	n, err := strconv.Atoi(params.Get("page"))
	if err != nil {
		return 0
	}
	return n
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	blogid := 48977
	GrabArticleList(blogid)
}
