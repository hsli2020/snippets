------------------------------------------------------------
func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}
------------------------------------------------------------
type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Todos []Todo

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    json.NewEncoder(w).Encode(todos)
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
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }
------------------------------------------------------------
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
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
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
------------------------------------------------------------
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
------------------------------------------------------------
type person struct {
    name string
    age  int
}

   s := person{name: "Sean", age: 50}
------------------------------------------------------------
package main

import "fmt"

type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
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

import "errors"
import "fmt"

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }

    return arg + 3, nil
}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
------------------------------------------------------------
package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
------------------------------------------------------------
package main

import "fmt"

func main() {
    messages := make(chan string)

    go func() { messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}
------------------------------------------------------------
package main

import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func main() {

    done := make(chan bool, 1)
    go worker(done)

    <-done
}
------------------------------------------------------------
package main

import "fmt"

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
------------------------------------------------------------
package main

import "time"
import "fmt"

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
------------------------------------------------------------
package main

import "time"
import "fmt"

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
------------------------------------------------------------
package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
------------------------------------------------------------
package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}
------------------------------------------------------------
package main

import "fmt"

func main() {
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
------------------------------------------------------------
package main

import "time"
import "fmt"

func main() {

    timer1 := time.NewTimer(time.Second * 2)

    <-timer1.C
    fmt.Println("Timer 1 expired")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
------------------------------------------------------------
package main

import "time"
import "fmt"

func main() {

    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
------------------------------------------------------------
package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 9; a++ {
        <-results
    }
}
------------------------------------------------------------
package main

import "time"
import "fmt"

func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(time.Millisecond * 200)

    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3)

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}
------------------------------------------------------------
package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

    var ops uint64 = 0

    for i := 0; i < 50; i++ {
        go func() {
            for {
                atomic.AddUint64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
------------------------------------------------------------
package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    var state = make(map[int]int)

    var mutex = &sync.Mutex{}

    var ops int64 = 0

    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)

                runtime.Gosched()
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
------------------------------------------------------------
package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    var ops int64 = 0

    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
------------------------------------------------------------
package main

import "fmt"
import "sort"

func main() {

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
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

import "fmt"
import "os"

func main() {

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
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
    fmt.Printf("%v\n", p)
    fmt.Printf("%+v\n", p)
    fmt.Printf("%#v\n", p)
    fmt.Printf("%T\n", p)
    fmt.Printf("%t\n", true)
    fmt.Printf("%d\n", 123)
    fmt.Printf("%b\n", 14)
    fmt.Printf("%c\n", 33)
    fmt.Printf("%x\n", 456)
    fmt.Printf("%f\n", 78.9)
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
    fmt.Printf("%s\n", "\"string\"")
    fmt.Printf("%q\n", "\"string\"")
    fmt.Printf("%x\n", "hex this")
    fmt.Printf("%p\n", &p)
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}

{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
true
123
1110
!
1c8
78.900000
1.234000e+08
1.234000E+08
"string"
"\"string\""
6865782074686973
0x42135100
|    12|   345|
|  1.20|  3.45|
|1.20  |3.45  |
|   foo|     b|
|foo   |b     |
a string
an error
------------------------------------------------------------
package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))
    fmt.Println(r.FindString("peach punch"))
    fmt.Println(r.FindStringIndex("peach punch"))
    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
    fmt.Println(r.FindAllString("peach punch pinch", 2))
    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}

true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH
------------------------------------------------------------
package main

import "encoding/json"
import "fmt"
import "os"

type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    num := dat["num"].(float64)
    fmt.Println(num)

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}

true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{"Page":1,"Fruits":["apple","peach","pear"]}
{"page":1,"fruits":["apple","peach","pear"]}
map[num:6.13 strs:[a b]]
6.13
a
{1 [apple peach]}
apple
{"apple":5,"lettuce":7}
------------------------------------------------------------
package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    now := time.Now()
    p(now)

    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    p(then.Weekday())

    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    diff := now.Sub(then)
    p(diff)

    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    p(then.Add(diff))
    p(then.Add(-diff))
}

