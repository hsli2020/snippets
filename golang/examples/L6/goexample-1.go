func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}
------------------------------------------------------------
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
------------------------------------------------------------
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
------------------------------------------------------------
json.NewEncoder(w).Encode(todos)
------------------------------------------------------------
kvs := map[string]string{"a": "apple", "b": "banana"}
for k, v := range kvs {
    fmt.Printf("%s -> %s\n", k, v)
}
------------------------------------------------------------
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
------------------------------------------------------------
package main

import "fmt"
import "math"

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)
}
------------------------------------------------------------
package main

import "sort"
import "fmt"

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
------------------------------------------------------------
package main

import "strings"
import "fmt"

func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))

    fmt.Println(Map(strs, strings.ToUpper))
}
------------------------------------------------------------
package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {

    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
------------------------------------------------------------
package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}
    fmt.Printf("%v\n", p)      //  {1 2}
    fmt.Printf("%+v\n", p)     //  {x:1 y:2}
    fmt.Printf("%#v\n", p)     //  main.point{x:1, y:2}
    fmt.Printf("%T\n", p)      //  main.point
    fmt.Printf("%t\n", true)   //  true
    fmt.Printf("%d\n", 123)    //  123
    fmt.Printf("%b\n", 14)     //  1110
    fmt.Printf("%c\n", 33)     //  !
    fmt.Printf("%x\n", 456)    //  1c8
    fmt.Printf("%f\n", 78.9)   //  78.900000
    fmt.Printf("%e\n", 123400000.0)    //  1.234000e+08
    fmt.Printf("%E\n", 123400000.0)    //  1.234000E+08
    fmt.Printf("%s\n", "\"string\"")   //  "string"
    fmt.Printf("%q\n", "\"string\"")   //  "\"string\""
    fmt.Printf("%x\n", "hex this")     //  6865782074686973
    fmt.Printf("%p\n", &p)         //  0x42135100
    fmt.Printf("|%6d|%6d|\n", 12, 345)           //  |    12|   345|
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)     //  |  1.20|  3.45|
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)   //  |1.20  |3.45  |
    fmt.Printf("|%6s|%6s|\n", "foo", "b")        //  |   foo|     b|
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")      //  |foo   |b     |

    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)   //  a string
    fmt.Fprintf(os.Stderr, "an %s\n", "error")   //  an error
}
------------------------------------------------------------
package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)  // true

    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))  // true
    fmt.Println(r.FindString("peach punch"))  // peach
    fmt.Println(r.FindStringIndex("peach punch"))  // [0 5]
    fmt.Println(r.FindStringSubmatch("peach punch"))  // [peach ea]
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))  // [0 5 1 3]
    fmt.Println(r.FindAllString("peach punch pinch", -1))  // [peach punch pinch]

    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
        // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
    fmt.Println(r.FindAllString("peach punch pinch", 2))  // [peach punch]
    fmt.Println(r.Match([]byte("peach")))  // true

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)  // p([a-z]+)ch
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))  // a <fruit>

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))  // a PEACH
}
------------------------------------------------------------
package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)  // 2012-10-31 15:50:13.793654 +0000 UTC

    then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)  // 2009-11-17 20:34:58.651387237 +0000 UTC

    p(then.Year())            // 2009
    p(then.Month())           // November
    p(then.Day())             // 17
    p(then.Hour())            // 20
    p(then.Minute())          // 34
    p(then.Second())          // 58
    p(then.Nanosecond())      // 651387237
    p(then.Location())        // UTC

    p(then.Weekday())         // Tuesday

    p(then.Before(now))       // true
    p(then.After(now))        // false
    p(then.Equal(now))        // false

    diff := now.Sub(then)
    p(diff)                 // 25891h15m15.142266763s

    p(diff.Hours())         // 25891.25420618521
    p(diff.Minutes())       // 1.5534752523711128e+06
    p(diff.Seconds())       // 9.320851514226677e+07
    p(diff.Nanoseconds())   // 93208515142266763

    p(then.Add(diff))   // 2012-10-31 15:50:13.793654 +0000 UTC
    p(then.Add(-diff))  // 2006-12-05 01:19:43.509120474 +0000 UTC
}
------------------------------------------------------------
package main

import "fmt"
import "time"

