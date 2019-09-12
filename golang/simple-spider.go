// ԭ�����ӣ�https://blog.csdn.net/qq_25504271/article/details/78569643
package main

import (
    "net/http"
    "github.com/PuerkitoBio/goquery"
    "log"
    "fmt"
    "strconv"
    "os"
)

const (
    baseUrl string = "https://studygolang.com/topics?p="
)

func main() {

    var page int = 1
    var count int =getPageCount()

    for  {
        str := baseUrl + strconv.Itoa(page)
        response := getResponse(str)
        if (response.StatusCode == 403) {
            fmt.Println("IP �ѱ���ֹ����")
            os.Exit(1)
        }
        if (response.StatusCode == 200) {
            dom, err := goquery.NewDocumentFromResponse(response)
            if err != nil {
                log.Fatalf("ʧ��ԭ��", response.StatusCode)
            }
            dom.Find(".topics .topic").Each(func(i int, content *goquery.Selection) {
                title := content.Find(".title a").Text()
                fmt.Println(title)
            })
        }
        page++
        if page >= count{
            fmt.Println("������ȡ��ɹ�"+strconv.Itoa(page)+"��")
            os.Exit(1)
        }
    }
}

/**
* ����response
*/
func getResponse(url string) *http.Response {
    client := &http.Client{}
    request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
    response, _ := client.Do(request)
    return response
}

/**
* �õ���������
*/
func getPageCount() int {
    response := getResponse(baseUrl)
    dom, err := goquery.NewDocumentFromResponse(response)
    if err != nil {
        log.Fatalf("ʧ��ԭ��", response.StatusCode)
    }
    resDom := dom.Find(".text-center .pagination-sm li a")
    //len := resDom.Length()
    count,_ := strconv.Atoi(resDom.Eq(resDom.Length()-2).Text())
    return count
}