2012-10-31 15:50:13.793654 +0000 UTC
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
25891h15m15.142266763s
25891.25420618521
1.5534752523711128e+06
9.320851514226677e+07
93208515142266763
2012-10-31 15:50:13.793654 +0000 UTC
2006-12-05 01:19:43.509120474 +0000 UTC
------------------------------------------------------------
package main

import "fmt"
import "time"

func main() {

    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}

2012-10-31 16:13:58.292387 +0000 UTC
1351700038
1351700038292
1351700038292387000
2012-10-31 16:13:58 +0000 UTC
2012-10-31 16:13:58.292387 +0000 UTC
------------------------------------------------------------
package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339))

    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}

2014-04-15T18:00:15-07:00
2012-11-01 22:08:41 +0000 +0000
6:00PM
Tue Apr 15 18:00:15 2014
2014-04-15T18:00:15.161182-07:00
0000-01-01 20:41:00 +0000 UTC
2014-04-15T18:00:15-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": ...
------------------------------------------------------------
package main

import "time"
import "fmt"
import "math/rand"

func main() {

    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    fmt.Println(rand.Float64())

    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}
------------------------------------------------------------
package main

import "strconv"
import "fmt"

func main() {

    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    _, e := strconv.Atoi("wat")
    fmt.Println(e)
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

    fmt.Println(u.Scheme)

    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}

postgres
user:pass
user
pass
host.com:5432
host.com
5432
/path
f
k=v
map[k:[v]]
v
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
package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    f, err := os.Open("/tmp/dat")
    check(err)

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    _, err = f.Seek(0, 0)
    check(err)

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    f.Close()
}
------------------------------------------------------------
package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    f, err := os.Create("/tmp/dat2")
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()
}
------------------------------------------------------------
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
------------------------------------------------------------
package main

import "flag"
import "fmt"

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

$ go build command-line-flags.go

$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]

$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
------------------------------------------------------------
package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    dateCmd := exec.Command("date")

    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    grepCmd := exec.Command("grep", "hello")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
------------------------------------------------------------
package main

import "syscall"
import "os"
import "os/exec"