func main() {

    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)  // 2012-10-31 16:13:58.292387 +0000 UTC

    millis := nanos / 1000000
    fmt.Println(secs)    // 1351700038
    fmt.Println(millis)  // 1351700038292
    fmt.Println(nanos)   // 1351700038292387000

    fmt.Println(time.Unix(secs, 0))   // 2012-10-31 16:13:58 +0000 UTC
    fmt.Println(time.Unix(0, nanos))  // 2012-10-31 16:13:58.292387 +0000 UTC
}
------------------------------------------------------------
package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339))  // 2014-04-15T18:00:15-07:00

    t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
    p(t1)  // 2012-11-01 22:08:41 +0000 +0000

    p(t.Format("3:04PM"))  // 6:00PM
    p(t.Format("Mon Jan _2 15:04:05 2006"))  // Tue Apr 15 18:00:15 2014
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))  // 2014-04-15T18:00:15.161182-07:00
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)  // 0000-01-01 20:41:00 +0000 UTC

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())  // 2014-04-15T18:00:15-00:00

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)  // parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": ...
}
------------------------------------------------------------
package main

import "fmt"
import "net"
import "net/url"

func main() {

    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    fmt.Println(u.Scheme)  // postgres

    fmt.Println(u.User)    // user:pass
    fmt.Println(u.User.Username())  // user
    p, _ := u.User.Password()
    fmt.Println(p)  // /path

    fmt.Println(u.Host)  // host.com:5432
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)  // host.com
    fmt.Println(port)  // 5432

    fmt.Println(u.Path)      // /path
    fmt.Println(u.Fragment)  // f

    fmt.Println(u.RawQuery)  // k=v
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)           // map[k:[v]]
    fmt.Println(m["k"][0])   // v
}
------------------------------------------------------------
package main

import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    h := sha1.New()
    h.Write([]byte(s))
    bs := h.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
------------------------------------------------------------
package main

import b64 "encoding/base64"
import "fmt"

