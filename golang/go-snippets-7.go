// ---------------------------------------------------------
type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

// var c Circle
// c := new(Circle)
// c := Circle{x: 0, y: 0, r: 5}
// c := Circle{0, 0, 5}
// fmt.Println(c.area())

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

type MultiShape struct {
    shapes []Shape
}

func (m *MultiShape) area() float64 {
    var area float64
    for _, s := range m.shapes {
        area += s.area()
    }
    return area 
}

Now a MultiShape can contain Circles, Rectangles or even other MultiShapes.

type Person struct {
    Name string
}

type Student struct {
    Person    // 继承|派生
    Birth string
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

func (p *Student) Talk() {  // 改写|overwrite
    fmt.Println("Hi, I'm a student, my name is", p.Name)
}

p := Person{Name:"tom"}
p.Talk()

s := new(Student)
s.Name = "sam"
s.Birth = "2000-02-20"
s.Talk()
// ---------------------------------------------------------
// x, err := f()
// x, ok  := f()

import m "golang-book/chapter11/math"
'm' is the alias

val, ok := rating["key"]
// check whether there is a key in map

// func TestAverage(t *testing.T) 
The 'go test' command will look for any tests in any of the files in the current folder and run them. Tests are identified by starting a function with the word 'Test' and taking one argument of type '*testing.T'. 

// arr := []byte("test")
// str := string([]byte{'t','e','s','t'})
Sometimes we need to work with strings as binary data. 
To convert a string to a slice of bytes (and vice- versa) do this:

// return error
return -1, errors.New("Can't work with 42")

// 类型转换 (isinstanceof)
if ae, ok := e.(*argError); ok {

// 这两种写法的区别，用指针表示类的状态会被改变，非指针则类的状态不变
// 类似C++中在成员函数后加上const的作用
func (this *Foo) Method() string { }
func (this   Foo) Method() string { }

// arr := [ 1, 2, 3, 4 ] // WRONG
// arr := { 1, 2, 3, 4 } // WRONG
// arr := []int{ 1, 2, 3, 4 } // OK

// h := md5.New()
// fmt.Printf("%x", h.Sum([]byte("134")))

// h := md5.New()
// io.WriteString(h, "The fog is getting thicker!")
// io.WriteString(h, "And Leon's getting laaarger!")
// fmt.Printf("%x", h.Sum(nil))

// f, _ := os.Open(filename)
// defer f.Close()

func sum(x []float64) float64 {
    var total float64 = 0.0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    return total;
}

func test_sum() {
    x := []float64{ 98, 93, 77, 82, 83 }
    s := sum(x)
    fmt.Println(s)
}

func sum(x []float64) float64 {
    var total float64 = 0.0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    return total;
}

func test_sum() {
    x := []float64{ 98, 93, 77, 82, 83 }
    s := sum(x)
    fmt.Println(s)
}

多个返回值
func ret2() (int, float64) {
     return 5, 6.6
}

func test_ret2() {
     x, y := ret2()
     fmt.Println(x, y);
}

可变参数
// func Println(a ...interface{}) (n int, err error)

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func test_add() {
     fmt.Println(add(1,2,3))
     xs := []int{1,2,3}
     fmt.Println(add(xs...))
}

Go没有static关键字，但可通过closure达到同样效果
func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
       ret = i
       i += 2
       return 
    }
}

func test_evenGen() {
    nextEven := makeEvenGenerator()
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
}

递归
func factorial(x uint64) uint64 {
    if x == 0 {
        return 1 
    }
    return x * factorial(x-1)
}

Go没有Try/Catch/Exception，可通过panic/recover达到类似效果
func test_panic() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}

func test_closure() {
    x := 0
    increment := func() int {
        x++
        return x 
    }
    fmt.Println(increment())
    fmt.Println(increment())
}