func main() {

    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    args := []string{"ls", "-a", "-l", "-h"}

    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
------------------------------------------------------------
package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
------------------------------------------------------------
package main

import(
    "encoding/json"
    "fmt"
    "os"
)

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

func main() {
    //var configs map[string]map[string][]Service
    pathToFile := "/Users/lex/dev/go/samples/src/bitbucket.org/l3x/unmarshal/services-json-encoding.json"

    dev_configs, err := LoadApiServers(pathToFile, "development")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("dev_configs: %v\n\n", dev_configs)

    fmt.Printf("dev_configs[\"speech\"][2].Name: %v\n", dev_configs["speech"][2].Name)
    fmt.Printf("dev_configs[\"sms\"][1].Host: %v\n", dev_configs["sms"][1].Host)
    fmt.Printf("dev_configs[\"payment\"][0].Port: %v\n", dev_configs["sms"][0].Port)
}

Output - encoding/json

dev_configs: map[speech:[{speech-server-1 127.0.0.1 3050 2050 3050} {speech-server-2 127.0.0.1 3051 2051 3051} {speech-server-3 127.0.0.1 3052 2052 3052}] sms:[{sms-server-1 127.0.0.1 4050 0 0} {sms-server-2 127.0.0.1 4051 0 0} {sms-server-3 127.0.0.1 4052 0 0}] payment:[{payment-server-1 127.0.0.1 0 2015 3015}]]

dev_configs["speech"][2].Name: speech-server-3
dev_configs["sms"][1].Host: 127.0.0.1
dev_configs["payment"][0].Port: 4050

Process finished with exit code 0
------------------------------------------------------------
Code Example - jsoncfgo

In this code example we demonstrate how to use the "jsoncfgo" package:

package main

import(
    "sync"
    "encoding/json"
    "log"
    "os"
    "fmt"
    "github.com/l3x/jsoncfgo"
    u "github.com/go-goodies/go_utils"
    "reflect"
)

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

func getService(svcMap interface{}) (*Service) {

    connMap := svcMap.(map[string]interface{})
    connConf := jsoncfgo.Obj(connMap)
    thisSvc := &Service{
        Host: connConf.OptionalString("host", ""),
        Port: connConf.OptionalUint("port", 0),
        RestPort: connConf.OptionalUint("restPort", 0),
        WsPort: connConf.OptionalUint("wsPort", 0),
    }
    return thisSvc
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

func main() {
    pathToFile := "/Users/lex/dev/go/samples/src/bitbucket.org/l3x/unmarshal/services_jsoncfgo.json"

    cfg, err := jsoncfgo.ReadFile(pathToFile)
    if err != nil {
        log.Fatal(err.Error())  // Handle error here
    }

    environmentsList := make(map[string]*Environment)
    environments := cfg.OptionalObject("environments")

    for alias, envMap := range environments {

        servicesMap, ok := envMap.(map[string]interface{})
        if !ok {
            log.Fatalf("entry %q in environments section is a %T, want an object", alias, envMap)
        }
        servicesConf := jsoncfgo.Obj(servicesMap)

        environment := &Environment{
            Speech:   servicesConf["speech"],
            Sms:   servicesConf["sms"],
            Payment:  servicesConf["payment"],
        }
        environmentsList[alias] = environment
        fmt.Printf("environment.Speech: %v\n", environment.Speech)
        speechAry := environment.Speech.([]interface{})
        fmt.Printf("speechAry[0]: %v\n", speechAry[0])
        fmt.Printf("speechAry[1]: %v\n", speechAry[1])
        fmt.Printf("speechAry[2]: %v\n\n", speechAry[2])

        fmt.Printf("environment.Sms: %v\n", environment.Sms)
        smsAry := environment.Sms.([]interface{})
        fmt.Printf("smsAry[0]: %v\n", smsAry[0])
        fmt.Printf("smsAry[1]: %v\n", smsAry[1])
        fmt.Printf("smsAry[2]: %v\n\n", smsAry[2])

        fmt.Printf("environment.Payment: %v\n", environment.Payment)
        paymentAry := environment.Payment.([]interface{})
        fmt.Printf("paymentAry[0]: %v\n", paymentAry[0])
        fmt.Println(u.Dashes(80))
    }

    fmt.Println(u.Dashes(80))
    for alias, envMap := range environmentsList {
        fmt.Printf("alias: %v\n", alias)
        fmt.Printf("envMap: %v\n", envMap)
        printEnvironment(envMap)
        fmt.Println(u.Dashes(80))
    }
}
------------------------------------------------------------
package main

import (
    "fmt"
    "log"
    "errors"
    "net/http"
    "io/ioutil"
    "html/template"
    "regexp"
    "encoding/json"
    "github.com/l3x/jsoncfgo"
    "github.com/go-goodies/go_oops"
)

var Dir string
var Users jsoncfgo.Obj
var AppContext *go_oops.Singleton

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

func printCookies(response http.ResponseWriter, request *http.Request) {
    cookieNameForUsername := AppContext.Data["CookieNameForUsername"].(string)
    fmt.Println("COOKIES:")
    for _, cookie := range request.Cookies() {
        fmt.Printf("%v: %v\n", cookie.Name, cookie.Value)
        if cookie.Name == cookieNameForUsername {
            SetUsernameCookie(response, cookie.Value)
        }
    }; fmt.Println("")
}

func UserHandler(response http.ResponseWriter, request *http.Request){
    response.Header().Set("Content-type", "application/json")
    // json data to send to client
    data := map[string]string { "api" : "user", "name" : "" }
    userApiURL := regexp.MustCompile(`^/user/(\w+)$`)
    usernameMatches := userApiURL.FindStringSubmatch(request.URL.Path)
    // regex matches example: ["/user/joesample", "joesample"]
    if len(usernameMatches) > 0 {
        printCookies(response, request)
        var userName string
        userName = usernameMatches[1]  // ex: joesample
        userObj := AppContext.Data[userName]
        fmt.Printf("userObj: %v\n", userObj)
        if userObj == nil {
            msg := fmt.Sprintf("Invalid username (%s)", userName)
            panic(errors.New(msg))
        } else {
            // Send JSON to the client
            thisUser := userObj.(jsoncfgo.Obj)
            fmt.Printf("thisUser: %v\n", thisUser)
            data["name"] = thisUser["firstname"].(string) + " " + thisUser["lastname"].(string)
        }
        json_bytes, _ := json.Marshal(data)
        fmt.Printf("json_bytes: %s\n", string(json_bytes[:]))
        fmt.Fprintf(response, "%s\n", json_bytes)

    } else {
        http.Error(response, "404 page not found", 404)
    }
}

func SetUsernameCookie(response http.ResponseWriter, userName string){
    // Add cookie to response
    cookieName := AppContext.Data["CookieNameForUsername"].(string)
    cookie := http.Cookie{Name: cookieName, Value: userName}
    http.SetCookie(response, &cookie)
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
    fmt.Fprintf(response, "<table>")
    fmt.Fprintf(response, "<tr><td><strong>request.Method    </strong></td><td>'%v'</td></tr>", request.Method)
    fmt.Fprintf(response, "<tr><td><strong>request.RequestURI</strong></td><td>'%v'</td></tr>", request.RequestURI)
    fmt.Fprintf(response, "<tr><td><strong>request.URL.Path  </strong></td><td>'%v'</td></tr>", request.URL.Path)
    fmt.Fprintf(response, "<tr><td><strong>request.Form      </strong></td><td>'%v'</td></tr>", request.Form)
    fmt.Fprintf(response, "<tr><td><strong>request.Cookies() </strong></td><td>'%v'</td></tr>", request.Cookies())
    fmt.Fprintf(response, "</table>")

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

var form = `
<h1>Debug Info (POST form)</h1>
<form method="POST" action="" name="frmTest">
<div>
    <label for="username">User Name</label>
    <input id="username" name="username" placeholder="joesample, alicesmith, or bobbrown" required="" type="text"
size="50">
</div>
<div><input type="submit" value="Submit"></div>
</form>

</form>
`

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

func doThis() error { return nil }
func doThat() error { return errors.New("ERROR - doThat") }

func wrappedHandler(w http.ResponseWriter, r *http.Request) error {
    log.Println("betterHandler...")
    if err := doThis(); err != nil {
        return fmt.Errorf("doing this: %v", err)
    }

    if err := doThat(); err != nil {
        return fmt.Errorf("doing that: %v", err)
    }
    return nil
}

func main() {
    cfg := jsoncfgo.Load("/Users/lex/dev/go/data/webserver/webserver-config.json")

    host := cfg.OptionalString("host", "localhost")
    fmt.Printf("host: %v\n", host)

    port := cfg.OptionalInt("port", 8080)
    fmt.Printf("port: %v\n", port)

    Dir = cfg.OptionalString("dir", "www/")
    fmt.Printf("web_dir: %v\n", Dir)

    redirect_code := cfg.OptionalInt("redirect_code", 307)
    fmt.Printf("redirect_code: %v\n\n", redirect_code)

    mux := http.NewServeMux()

    fileServer := http.Dir(Dir)
    fileHandler := http.FileServer(fileServer)
    mux.Handle("/", fileHandler)

    rdh := http.RedirectHandler("http://example.org", redirect_code)
    mux.Handle("/redirect", rdh)
    mux.Handle("/notFound", http.NotFoundHandler())

    mux.Handle("/help", http.HandlerFunc( HelpHandler ))

    mux.Handle("/debugForm", http.HandlerFunc( DebugFormHandler ))
    mux.Handle("/debugQuery", http.HandlerFunc( DebugQueryHandler ))

    mux.Handle("/user/", http.HandlerFunc( UserHandler ))
    mux.Handle("/ajax", http.HandlerFunc( AjaxHandler ))

    mux.Handle("/adapter", errorHandler(wrappedHandler))

    log.Printf("Running on port %d\n", port)

    addr := fmt.Sprintf("%s:%d", host, port)

    Users := jsoncfgo.Load("/Users/lex/dev/go/data/webserver/users.json")

    joesample := Users.OptionalObject("joesample")
    fmt.Printf("joesample: %v\n", joesample)
    fmt.Printf("joesample['firstname']: %v\n", joesample["firstname"])
    fmt.Printf("joesample['lastname']: %v\n\n", joesample["lastname"])

    alicesmith := Users.OptionalObject("alicesmith")
    fmt.Printf("alicesmith: %v\n", alicesmith)
    fmt.Printf("alicesmith['firstname']: %v\n", alicesmith["firstname"])
    fmt.Printf("alicesmith['lastname']: %v\n\n", alicesmith["lastname"])

    bobbrown := Users.OptionalObject("bobbrown")
    fmt.Printf("bobbrown: %v\n", bobbrown)
    fmt.Printf("bobbrown['firstname']: %v\n", bobbrown["firstname"])
    fmt.Printf("bobbrown['lastname']: %v\n\n", bobbrown["lastname"])

    AppContext = go_oops.NewSingleton()
    AppContext.Data["CookieNameForUsername"] = "testapp-username"
    AppContext.Data["joesample"] = joesample
    AppContext.Data["alicesmith"] = alicesmith
    AppContext.Data["bobbrown"] = bobbrown
    fmt.Printf("AppContext: %v\n", AppContext)
    fmt.Printf("AppContext.Data[\"joesample\"]: %v\n", AppContext.Data["joesample"])
    fmt.Printf("AppContext.Data[\"alicesmith\"]: %v\n", AppContext.Data["alicesmith"])
    fmt.Printf("AppContext.Data[\"bobbrown\"]: %v\n\n", AppContext.Data["bobbrown"])

    err := http.ListenAndServe(addr, mux)
    fmt.Println(err.Error())
}
------------------------------------------------------------
package main    // Executable commands must always use package main.

import (
    "fmt"       // fmt.Println formats output to console
    "math"      // provides math.Sqrt function
)

// ----------------------
//    Shape interface
// ----------------------
// Shape interface defines a method set (consisting of the area method)
type Shape interface {
    area() float64          // any type that implements an area method is considered a Shape
}
// Calculate total area of all shapes via polymorphism (all shapes implement the area method)
func totalArea(shapes ...Shape) float64 {   // Use interface type as as function argument
    var area float64                        // "..." makes shapes "variadic" (can send one or more)
    for _, s := range shapes {
        area += s.area()    // the current Shape implements/receives the area method
    }                       // go passes the pointer to the shape to the area method
    return area
}

// ----------------------
//    Drawer interface
// ----------------------
type Drawer interface {
    draw()                  // does not return a type
}
func drawShape(d Drawer) {  // associate this method with the Drawer interface
    d.draw()
}

// ----------------------
//      Circle Type
// ----------------------
type Circle struct {        // Since "Circle" is capitalized, it is visible outside this package
    x, y, r float64         // a Circle struct is a collection of fields: x, y, r
}
// Circle implements Shape interface b/c it has an area method
// area is a method, which is special type of function that is associated with the Circle struct
// The Circle struct becomes the "receiver" of this method, so we can use the "." operator
func (c *Circle) area() float64 {   // dereference Circle type (data pointed to by c)
    return math.Pi * c.r * c.r      // Pi is a constant in the math package
}
func (c Circle) draw() {
    fmt.Println("Circle drawing with radius: ", c.r)    // encapsulated draw implementation for Circle type
}
// ----------------------
//     Rectangle Type
// ----------------------
type Rectangle struct {     // a struct contains named fields of data
    x1, y1, x2, y2 float64  // define multiple fields with same data type on one line
}
func distance(x1, y1, x2, y2 float64) float64 {         // lowercase functin name visible only in this package
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a * a + b * b)
}
// Rectangle implements Shape interface b/c it has an area method
func (r *Rectangle) area() float64 {       // "r" is passed by reference
    l := distance(r.x1, r.y1, r.x1, r.y2)  // define and assign local variable "l"
    w := distance(r.x1, r.y1, r.x2, r.y1)  // l and w only available within scope of area function
    return l * w
}
func (r Rectangle) draw() {                // "r" is passed by value
    fmt.Printf("Rectangle drawing with point1: (%f, %f) and point2: (%f, %f)\n", r.x1, r.y1, r.x2, r.y2)
}
// ----------------------
//    MultiShape Type
// ----------------------
type MultiShape struct {
    shapes []Shape  // shapes field is a slice of interfaces
}
//
func (m *MultiShape) area() float64 {
    var area float64
    for _, shape := range m.shapes {    // iterate through shapes ("_" indicates that index is not used)
        area += shape.area()            // execute polymorphic area method for this shape
    }
    return area
}

func main() {
    c := Circle{0, 0, 5}                                            // initialize new instance of Circle type by field order "struct literal"
                                                                    // The new function allocates memory for all  fields, sets each to their zero value and returns a pointer
    c2 := new(Circle)                                               // c2 is a pointer to the instantiated Circle type
    c2.x = 0; c2.y = 0; c2.r = 10                                   // initialize data with multiple statements on one line
    fmt.Println("Circle Area:", totalArea(&c))                      // pass address of circle (c)
    fmt.Println("Circle2 Area:", totalArea(c2))                     // c2 was defined using built-in new function
    r := Rectangle{x1: 0, y1: 0, x2: 5, y2: 5}                      // "struct literal" rectangle (r) initialized by field name
    fmt.Println("Rectangle Area:", totalArea(&r))                   // pass address of rectangle (r)
    fmt.Println("Rectangle + Circle Area:", totalArea(&c, c2, &r))  // can pass multiple shapes
    m := MultiShape{[]Shape{&r, &c, c2}}                            // pass slice of shapes
    fmt.Println("Multishape Area:", totalArea(&m))                  // calculate total area of all shapes
    fmt.Println("Area Totals:", totalArea(&c, c2, &r))              // c2 is a pointer to a circle, &c and &r are addresses of shapes
    fmt.Println("2 X Area Totals:", totalArea(&c, c2, &r, &m))      // twice the size of all areas
    drawShape(c)                                                    // execute polymorphic method call
    drawShape(c2)
    drawShape(r)
}
------------------------------------------------------------
package main

import (
    "fmt"
    "strings"
)

type Car struct {
    Make  string
    Model  string
    Options []string
}

func main() {

    dashes := strings.Repeat("-", 50)

    is250 := &Car{"Lexus", "IS250", []string{"GPS", "Alloy Wheels", "Roof Rack", "Power Outlets", "Heated Seats"}}
    accord := &Car{"Honda", "Accord", []string{"Alloy Wheels", "Roof Rack"}}
    blazer := &Car{"Chevy", "Blazer", []string{"GPS", "Roof Rack", "Power Outlets"}}

    cars := []*Car{is250, accord, blazer}
    fmt.Printf("Cars:\n%v\n\n", cars)  // cars is a slice of pointers to our three cars

    // Create a map to associate options with each car
    car_options := make(map[string][]*Car)

    fmt.Printf("CARS:\n%s\n", dashes)
    for _, car := range cars {
        fmt.Printf("%v\n", car)
        for _, option := range car.Options {
            // Associate this car with each of it's options
            car_options[option] = append(car_options[option], car)
            fmt.Printf("car_options[option]: %s\n", option)
        }
        fmt.Println(dashes)
    }
    fmt.Println(dashes)

    // Print a list of cars with the "GPS" option
    for _, p := range car_options["GPS"] {
        fmt.Println(p.Make, "has GPS.")
    }

    fmt.Println("")
    fmt.Println(len(car_options["Alloy Wheels"]), "has Alloy Wheels.")
}
------------------------------------------------------------
package main

import (
    "fmt"
    "errors"
    "strings"
)

type Value struct {
    Name string
    MilesAway int
}

type Node struct {
    Value               // Embedded struct
    next, prev  *Node
}

type List struct {
    head, tail *Node
}

func (l *List) First() *Node {
    return l.head
}

func (n *Node) Next() *Node {
    return n.next
}

func (n *Node) Prev() *Node {
    return n.prev
}

// Create new node with value
func (l *List) Push(v Value) *List {
    n := &Node{Value: v}
    if l.head == nil {
        l.head = n      // First node
    } else {
        l.tail.next = n // Add after prev last node
        n.prev = l.tail // Link back to prev last node
    }
    l.tail = n          // reset tail to newly added node
    return l
}

func (l *List) Find(name string) *Node {
    found := false
    var ret *Node = nil
    for n := l.First(); n != nil && !found; n = n.Next() {
        if n.Value.Name == name {
            found = true
            ret = n
        }
    }
    return ret
}

func (l *List) Delete(name string) bool {
    success := false
    node2del := l.Find(name)
    if node2del != nil {
        fmt.Println("Delete - FOUND: ", name)
        prev_node := node2del.prev
        next_node := node2del.next
        // Remove this node
        prev_node.next = node2del.next
        next_node.prev = node2del.prev
        success = true
    }
    return success
}

var errEmpty = errors.New("ERROR - List is empty")

// Pop last item from list
func (l *List) Pop() (v Value, err error) {
    if l.tail == nil {
        err = errEmpty
    } else {
        v = l.tail.Value
        l.tail = l.tail.prev
        if l.tail == nil {
            l.head = nil
        }
    }
    return v, err
}

func main() {
    dashes := strings.Repeat("-", 50)
    l := new(List)  // Create Doubly Linked List

    l.Push(Value{Name: "Atlanta", MilesAway: 0})
    l.Push(Value{Name: "Las Vegas", MilesAway: 1961})
    l.Push(Value{Name: "New York", MilesAway: 881})

    processed := make(map[*Node]bool)

    fmt.Println("First time through list...")
    for n := l.First(); n != nil; n = n.Next() {
        fmt.Printf("%v\n", n.Value)
        if processed[n] {
            fmt.Printf("%s as been processed\n", n.Value)
        }
        processed[n] = true
    }

    fmt.Println(dashes)
    fmt.Println("Second time through list...")
    for n := l.First(); n != nil; n = n.Next() {
        fmt.Printf("%v", n.Value)
        if processed[n] {
            fmt.Println(" has been processed")
        } else { fmt.Println() }
        processed[n] = true
    }

    fmt.Println(dashes)

    var found_node *Node
    city_to_find := "New York"
    found_node = l.Find(city_to_find)
    if found_node == nil {
        fmt.Printf("NOT FOUND: %v\n", city_to_find)
    } else {
        fmt.Printf("FOUND: %v\n", city_to_find)
    }

    city_to_find = "Chicago"
    found_node = l.Find(city_to_find)
    if found_node == nil {
        fmt.Printf("NOT FOUND: %v\n", city_to_find)
    } else {
        fmt.Printf("FOUND: %v\n", city_to_find)
    }

    fmt.Println(dashes)
    city_to_remove := "Las Vegas"
    successfully_removed_city := l.Delete(city_to_remove)
    if successfully_removed_city {
        fmt.Printf("REMOVED: %v\n", city_to_remove)
    } else {
        fmt.Printf("DID NOT REMOVE: %v\n", city_to_remove)
    }

    city_to_remove = "Chicago"
    successfully_removed_city = l.Delete(city_to_remove)
    if successfully_removed_city {
        fmt.Printf("REMOVED: %v\n", city_to_remove)
    } else {
        fmt.Printf("DID NOT REMOVE: %v\n", city_to_remove)
    }

    fmt.Println(dashes)
    fmt.Println("* Pop each value off list...")
    for v, err := l.Pop(); err == nil; v, err = l.Pop() {
        fmt.Printf("%v\n", v)
    }
    fmt.Println(l.Pop())  // Generate error - attempt to pop from empty list
}
------------------------------------------------------------
package main

import (
    "fmt"
    "log"
    "strings"
    "errors"
    u "github.com/go-goodies/go_utils"
)
// enums indicating number of characters for that type of word
// ex: a TINY word has 4 or fewer characters
const (
    TEENINY WordSize = 1
    SMALL   WordSize = 4 << iota
    MEDIUM                      // assigned 8 from iota
    LARGE                       // assigned 16 from iota
    XLARGE  WordSize = 32000
)

type WordSize int

func (ws WordSize) String() string {
    var s string
    if ws&TEENINY == TEENINY {
        s = "TEENINY"
    }
    return s
}

// ChainLink allows us to chain function/method calls.  It also keeps  
// data internal to ChainLink, avoiding the side effect of mutated data.
type ChainLink struct {
    Data []string
}

func (v *ChainLink)Value() []string {
    return v.Data
}

// stringFunc is a first-class method, used as a parameter to _map
type stringFunc func(s string) (result string)

// _map uses stringFunc to modify (up-case) each string in the slice
func (v *ChainLink)_map(fn stringFunc) *ChainLink {
    var mapped []string
    orig := *v
    for _, s := range orig.Data {
        mapped = append(mapped, fn(s))  // first-class function
    }
    v.Data = mapped
    return v
}

// _filter uses embedded logic to filter the slice of strings
// Note: We could have chosen to use a first-class function
func (v *ChainLink)_filter(max WordSize) *ChainLink {
    filtered := []string{}
    orig := *v
    for _, s := range orig.Data {
        if len(s) <= int(max) {             // embedded logic
            filtered = append(filtered, s)
        }
    }
    v.Data = filtered
    return v
}


func main() {
    nums := []string{
        "tiny",
        "marathon",
        "philanthropinist",
        "supercalifragilisticexpialidocious"}

    data := ChainLink{nums};
    orig_data := data.Value()
    fmt.Printf("unfiltered: %#v\n", data.Value())

    filtered := data._filter(MEDIUM)
    fmt.Printf("filtered: %#v\n", filtered)

    fmt.Printf("filtered and mapped (MEDIUM sized words): %#v\n",
        filtered._map(strings.ToUpper).Value())

    data = ChainLink{nums}
    fmt.Printf("filtered and mapped (MEDIUM sized words): %#v\n",
        data._filter(MEDIUM)._map(strings.ToUpper).Value())

    data = ChainLink{nums}
    fmt.Printf("filtered twice and mapped (SMALL sized words): %#v\n",
        data._filter(MEDIUM)._map(strings.ToUpper)._filter(SMALL).Value())

    data = ChainLink{nums}
    val := data._map(strings.ToUpper)._filter(XLARGE).Value()
    fmt.Printf("mapped and filtered (XLARGE sized words): %#v\n", val)

    // heredoc with interpoloation
    constants := `
** Constants ***
SMALL: %d
MEDIUM: %d
LARGE: %d
XLARGE: %d
`
    fmt.Printf(constants, SMALL, MEDIUM, LARGE, XLARGE)
    fmt.Printf("TEENINY: %s\n\n", TEENINY)

    fmt.Printf("Join(nums, \"|\")     : %v\n", u.Join(nums, "|"))
    fmt.Printf("Join(orig_data, \"|\"): %v\n", u.Join(orig_data, "|"))
    fmt.Printf("Join(data, \"|\")     : %v\n\n", u.Join(data.Value(), "|"))

    if u.Join(nums, "|") == u.Join(orig_data, "|") {
        fmt.Println("No Side Effects!")
    } else {
        log.Print(errors.New("WARNING - Side Effects!"))

    }
}
------------------------------------------------------------

