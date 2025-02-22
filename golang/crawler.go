package main

import (
	"io/ioutil"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/unicode"
	"log"
	"regexp"
	"strconv"
	"net/http"
	"fmt"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

/* start Fetch.go*/
func Fetch(url string)([]byte ,error){
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

	bodyReader:= bufio.NewReader(resp.Body)
	e:= determineEncoding(bodyReader)
	utf8reader:= transform.NewReader(bodyReader,e.NewDecoder())

	return ioutil.ReadAll(utf8reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding{

	bytes,err := bufio.NewReader(r).Peek(1024)
	if err !=nil{
		log.Printf("Fetcher error:%v",err)
		return unicode.UTF8
	}
	e,_,_:= charset.DetermineEncoding(bytes,"")
	return e
}
/* end  Fetch.go*/

/* start Type.go*/
type Request struct{
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct{
	Requests []Request
	Items []interface{}
}

func NilParser([]byte) ParseResult{
	return ParseResult{}
}
/* end Type.go*/

/* start parser/city.go  爬取城市下每一个用户和网址*/
const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) ParseResult{
	re:= regexp.MustCompile(cityRe)
	matches:= re.FindAllSubmatch(contents,-1)

	result := ParseResult{}
	for _,m:= range matches{
		name:=string(m[2])
		println(string(m[1]))
		result.Items = append(result.Items,"User:"+string(m[2]))
		result.Requests = append(result.Requests,Request{
			Url:string(m[1]),
			ParserFunc:func(c []byte) ParseResult{
				return PaesrProfile(
					c,name)
			},
		})
	}
	return result
}
/* end parser/city.go */

/* start parser/citylist.go */

const cityListRe = `(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) ParseResult{
	re:=regexp.MustCompile(cityListRe)

	matches:= re.FindAllSubmatch(contents,-1)
	result:=ParseResult{}

	for _,m :=range matches{
		result.Items = append(result.Items,string(m[2]))
		result.Requests  = append(
			result.Requests,Request{
				Url:string(m[1]),
				ParserFunc:ParseCity,
			}
		)
	}

	return result
}
/* end parser/citylist.go */

/* start profile.go */
type Profile struct {
	Name string
	Age int
	Marry string
	Constellation string
	Height int
	Weight int
	Salary string
}

func (p Profile) String() string{
	return  p.Name +" " + p.Marry + strconv.Itoa(p.Age) +"olds "+ 
		strconv.Itoa(p.Age) + "cm " +  strconv.Itoa(p.Weight)+ "kg "
}
/* end profile.go */

/* start parser/profile.go */

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)岁</div>`)
var marry =   regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(已婚)</div>`)
var constellation =   regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>(.*)座</div>`)
var height  =regexp.MustCompile(`160cm`)
var weight =regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\d]+)kg</div>`)
var salary = 	regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>月收入:([^<]+)</div>`)

//解析器 解析用户
//name为上一级传递过来的
func PaesrProfile(contents []byte,name string) ParseResult{
	//ioutil.WriteFile("test.html",contents,0x777)
	//用户结构体
	profile:=Profile{}
	profile.Name = name

	//年龄   string转换为int
	age,err:= strconv.Atoi(extractString(contents,ageRe))
	if err==nil{
		profile.Age = age
	}
		//身高
	height,err:= strconv.Atoi(extractString(contents,height))
	if err==nil{
		profile.Height = height
	}
	//体重
	weight,err:= strconv.Atoi(extractString(contents,weight))
	if err==nil{
		profile.Weight = weight
	}

	//薪水
	profile.Salary = extractString(contents,salary)

	//星座
	profile.Constellation = extractString(contents,constellation)
	if extractString(contents,marry)== ""{
		profile.Marry ="未婚"
	}else{
		profile.Marry ="已婚"
	}

	result:=ParseResult{
		Items:[]interface{}{profile},
	}

	return result
}

//封装 正则表达式匹配
func extractString(contents []byte,re *regexp.Regexp) string{
	match:=re.FindSubmatch(contents)

	if len(match)>=2{
		return string(match[1])
	}else{
		return ""
	}
}
/* end parser/profile.go */