strings
fmt.Println(
    strings.Contains("test", "es"),    // true
    strings.Count("test", "t"),        // 2
    strings.HasPrefix("test", "te"),   // true
    strings.HasSuffix("test", "st"),   // true
    strings.Index("test", "e"),        // 1
    strings.Join([]string{"a","b"}, "-"),    // "a-b"
    strings.Repeat("a", 5),                  // == "aaaaa"
    strings.Replace("aaaa", "a", "b", 2),    // "bbaa"
    strings.Split("a-b-c-d-e", "-"),         // []string{"a","b","c","d","e"}
    strings.ToLower("TEST"),    // "test"
    strings.ToUpper("test"),    // "TEST"
)
// ---------------------------------------------------------
package main

import (
    "fmt"
    "time"
    "math/rand"
)

func f(n int) {
    for i := 0; i < 10; i++ {
        fmt.Println(n, ":", i)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(time.Millisecond * amt)
    } 
}

func main() {
    for i := 0; i < 10; i++ {
        go f(i)
    }
    var input string
    fmt.Scanln(&input)
}
// ---------------------------------------------------------
package main

// 这两种写法的区别，用指针表示类的状态会被改变，非指针则类的状态不变
// 类似C++中在成员函数后加上const的作用
func (this *Foo) Method() string { }
func (this   Foo) Method() string { }

import "fmt"

type Foo struct {
    value int
}

func (this *Foo) Set(val int) {  // try to remove the star *
    this.value = val
}

func (this Foo) Get() int {
    return this.value
}

func main() {
    v := new(Foo)
    fmt.Println("Foo.value=", v.Get());
    //fmt.Println(v);

    v.Set(123);
    fmt.Println("Foo.value=", v.Get());
    //fmt.Println(v);
}

import "fmt"
import "crypto/md5"
import "encoding/hex"

func main() {
    println("Let's Go!");

    hash := md5.Sum([]byte("1"));

    // output: method 1
    str := hex.EncodeToString(hash[:]);
    fmt.Println(str);

    // output: method 2
    fmt.Printf("%x\n", hash);
}
// ---------------------------------------------------------
GoLang builtin funcs

func copy(dst, src []Type) int
func append(slice []Type, elems ...Type) []Type
func delete(m map[Type]Type1, key Type)

func len(v Type) int
func cap(v Type) int

func new(Type) *Type
func make(Type, size IntegerType) Type
func close(c chan<- Type)

func panic(v interface{})
func recover() interface{}

func print(args ...Type)
func println(args ...Type)

func real(c ComplexType) FloatType
func imag(c ComplexType) FloatType
func complex(r, i FloatType) ComplexType

var ErrTooLarge = errors.New("bytes.Buffer: too large")
// ---------------------------------------------------------
////////////////////////////////////////////////////////////
// read file
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()

    // get the file size
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // read the file
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}

////////////////////////////////////////////////////////////
// read file
package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    bs, err := ioutil.ReadFile("test.txt")
    if err != nil {
        return
    }
    str := string(bs)
    fmt.Println(str)
}

////////////////////////////////////////////////////////////
package main

import (
    "os"
)

func main() {
    file, err := os.Create("test.txt")
    if err != nil {
        // handle the error here
        return
    }
    defer file.Close()
    file.WriteString("test")
}

////////////////////////////////////////////////////////////
package main

import (
    "fmt"
    "os"
)

func main() {
    dir, err := os.Open(".")
    if err != nil {
        return
    }
    defer dir.Close()

    fileInfos, err := dir.Readdir(-1)
    if err != nil {
        return
    }
    for _, fi := range fileInfos {
        fmt.Println(fi.Name())
    }
}

////////////////////////////////////////////////////////////
package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
        fmt.Println(path)
        return nil
    })
}

////////////////////////////////////////////////////////////
package main

import "errors"

func main() {
    err := errors.New("error message")
}

////////////////////////////////////////////////////////////
package main

import ("fmt" ; "container/list")

func main() {
    var x list.List
    x.PushBack(1)
    x.PushBack(2)
    x.PushBack(3)

    for e := x.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(int))
    }
}

////////////////////////////////////////////////////////////
package main

import ("fmt" ; "sort")

type Person struct {
    Name string
    Age int
}

type ByName []Person