func main() {

    data := "abc123!?$*&()'-=@~"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
------------------------------------------------------------
type Service struct {
    Name     string `json:"name"`
    Host     string `json:"host"`
    Port     uint   `json:"port"`
    RestPort uint   `json:"restPort"`
    WsPort   uint   `json:"wsPort"`
}

func LoadApiServers(filepath, env string) (map[string][]Service, error) {
    file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    configs := make(map[string]map[string][]Service, 0)
    err = json.NewDecoder(file).Decode(&configs)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return configs[env], err
}
------------------------------------------------------------
type Speech struct {
    Services []Service
}
type Sms struct {
    Services []Service
}
type Payment struct {
    Services []Service
}

type Environment struct {
    Speech interface{}
    Sms interface{}
    Payment interface{}
}

type Service struct {
    Name string `json:"name"`
    Host string `json:"host"`
    Port uint `json:"port"`
    RestPort uint `json:"restPort"`
    WsPort uint `json:"wsPort"`
}

type Config struct {
    Services []Service
    Master Service
    Mutex sync.RWMutex
}

func LoadApiServers(filepath, env string) (map[string][]Service, error) {
    file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    configs := make(map[string]map[string][]Service, 0)
    err = json.NewDecoder(file).Decode(&configs)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    return configs[env], err
}

func printEnvironment(envMap *Environment) {
    el := reflect.ValueOf(envMap).Elem()
    fldName := ""
    fldType := reflect.ValueOf(envMap).Elem().Type()
    var speech Speech
    var sms Sms
    var payment Payment
    for i := 0; i < el.NumField(); i++ {
        f := el.Field(i)
        fldName = fldType.Field(i).Name
        thisValAry := f.Interface().([]interface{})
        switch fldName {
        case "Speech":
            for _, connectorMap := range thisValAry {
                thisSvc := getService(connectorMap)
                speech.Services = append(speech.Services, *thisSvc)
            }
        case "Sms":
            for _, chatMap := range thisValAry {
                thisSvc := getService(chatMap)
                sms.Services = append(sms.Services, *thisSvc)
            }
        case "Payment":
            for _, gatetMap := range thisValAry {
                thisSvc := getService(gatetMap)
                payment.Services = append(payment.Services, *thisSvc)
            }
        }
    }
    environment := &Environment{
        Speech: speech,
        Sms: sms,
        Payment: payment,
    }
    fmt.Printf("environment: %+v\n", environment)
}

var Dir string

func HtmlFileHandler(response http.ResponseWriter, request *http.Request, filename string){
    response.Header().Set("Content-type", "text/html")
    webpage, err := ioutil.ReadFile(Dir + filename)  // read whole the file
    if err != nil {
        http.Error(response, fmt.Sprintf("%s file error %v", filename, err), 500)
    }
    fmt.Fprint(response, string(webpage));
}

func HelpHandler(response http.ResponseWriter, request *http.Request){
    HtmlFileHandler(response, request, "/help.html")
}

func AjaxHandler(response http.ResponseWriter, request *http.Request){
    HtmlFileHandler(response, request, "/ajax.html")
}

func DebugFormHandler(response http.ResponseWriter, request *http.Request){

    printCookies(response, request)

    err := request.ParseForm()  // Parse URL and POST data into request.Form
    if err != nil {
        http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
    }

    // Set cookie and MIME type in the HTTP headers.
    fmt.Printf("request.Form: %v\n", request.Form)
    if request.Form["username"] != nil {
        cookieVal := request.Form["username"][0]
        fmt.Printf("cookieVal: %s\n", cookieVal)
        SetUsernameCookie(response, cookieVal)
    }; fmt.Println("")

    templateHandler(response, request)
    response.Header().Set("Content-type", "text/plain")

    // Send debug diagnostics to client
    fmt.Fprintf(response, "request.Method     '%v'", request.Method)
    fmt.Fprintf(response, "request.RequestURI '%v'", request.RequestURI)
    fmt.Fprintf(response, "request.URL.Path   '%v'", request.URL.Path)
    fmt.Fprintf(response, "request.Form       '%v'", request.Form)
    fmt.Fprintf(response, "request.Cookies()  '%v'", request.Cookies())
}

func DebugQueryHandler(response http.ResponseWriter, request *http.Request){

    // Set cookie and MIME type in the HTTP headers.
    response.Header().Set("Content-type", "text/plain")

    // Parse URL and POST data into the request.Form
    err := request.ParseForm()
    if err != nil {
        http.Error(response, fmt.Sprintf("error parsing url %v", err), 500)
    }

    // Send debug diagnostics to client
    fmt.Fprintf(response, " request.Method     '%v'\n", request.Method)
    fmt.Fprintf(response, " request.RequestURI '%v'\n", request.RequestURI)
    fmt.Fprintf(response, " request.URL.Path   '%v'\n", request.URL.Path)
    fmt.Fprintf(response, " request.Form       '%v'\n", request.Form)
    fmt.Fprintf(response, " request.Cookies()  '%v'\n", request.Cookies())
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.New("form.html").Parse(form)
    t.Execute(w, "")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    log.Println(r.Form)
    templateHandler(w, r)
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("errorHandler...")
        err := f(w, r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            log.Printf("handling %q: %v", r.RequestURI, err)
        }
    }
}
------------------------------------------------------------
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    area() float64
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type Drawer interface {
    draw()
}
func drawShape(d Drawer) {
    d.draw()
}

type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}
func (c Circle) draw() {
    fmt.Println("Circle drawing with radius: ", c.r)
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a * a + b * b)
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}
func (r Rectangle) draw() {
    fmt.Printf("Rectangle drawing with point1: (%f, %f) and point2: (%f, %f)\n",
        r.x1, r.y1, r.x2, r.y2)
}

type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
    var area float64
    for _, shape := range m.shapes {
        area += shape.area()
    }
    return area
}

func main() {
    c := Circle{0, 0, 5}

    c2 := new(Circle)
    c2.x = 0; c2.y = 0; c2.r = 10
    fmt.Println("Circle Area:", totalArea(&c))
    fmt.Println("Circle2 Area:", totalArea(c2))

    r := Rectangle{x1: 0, y1: 0, x2: 5, y2: 5}
    fmt.Println("Rectangle Area:", totalArea(&r))
    fmt.Println("Rectangle + Circle Area:", totalArea(&c, c2, &r))

    m := MultiShape{[]Shape{&r, &c, c2}}
    fmt.Println("Multishape Area:", totalArea(&m))
    fmt.Println("Area Totals:", totalArea(&c, c2, &r))
    fmt.Println("2 X Area Totals:", totalArea(&c, c2, &r, &m))

    drawShape(c)
    drawShape(c2)
    drawShape(r)
}
------------------------------------------------------------
    dashes := strings.Repeat("-", 50)
    fmt.Println(dashes)
------------------------------------------------------------