/* start engine.go   单任务版引擎*/
func Run(seeds ...Request){
	var requests []Request

	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) >0{
		r:=requests[0]

		requests = requests[1:]
		fmt.Printf("Fetching %s",r.Url)
		body,err:= Fetch(r.Url)

		if err!=nil{
			log.Printf("Fetcher:error "+ "fetching url %s, : %v",r.Url,err)
			continue
		}

		parseResult:= r.ParserFunc(body)

		requests = append(requests,parseResult.Requests...)

		for _,item:= range parseResult.Items{
			fmt.Printf("Got item %s\n",item)
		}
	}
}

//具体的工作  传递一个request，通过解析器对url进行解析
func worker(r Request)(ParseResult,error){
	fmt.Printf("Fetching %s\n",r.Url)
	body,err:= Fetch(r.Url)

	if err!=nil{
		log.Printf("Fetcher:error "+ "fetching url %s, : %v",r.Url,err)
		return ParseResult{},err
	}
	return 	r.ParserFunc(body),nil
}

// 并发版爬虫引擎  包含了调度器 与 工人数
type ConcurrentEngine struct{
	Scheduler Scheduler
	WorkerCount int
}

//调度器是一个接口，扩展性
type Scheduler interface {
	//提交Request进行执行
	Submit(Request)
	//配置通道
	ConfigureMasterWorkChan(chan Request)
}

//并发版爬虫引擎
func (e *ConcurrentEngine) Run(seeds ...Request){
	in:=make(chan Request)
	out:= make(chan ParseResult)
	//配置调度器通道
	e.Scheduler.ConfigureMasterWorkChan(in)

	//开启WorkerCount个工作
	for i:=0;i<e.WorkerCount;i++{
		createWorker(in,out)
	}
		//种子首先运行
	for _,r:=range seeds{
		e.Scheduler.Submit(r)
	}
	itemsCount:=0
	for{
		//out等待接受ParseResult
		result:=<-out

		//打印出接收到的数据，以及个数。
		for _,item:= range result.Items{
			fmt.Printf("Got item: #%d: %v\n",itemsCount,item)
			itemsCount++
		}

		//每一个request都通过调度器发送给in通道，由于在调度器中使用了协程的方式，不会卡死在这里陷入死锁。
		for _,request:= range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

//工作函数，逻辑是 in通道接收到request，即会调用worker函数爬每一个request中的网址，用对应的解析器。 
//解析完成后，将ParseResult返回给通道out
func createWorker(in chan Request, out chan ParseResult) {
	go func(){
		for{
			request := <-in
			result,err:= worker(request)
			if err!=nil{
				continue
			}
			out <-result
		}
	}()
}
/* end engine.go */

/* start scheduler.go 简单版调度器，用于分配工作任务 */
type SimpleScheduler struct{
	//通道
	workerChan chan Request
}

func ( s *SimpleScheduler) Submit( r Request) {
	//为了防止死锁，在调度器中建立go的协程  分配任务到通道中。
	go func(){s.workerChan <- r}()
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan Request) {
	//in通道
	s.workerChan = c
}
/* end scheduler.go */

func main(){
	//并发版爬虫
	e:= ConcurrentEngine{
		Scheduler:&SimpleScheduler{},
		WorkerCount:100,
	}

	e.Run(Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:ParseCityList,
	})

	//单任务版爬虫
	//Run(Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:ParseCityList,
	//})
	//paseTest()
}

//测试
func paseTest(){
	contents,_:= ioutil.ReadFile("test.html")

	profile:=Profile{}
	age,err:= strconv.Atoi(extractString(contents,ageRe))

	if err!=nil{
		profile.Age = age
	}

	height,err:= strconv.Atoi(extractString(contents,height))
	if err!=nil{
		profile.Height = height
	}

	weight,err:= strconv.Atoi(extractString(contents,weight))
	if err!=nil{
		profile.Weight = weight
	}

	profile.Salary = extractString(contents,salary)

	profile.Constellation = extractString(contents,constellation)
	if extractString(contents,marry)== ""{
		profile.Marry ="未婚"
	}else{
		profile.Marry ="已婚"
	}

	fmt.Printf("%s",profile)
}