func (this ByName) Len() int {
    return len(this)
}
func (this ByName) Less(i, j int) bool {
    return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}

func main() {
    kids := []Person{
        {"Jill", 9},
        {"Jack", 10},
    }
    sort.Sort(ByName(kids))
    fmt.Println(kids)
}

type ByAge []Person

func (this ByAge) Len() int {
    return len(this)
}
func (this ByAge) Less(i, j int) bool {
    return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
}
// ---------------------------------------------------------
type Duration int64

const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)

func (d Duration) Hours() float64
func (d Duration) Minutes() float64
func (d Duration) Nanoseconds() int64
func (d Duration) Seconds() float64

func (d Duration) String() string

type Month int

const (
    January Month = 1 + iota
    February
    March
    April
    May
    June
    July
    August
    September
    October
    November
    December
)

func (m Month) String() string
    String returns the English name of the month ("January", "February", ...).

type Weekday int

const (
    Sunday Weekday = iota // (Sunday = 0, ...).
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (d Weekday) String() string // String returns the English name of the day ("Sunday", "Monday", ...).

func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
func Now() Time

func Parse(layout, value string) (Time, error)
func ParseInLocation(layout, value string, loc *Location) (Time, error)
func Unix(sec int64, nsec int64) Time

func (t Time) Add(d Duration) Time
func (t Time) AddDate(years int, months int, days int) Time
func (t Time) After(u Time) bool
func (t Time) Before(u Time) bool
func (t Time) Clock() (hour, min, sec int)
func (t Time) Date() (year int, month Month, day int)
func (t Time) Day() int
func (t Time) Equal(u Time) bool
func (t Time) Format(layout string) string
func (t *Time) GobDecode(data []byte) error
func (t Time) GobEncode() ([]byte, error)
func (t Time) Hour() int
func (t Time) ISOWeek() (year, week int)
func (t Time) In(loc *Location) Time
func (t Time) IsZero() bool
func (t Time) Local() Time
func (t Time) Location() *Location
func (t Time) MarshalBinary() ([]byte, error)
func (t Time) MarshalJSON() ([]byte, error)
func (t Time) MarshalText() ([]byte, error)
func (t Time) Minute() int
func (t Time) Month() Month
func (t Time) Nanosecond() int
func (t Time) Round(d Duration) Time
func (t Time) Second() int
func (t Time) String() string
func (t Time) Sub(u Time) Duration
func (t Time) Truncate(d Duration) Time
func (t Time) UTC() Time
func (t Time) Unix() int64
func (t Time) UnixNano() int64
func (t *Time) UnmarshalBinary(data []byte) error
func (t *Time) UnmarshalJSON(data []byte) (err error)
func (t *Time) UnmarshalText(data []byte) (err error)
func (t Time) Weekday() Weekday
func (t Time) Year() int
func (t Time) YearDay() int
func (t Time) Zone() (name string, offset int)

package unicode

func In(r rune, ranges ...*RangeTable) bool
func Is(rangeTab *RangeTable, r rune) bool
func IsOneOf(ranges []*RangeTable, r rune) bool

func IsControl(r rune) bool
func IsDigit(r rune) bool
func IsGraphic(r rune) bool
func IsLetter(r rune) bool
func IsUpper(r rune) bool
func IsLower(r rune) bool
func IsMark(r rune) bool
func IsNumber(r rune) bool
func IsPrint(r rune) bool
func IsPunct(r rune) bool
func IsSymbol(r rune) bool
func IsTitle(r rune) bool
func ToLower(r rune) rune
func ToTitle(r rune) rune
func ToUpper(r rune) rune
func IsSpace(r rune) bool

func SimpleFold(r rune) rune
func To(_case int, r rune) rune

package unsafe

func Alignof(v ArbitraryType) uintptr
func Offsetof(v ArbitraryType) uintptr
func Sizeof(v ArbitraryType) uintptr
// ---------------------------------------------------------
godoc builtin // 列出这个包所有函数说明，这个包含有很多很有用的函数
godoc net/http
godoc fmt Println
godoc -src fmt Println

// 从浏览器查看go文档
godoc -http=:8080   // 启动服务
http://localhost:8080 // 查看
// ---------------------------------------------------------
GoLang bytes package funcs

package bytes

func Compare(a, b []byte) int
func Contains(b, subslice []byte) bool
func Count(s, sep []byte) int
func Equal(a, b []byte) bool
func EqualFold(s, t []byte) bool

func Fields(s []byte) [][]byte
func FieldsFunc(s []byte, f func(rune) bool) [][]byte

func HasPrefix(s, prefix []byte) bool
func HasSuffix(s, suffix []byte) bool

func Index(s, sep []byte) int
func IndexAny(s []byte, chars string) int
func IndexByte(s []byte, c byte) int
func IndexFunc(s []byte, f func(r rune) bool) int
func IndexRune(s []byte, r rune) int

func Join(s [][]byte, sep []byte) []byte

func LastIndex(s, sep []byte) int
func LastIndexAny(s []byte, chars string) int
func LastIndexFunc(s []byte, f func(r rune) bool) int

func Map(mapping func(r rune) rune, s []byte) []byte
func Repeat(b []byte, count int) []byte
func Replace(s, old, new []byte, n int) []byte

func Runes(s []byte) []rune

func Split(s, sep []byte) [][]byte
func SplitAfter(s, sep []byte) [][]byte
func SplitAfterN(s, sep []byte, n int) [][]byte
func SplitN(s, sep []byte, n int) [][]byte

func Title(s []byte) []byte

func ToLower(s []byte) []byte
func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte

func ToTitle(s []byte) []byte
func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte

func ToUpper(s []byte) []byte
func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte

func Trim(s []byte, cutset string) []byte
func TrimFunc(s []byte, f func(r rune) bool) []byte

func TrimLeft(s []byte, cutset string) []byte
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte

func TrimPrefix(s, prefix []byte) []byte
func TrimSuffix(s, suffix []byte) []byte

func TrimRight(s []byte, cutset string) []byte
func TrimRightFunc(s []byte, f func(r rune) bool) []byte

func TrimSpace(s []byte) []byte
// ---------------------------------------------------------
package list

	for e := l.Front(); e != nil; e = e.Next() {
		// do something with e.Value
	}

type Element struct {
    // The value stored with this element.
    Value interface{}
    // contains filtered or unexported fields
}

func (e *Element) Next() *Element
func (e *Element) Prev() *Element

type List struct {
    // contains filtered or unexported fields
}

func New() *List

func (l *List) Back() *Element
func (l *List) Front() *Element

func (l *List) Init() *List

func (l *List) InsertAfter(v interface{}, mark *Element) *Element
func (l *List) InsertBefore(v interface{}, mark *Element) *Element

func (l *List) Len() int

func (l *List) MoveAfter(e, mark *Element)
func (l *List) MoveBefore(e, mark *Element)
func (l *List) MoveToBack(e *Element)
func (l *List) MoveToFront(e *Element)

func (l *List) PushBack(v interface{}) *Element
func (l *List) PushBackList(other *List)
func (l *List) PushFront(v interface{}) *Element
func (l *List) PushFrontList(other *List)

func (l *List) Remove(e *Element) interface{}

package ring

type Ring struct {
    Value interface{} // for use by client; untouched by this library
    // contains filtered or unexported fields
}

func New(n int) *Ring

func (r *Ring) Do(f func(interface{}))
func (r *Ring) Len() int
func (r *Ring) Link(s *Ring) *Ring
func (r *Ring) Move(n int) *Ring
func (r *Ring) Next() *Ring
func (r *Ring) Prev() *Ring
func (r *Ring) Unlink(n int) *Ring

package hex

func Decode(dst, src []byte) (int, error)
func DecodeString(s string) ([]byte, error)
func DecodedLen(x int) int

func Dump(data []byte) string
func Dumper(w io.Writer) io.WriteCloser

func Encode(dst, src []byte) int
func EncodeToString(src []byte) string
func EncodedLen(n int) int
// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------

// ---------------------------------------------------------
