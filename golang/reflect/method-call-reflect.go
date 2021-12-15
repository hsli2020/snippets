package main

import (
	"fmt"
	"reflect"
	"sync"
)

type Spider struct {
}

func (s Spider) GetZhihu() map[string]interface{} {
	return map[string]interface{}{"From": "--Get Zhihu"}
}

func (s Spider) GetWeibo() map[string]interface{} {
	return map[string]interface{}{"From": "--Get Weibo"}
}

func (s Spider) GetTieba() map[string]interface{} {
	return map[string]interface{}{"From": "--Get Tieba"}
}

func (s Spider) GetDouban() map[string]interface{} {
	return map[string]interface{}{"From": "--Get Douban"}
}

func (s Spider) GetHupu() map[string]interface{} {
	return map[string]interface{}{"From": "--Get Hupu"}
}

func ExecCrawl(siteName string) {
	spider := Spider{}
	reflectValue := reflect.ValueOf(spider)
	method := reflectValue.MethodByName("Get" + siteName)
	result := method.Call(nil)
	data := result[0].Interface().(map[string]interface{})
	fmt.Println(data["From"])
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	allSites := []string{
		"Zhihu",
		"Weibo",
		"Tieba",
		"Douban",
		"Hupu",
	}

	wg.Add(len(allSites))

	for _, site := range allSites {
		fmt.Println("Crawl" + site)
		go ExecCrawl(site)
	}

	wg.Wait()
	fmt.Println("End")
}
