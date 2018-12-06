package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

/* start Fetch.go*/
func Fetch(url string) ([]byte, error) {
	//resp,err:= http.Get(url)
	//
	//if err!=nil{
	//	return nil,err
	//}
	//
	//defer resp.Body.Close()
	//if resp.StatusCode != http.StatusOK{
	//	return nil,fmt.Errorf("Error: status code:%d",resp.StatusCode)
	//}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

/* end  Fetch.go*/
/* start Type.go*/
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

/* end Type.go*/

/* start parser/city.go  爬取城市下每一个用户和网址*/

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		println(string(m[1]))
		result.Items = append(result.Items, "User:"+string(m[2]))
		result.Requests = append(result.Requests, Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) ParseResult {
				return PaesrProfile(
					c, name)
			},
		})
	}

	return result

}

/* end parser/city.go */

/* start parser/citylist.go */

const cityListRe = `(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) ParseResult {
	re := regexp.MustCompile(cityListRe)

	matches := re.FindAllSubmatch(contents, -1)
	result := ParseResult{}
	//测试，限制10个城市
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(
			result.Requests, Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})

		limit--
		if limit == 0 {
			break
		}
	}

	return result
}

/* end parser/citylist.go */

/* start profile.go */
type Profile struct {
	Name          string
	Age           int
	Marry         string
	Constellation string
	Height        int
	Weight        int
	Salary        string
}

func (p Profile) String() string {
	return p.Name + " " + p.Marry + strconv.Itoa(p.Age) + "olds " + strconv.Itoa(p.Age) + "cm " + strconv.Itoa(p.Weight) + "kg "
}

/* end profile.go */

/* start parser/profile.go */

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var marry = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(已婚)</div>`)
var constellation = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(.*)座</div>`)
var height = regexp.MustCompile(`160cm`)
var weight = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
var salary = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)

//name为上一级传递过来的
func PaesrProfile(contents []byte, name string) ParseResult {

	//ioutil.WriteFile("test.html",contents,0x777)

	profile := Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, height))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weight))
	if err == nil {
		profile.Weight = weight
	}

	profile.Salary = extractString(contents, salary)

	profile.Constellation = extractString(contents, constellation)
	if extractString(contents, marry) == "" {
		profile.Marry = "未婚"
	} else {
		profile.Marry = "已婚"
	}

	result := ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

/* end parser/profile.go */

/* start engine.go */
func Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]

		requests = requests[1:]
		fmt.Printf("Fetching %s", r.Url)
		body, err := Fetch(r.Url)

		if err != nil {
			log.Printf("Fetcher:error "+"fetching url %s, : %v", r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			fmt.Printf("Got item %s\n", item)
		}
	}
}

/* end engine.go */

func main() {

	Run(Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: ParseCityList,
	})
	//paseTest()
}

func paseTest() {

	contents, _ := ioutil.ReadFile("test.html")

	profile := Profile{}
	age, err := strconv.Atoi(extractString(contents, ageRe))

	if err != nil {

		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, height))
	if err != nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weight))
	if err != nil {
		profile.Weight = weight
	}

	profile.Salary = extractString(contents, salary)

	profile.Constellation = extractString(contents, constellation)
	if extractString(contents, marry) == "" {
		profile.Marry = "未婚"
	} else {
		profile.Marry = "已婚"
	}

	fmt.Printf("%s", profile)
}
