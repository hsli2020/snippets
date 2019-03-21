package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func MakeRequest(client http.Client, url string, ch chan<- string) {
    reqest, _ := http.NewRequest("GET", url, nil)

    response, _ := client.Do(reqest)
    body, _ := ioutil.ReadAll(response.Body)
    ch <- string(body)
}

func main() {
    start := time.Now()
    urls := []string{"https://sina.com.cn", "https://news.baidu.com"}
    client := &http.Client{}
    ch := make(chan string, len(urls))
    for _, url := range urls {
        go MakeRequest(*client, url, ch)
    }

    for range urls {
        fmt.Println(<-ch)
    }

    fmt.Printf("%.8fs elapsed\n", time.Since(start).Seconds())
}
