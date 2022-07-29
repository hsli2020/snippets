// ---------------------------------------------------------
Simple file server

package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./")))
    http.ListenAndServe(":8123", nil)
}
// ---------------------------------------------------------
package main

import "fmt"

type T int

func Start() T {
        fmt.Println("Start")
        return T(0)
}

func (t T) End() {
        fmt.Println("End")
}

func main() {
        defer Start().End()
        fmt.Println("During")
}
// ---------------------------------------------------------
func Sqrt(x float64) float64 {
        z := 0.0
        for i := 0; i < 1000; i++ {
            z -= (z*z - x) / (2 * x)
        }
        return z
}

func palindrome(num int) bool {
  var n, reverse, dig int
  n = num
  reverse = 0
  for (num > 0){
    dig = num % 10
    reverse = reverse * 10 + dig
    num = num / 10
  }
  return n == reverse
}

func get_size(file_bytes uint64) string {
    var (
        units []string
        size  string
        i     int
    )
    units = []string{"B", "K", "M", "G", "T", "P"}
    i = 0
    for {
        i++
        file_bytes = file_bytes / 1024
        if file_bytes < 1024 {
            size = fmt.Sprintf("%d", file_bytes) + units[i]
            break
        }
    }
    return size
}

func get_size(file_bytes uint64) (size string) { 
    var size  string
    if file_bytes<=0 { 
      return 
    } 
    units := []string{"B", "K", "M", "G", "T", "P"} 
    i := 0 
    for { 
        i++ 
        file_bytes = file_bytes / 1024 
        if file_bytes < 1024 { 
            size = fmt.Sprintf("%d%s", file_bytes,units[i]) 
            break 
        } 
    } 
    return 
} 
// ---------------------------------------------------------
shuffle.go:

package main

import (
        "flag"
        "fmt"
        "sort"
        "rand"
        "time"
)

// Types that satisfy this interface can be shuffled.  Note that this is a
// subset of sort.Interface.
type Interface interface {
        Len() int
        Swap(i, j int)
}

func Shuffle(data Interface) {
        for i := data.Len() - 1; i > 0; i-- {
                if j := rand.Intn(i + 1); i != j {
                        data.Swap(i, j)
                }
        }
}

var n *int = flag.Int("n", 10, "shuffle array size")

func main() {
        flag.Parse()
        ai := make([]int, *n)
        af := make([]float, *n)
        for i := 0; i < *n; i++ {
                ai[i] = i
                af[i] = float(i) / 10.0
        }
        rand.Seed(time.Nanoseconds())
        Shuffle(sort.IntArray(ai))
        Shuffle(sort.FloatArray(af))
        fmt.Println(ai)
        fmt.Println(af)
}
// ---------------------------------------------------------
package main
 
import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)
var (
    endNum int //设置生成数的范围
)
func main() {
    i := createRandomNumber(endNum)
    //fmt.Println("生成规定范围内的整数:", i)    //本句调试用
 
    fmt.Println("请输入整数,范围为:0-", endNum)
 
    flag := true
    reader := bufio.NewReader(os.Stdin)
 
    for flag {
        data, _, _ := reader.ReadLine()
 
        command, err := strconv.Atoi(string(data)) //string to int,并作输入格式判断
        if err != nil {
            fmt.Println("格式不对，请输入数字")
        } else {
            fmt.Println("你输入的数字:", command)
 
            if command == i {
                flag = false
                fmt.Println("恭喜你，答对了~")
            } else if command < i {
                fmt.Println("你输入的数字小于生成的数字，别灰心！再来一次~")
            } else if command > i {
                fmt.Println("你输入的数字大于生成的数字，别灰心！再来一次~")
            }
        }
    }
}

func init() {
    endNum = 10
}
 
//生成规定范围内的整数
//设置起始数字范围，0开始,endNum截止
func createRandomNumber(endNum int) int {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    return r.Intn(endNum)
}
// ---------------------------------------------------------
package main
 
import (
  "fmt"
  "math"
)
 
func palindrome(num int) bool {
  var n, reverse, dig int
  n = num
  reverse = 0
  for (num > 0){
    dig = num % 10
    reverse = reverse * 10 + dig
    num = num / 10
  }
  return n == reverse
}
 
func main() {
  var cases int
  fmt.Scan(&cases)
  for i := 0; i < cases; i++ {
    var found, start, finish, sqrt_start, sqrt_finish, square int
    fmt.Scan(&start, &finish)
    sqrt_start = int(math.Sqrt(float64(start)))
    sqrt_finish = int(math.Sqrt(float64(finish)))
    for j := sqrt_start; j <= sqrt_finish; j++ {
      if palindrome(j){
        square = j*j
        if palindrome(square) && square >= start && square <= finish {
          found += 1
        }
      }
    }
    fmt.Print("Case #", (i + 1), ": ", found, "\n")
  }
}
// ---------------------------------------------------------
// use closure method to implement static variable
func f() func() int {
  i := 100
  return func() int { i++ ; return i }
}

// but the following code shows that its behavior isn't exactly same as static variable
func f() func() int {
  i := 0
  return func() int { i++ ; return i }
}

func meth(){
  for i := 0; i < 2; i++ {
    next := f()
    println(next())
    println(next(), next())
  }
}

func main() {
  meth()
}

O/P: 
1
2 3
1
2 3
// ---------------------------------------------------------
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("%+v\n\n", r);
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
// ---------------------------------------------------------
    func Errorf(format string, a ...interface{}) error

    func Fprint(w io.Writer, a ...interface{}) (n int, err error)
    func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
    func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

    func Fscan(r io.Reader, a ...interface{}) (n int, err error)
    func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
    func Fscanln(r io.Reader, a ...interface{}) (n int, err error)

    func Print(a ...interface{}) (n int, err error)
    func Printf(format string, a ...interface{}) (n int, err error)
    func Println(a ...interface{}) (n int, err error)

    func Scan(a ...interface{}) (n int, err error)
    func Scanf(format string, a ...interface{}) (n int, err error)
    func Scanln(a ...interface{}) (n int, err error)

    func Sprint(a ...interface{}) string
    func Sprintf(format string, a ...interface{}) string
    func Sprintln(a ...interface{}) string

    func Sscan(str string, a ...interface{}) (n int, err error)
    func Sscanf(str string, format string, a ...interface{}) (n int, err error)
    func Sscanln(str string, a ...interface{}) (n int, err error)

    type Formatter
    type GoStringer
    type ScanState
    type Scanner
    type State
    type Stringer
// ---------------------------------------------------------
golang跑满cpu

package main
 
import (
    "runtime"
)
 
func main() {
 
    runtime.GOMAXPROCS(runtime.NumCPU())
 
    for i := 0; i < runtime.NumCPU(); i++ {
        go func() {
            for {
            }
        }()
    }
 
    for {
    }
}

func main() { for { go func() { for { select { default: for { } } } }() } }
// ---------------------------------------------------------
func PrintRequest(r *http.Request) {
    //fmt.Printf("%v\n\n", r);   // value
    //fmt.Printf("%+v\n\n", r);  // fieldname + value
    //fmt.Printf("%#v\n\n", r);  // type + fieldname + value

    fmt.Printf("Method = %#v\n", r.Method)

    fmt.Printf("URL = %+v\n", r.URL)
    fmt.Printf("\tScheme = %#v\n", r.URL.Scheme)
    fmt.Printf("\tOpaque = %#v\n", r.URL.Opaque)
    fmt.Printf("\tUser = %#v\n", r.URL.User)
    fmt.Printf("\tHost = %#v\n", r.URL.Host)
    fmt.Printf("\tPath = %#v\n", r.URL.Path)
    fmt.Printf("\tRawQuery = %#v\n", r.URL.RawQuery)
    fmt.Printf("\tFragment = %#v\n", r.URL.Fragment)

    fmt.Printf("Proto = %#v\n", r.Proto)
    fmt.Printf("ProtoMajor = %#v\n", r.ProtoMajor)
    fmt.Printf("ProtoMinor = %#v\n", r.ProtoMinor)

    fmt.Printf("Header = %+v\n", r.Header)
    fmt.Printf("\tHeader.User-Agent = %#v\n", r.Header["User-Agent"])
    fmt.Printf("\tHeader.Accept = %#v\n", r.Header["Accept"])
    fmt.Printf("\tHeader.Accept-Language = %#v\n", r.Header["Accept-Language"])
    fmt.Printf("\tHeader.Accept-Encodinge = %#v\n", r.Header["Accept-Encoding"])
    fmt.Printf("\tHeader.Connection = %#v\n", r.Header["Connection"])

    fmt.Printf("Body = %#v\n", r.Body)
    fmt.Printf("ContentLength = %#v\n", r.ContentLength)
    fmt.Printf("TransferEncoding = %#v\n", r.TransferEncoding)
    fmt.Printf("Close = %#v\n", r.Close)
    fmt.Printf("Host = %#v\n", r.Host)
    fmt.Printf("Form = %#v\n", r.Form)
    fmt.Printf("PostForm = %#v\n", r.PostForm)
    fmt.Printf("MultipartForm = %#v\n", r.MultipartForm)
    fmt.Printf("Trailer = %#v\n", r.Trailer)
    fmt.Printf("RemoteAddr = %#v\n", r.RemoteAddr)
    fmt.Printf("RequestURI = %#v\n", r.RequestURI)
    fmt.Printf("TLS = %#v\n", r.TLS)
}

&{
Method:GET 
URL:/laptop?ref=lihs&name=andy 
Proto:HTTP/1.1 
ProtoMajor:1 
ProtoMinor:1 
Header:map[
    User-Agent:[Mozilla/5.0 (Windows NT 6.1; rv:21.0) Gecko/20100101 Firefox/21.0] 
    Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8] 
    Accept-Language:[en-US,en;q=0.5] 
    Accept-Encoding:[gzip, deflate] 
    Connection:[keep-alive]
] 
Body:0x11939220 
ContentLength:0 
TransferEncoding:[] 
Close:false 
Host:localhost:8080 
Form:map[] 
PostForm:map[] 
MultipartForm:<nil> 
Trailer:map[] 
RemoteAddr:[::1]:49609 
RequestURI:/laptop?ref=lihs&name=andy 
TLS:<nil>
}

*/
// ---------------------------------------------------------
GOLANG strings

func Contains(s, substr string) bool
func ContainsAny(s, chars string) bool
func ContainsRune(s string, r rune) bool
func Count(s, sep string) int
func EqualFold(s, t string) bool
func Fields(s string) []string
func FieldsFunc(s string, f func(rune) bool) []string
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Index(s, sep string) int
func IndexAny(s, chars string) int
func IndexFunc(s string, f func(rune) bool) int
func IndexRune(s string, r rune) int
func Join(a []string, sep string) string
func LastIndex(s, sep string) int
func LastIndexAny(s, chars string) int
func LastIndexFunc(s string, f func(rune) bool) int
func Map(mapping func(rune) rune, s string) string
func Repeat(s string, count int) string
func Replace(s, old, new string, n int) string
func Split(s, sep string) []string
func SplitAfter(s, sep string) []string
func SplitAfterN(s, sep string, n int) []string
func SplitN(s, sep string, n int) []string
func Title(s string) string
func ToLower(s string) string
func ToLowerSpecial(_case unicode.SpecialCase, s string) string
func ToTitle(s string) string
func ToTitleSpecial(_case unicode.SpecialCase, s string) string
func ToUpper(s string) string
func ToUpperSpecial(_case unicode.SpecialCase, s string) string
func Trim(s string, cutset string) string
func TrimFunc(s string, f func(rune) bool) string
func TrimLeft(s string, cutset string) string
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimPrefix(s, prefix string) string
func TrimRight(s string, cutset string) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimSpace(s string) string
func TrimSuffix(s, suffix string) string
type Reader
    func NewReader(s string) *Reader
    func (r *Reader) Len() int
    func (r *Reader) Read(b []byte) (n int, err error)
    func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
    func (r *Reader) ReadByte() (b byte, err error)
    func (r *Reader) ReadRune() (ch rune, size int, err error)
    func (r *Reader) Seek(offset int64, whence int) (int64, error)
    func (r *Reader) UnreadByte() error
    func (r *Reader) UnreadRune() error
    func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
type Replacer
    func NewReplacer(oldnew ...string) *Replacer
    func (r *Replacer) Replace(s string) string
    func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
// ---------------------------------------------------------
GOLANG 问题

1、构造函数？
2、静态变量？
3、静态方法(类方法)?
4、类常量？
5、类继承？
6、接口继承？
7、类成员和类方法分离的好处？
8、接口定义和实现分离的好处？
9、返回错误值与抛出异常的优劣比较？
0、如何表示“任意”类型？like void*
1、可变参数？
// ---------------------------------------------------------
GOLANG ioutil

var Discard io.Writer = devNull(0)

func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadDir(dirname string) ([]os.FileInfo, error)
func ReadFile(filename string) ([]byte, error)
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, prefix string) (f *os.File, err error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
// ---------------------------------------------------------
GOLANG bufio

const (
    // Maximum size used to buffer a token. The actual maximum token size
    // may be smaller as the buffer may need to include, for instance, a newline.
    MaxScanTokenSize = 64 * 1024
)

var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
var (
    ErrTooLong         = errors.New("bufio.Scanner: token too long")
    ErrNegativeAdvance = errors.New("bufio.Scanner: SplitFunc returns negative advance count")
    ErrAdvanceTooFar   = errors.New("bufio.Scanner: SplitFunc returns advance count beyond input")
)

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
type ReadWriter
    func NewReadWriter(r *Reader, w *Writer) *ReadWriter
type Reader
    func NewReader(rd io.Reader) *Reader
    func NewReaderSize(rd io.Reader, size int) *Reader
    func (b *Reader) Buffered() int
    func (b *Reader) Peek(n int) ([]byte, error)
    func (b *Reader) Read(p []byte) (n int, err error)
    func (b *Reader) ReadByte() (c byte, err error)
    func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
    func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
    func (b *Reader) ReadRune() (r rune, size int, err error)
    func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
    func (b *Reader) ReadString(delim byte) (line string, err error)
    func (b *Reader) UnreadByte() error
    func (b *Reader) UnreadRune() error
    func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
type Scanner
    func NewScanner(r io.Reader) *Scanner
    func (s *Scanner) Bytes() []byte
    func (s *Scanner) Err() error
    func (s *Scanner) Scan() bool
    func (s *Scanner) Split(split SplitFunc)
    func (s *Scanner) Text() string
type SplitFunc
type Writer
    func NewWriter(wr io.Writer) *Writer
    func NewWriterSize(wr io.Writer, size int) *Writer
    func (b *Writer) Available() int
    func (b *Writer) Buffered() int
    func (b *Writer) Flush() error
    func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
    func (b *Writer) Write(p []byte) (nn int, err error)
    func (b *Writer) WriteByte(c byte) error
    func (b *Writer) WriteRune(r rune) (size int, err error)
    func (b *Writer) WriteString(s string) (int, error)
// ---------------------------------------------------------
GOLANG os

const (
    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
    O_RDWR   int = syscall.O_RDWR   // open the file read-write.
    O_APPEND int = syscall.O_APPEND // append data to the file when writing.
    O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
    O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
    O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
    O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)
Flags to Open wrapping those of the underlying system. Not all flags may be implemented on a given system.

const (
    SEEK_SET int = 0 // seek relative to the origin of the file
    SEEK_CUR int = 1 // seek relative to the current offset
    SEEK_END int = 2 // seek relative to the end
)
Seek whence values.

const (
    PathSeparator     = '/' // OS-specific path separator
    PathListSeparator = ':' // OS-specific path list separator
)
const DevNull = "/dev/null"
DevNull is the name of the operating system's “null device.” On Unix-like systems, it is "/dev/null"; on Windows, "NUL".

Variables

var (
    ErrInvalid    = errors.New("invalid argument")
    ErrPermission = errors.New("permission denied")
    ErrExist      = errors.New("file already exists")
    ErrNotExist   = errors.New("file does not exist")
)
Portable analogs of some common system call errors.

var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.

var Args []string
Args hold the command-line arguments, starting with the program name.

func Chdir(dir string) error
func Chmod(name string, mode FileMode) error
func Chown(name string, uid, gid int) error
func Chtimes(name string, atime time.Time, mtime time.Time) error
func Clearenv()
func Environ() []string
func Exit(code int)
func Expand(s string, mapping func(string) string) string
func ExpandEnv(s string) string
func Getegid() int
func Getenv(key string) string
func Geteuid() int
func Getgid() int
func Getgroups() ([]int, error)
func Getpagesize() int
func Getpid() int
func Getppid() int
func Getuid() int
func Getwd() (pwd string, err error)
func Hostname() (name string, err error)
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPathSeparator(c uint8) bool
func IsPermission(err error) bool
func Lchown(name string, uid, gid int) error
func Link(oldname, newname string) error
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error
func NewSyscallError(syscall string, err error) error
func Readlink(name string) (string, error)
func Remove(name string) error
func RemoveAll(path string) error
func Rename(oldname, newname string) error
func SameFile(fi1, fi2 FileInfo) bool
func Setenv(key, value string) error
func Symlink(oldname, newname string) error
func TempDir() string
func Truncate(name string, size int64) error
type File
    func Create(name string) (file *File, err error)
    func NewFile(fd uintptr, name string) *File
    func Open(name string) (file *File, err error)
    func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
    func Pipe() (r *File, w *File, err error)
    func (f *File) Chdir() error
    func (f *File) Chmod(mode FileMode) error
    func (f *File) Chown(uid, gid int) error
    func (f *File) Close() error
    func (f *File) Fd() uintptr
    func (f *File) Name() string
    func (f *File) Read(b []byte) (n int, err error)
    func (f *File) ReadAt(b []byte, off int64) (n int, err error)
    func (f *File) Readdir(n int) (fi []FileInfo, err error)
    func (f *File) Readdirnames(n int) (names []string, err error)
    func (f *File) Seek(offset int64, whence int) (ret int64, err error)
    func (f *File) Stat() (fi FileInfo, err error)
    func (f *File) Sync() (err error)
    func (f *File) Truncate(size int64) error
    func (f *File) Write(b []byte) (n int, err error)
    func (f *File) WriteAt(b []byte, off int64) (n int, err error)
    func (f *File) WriteString(s string) (ret int, err error)
type FileInfo
    func Lstat(name string) (fi FileInfo, err error)
    func Stat(name string) (fi FileInfo, err error)
type FileMode
    func (m FileMode) IsDir() bool
    func (m FileMode) IsRegular() bool
    func (m FileMode) Perm() FileMode
    func (m FileMode) String() string
type LinkError
    func (e *LinkError) Error() string
type PathError
    func (e *PathError) Error() string
type ProcAttr
type Process
    func FindProcess(pid int) (p *Process, err error)
    func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
    func (p *Process) Kill() error
    func (p *Process) Release() error
    func (p *Process) Signal(sig Signal) error
    func (p *Process) Wait() (*ProcessState, error)
type ProcessState
    func (p *ProcessState) Exited() bool
    func (p *ProcessState) Pid() int
    func (p *ProcessState) String() string
    func (p *ProcessState) Success() bool
    func (p *ProcessState) Sys() interface{}
    func (p *ProcessState) SysUsage() interface{}
    func (p *ProcessState) SystemTime() time.Duration
    func (p *ProcessState) UserTime() time.Duration
type Signal
type SyscallError
    func (e *SyscallError) Error() string
// ---------------------------------------------------------
package main
 
import (
    "fmt"
    "io/ioutil"
    "os/exec"
    "time"
)
 
func run() {
    cmd := exec.Command("/bin/sh", "-c", "ping 127.0.0.1")
    _, err := cmd.Output()
    if err != nil {
        panic(err.Error())
    }
 
    if err := cmd.Start(); err != nil {
        panic(err.Error())
    }
 
    if err := cmd.Wait(); err != nil {
        panic(err.Error())
    }
}
 
func main() {
    go run()
    time.Sleep(1e9)
 
    cmd := exec.Command("/bin/sh", "-c", `ps -ef | grep -v "grep" | grep "ping"`)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("StdoutPipe: " + err.Error())
        return
    }
 
    stderr, err := cmd.StderrPipe()
    if err != nil {
        fmt.Println("StderrPipe: ", err.Error())
        return
    }
 
    if err := cmd.Start(); err != nil {
        fmt.Println("Start: ", err.Error())
        return
    }
 
    bytesErr, err := ioutil.ReadAll(stderr)
    if err != nil {
        fmt.Println("ReadAll stderr: ", err.Error())
        return
    }
 
    if len(bytesErr) != 0 {
        fmt.Printf("stderr is not nil: %s", bytesErr)
        return
    }
 
    bytes, err := ioutil.ReadAll(stdout)
    if err != nil {
        fmt.Println("ReadAll stdout: ", err.Error())
        return
    }
 
    if err := cmd.Wait(); err != nil {
        fmt.Println("Wait: ", err.Error())
        return
    }
 
    fmt.Printf("stdout: %s", bytes)
}
 
// 运行命令: go run main.go
// ---------------------------------------------------------
// MarkGuid project main.go
package main
 
import (
    "crypto/md5"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "io"
)
 
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}
 
func GetGuid() string {
    b := make([]byte, 48)
 
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
 
func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(GetGuid())
    }
}
// ---------------------------------------------------------
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}

type CookieJar interface {
    // SetCookies handles the receipt of the cookies in a reply for the
    // given URL.  It may or may not choose to save the cookies, depending
    // on the jar's policy and implementation.
    SetCookies(u *url.URL, cookies []*Cookie)

    // Cookies returns the cookies to send in a request for the given URL.
    // It is up to the implementation to honor the standard cookie use
    // restrictions such as in RFC 6265.
    Cookies(u *url.URL) []*Cookie
}

type Request struct {
    Method string // GET, POST, PUT, etc.

    // URL is created from the URI supplied on the Request-Line
    // as stored in RequestURI.
    //
    // For most requests, fields other than Path and RawQuery
    // will be empty. (See RFC 2616, Section 5.1.2)
    URL *url.URL

    // The protocol version for incoming requests.
    // Outgoing requests always use HTTP/1.1.
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0

    // A header maps request lines to their values.
    // If the header says
    //
    //	accept-encoding: gzip, deflate
    //	Accept-Language: en-us
    //	Connection: keep-alive
    //
    // then
    //
    //	Header = map[string][]string{
    //		"Accept-Encoding": {"gzip, deflate"},
    //		"Accept-Language": {"en-us"},
    //		"Connection": {"keep-alive"},
    //	}
    //
    // HTTP defines that header names are case-insensitive.
    // The request parser implements this by canonicalizing the
    // name, making the first character and any characters
    // following a hyphen uppercase and the rest lowercase.
    Header Header

    // The message body.
    Body io.ReadCloser

    // ContentLength records the length of the associated content.
    // The value -1 indicates that the length is unknown.
    // Values >= 0 indicate that the given number of bytes may
    // be read from Body.
    // For outgoing requests, a value of 0 means unknown if Body is not nil.
    ContentLength int64

    // TransferEncoding lists the transfer encodings from outermost to
    // innermost. An empty list denotes the "identity" encoding.
    // TransferEncoding can usually be ignored; chunked encoding is
    // automatically added and removed as necessary when sending and
    // receiving requests.
    TransferEncoding []string

    // Close indicates whether to close the connection after
    // replying to this request.
    Close bool

    // The host on which the URL is sought.
    // Per RFC 2616, this is either the value of the Host: header
    // or the host name given in the URL itself.
    // It may be of the form "host:port".
    Host string

    // Form contains the parsed form data, including both the URL
    // field's query parameters and the POST or PUT form data.
    // This field is only available after ParseForm is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from POST or PUT
    // body parameters.
    // This field is only available after ParseForm is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values

    // MultipartForm is the parsed multipart form, including file uploads.
    // This field is only available after ParseMultipartForm is called.
    // The HTTP client ignores MultipartForm and uses Body instead.
    MultipartForm *multipart.Form

    // Trailer maps trailer keys to values.  Like for Header, if the
    // response has multiple trailer lines with the same key, they will be
    // concatenated, delimited by commas.
    // For server requests, Trailer is only populated after Body has been
    // closed or fully consumed.
    // Trailer support is only partially complete.
    Trailer Header

    // RemoteAddr allows HTTP servers and other software to record
    // the network address that sent the request, usually for
    // logging. This field is not filled in by ReadRequest and
    // has no defined format. The HTTP server in this package
    // sets RemoteAddr to an "IP:port" address before invoking a
    // handler.
    // This field is ignored by the HTTP client.
    RemoteAddr string

    // RequestURI is the unmodified Request-URI of the
    // Request-Line (RFC 2616, Section 5.1) as sent by the client
    // to a server. Usually the URL field should be used instead.
    // It is an error to set this field in an HTTP client request.
    RequestURI string

    // TLS allows HTTP servers and other software to record
    // information about the TLS connection on which the request
    // was received. This field is not filled in by ReadRequest.
    // The HTTP server in this package sets the field for
    // TLS-enabled connections before invoking a handler;
    // otherwise it leaves the field nil.
    // This field is ignored by the HTTP client.
    TLS *tls.ConnectionState
}

type Response struct {
    Status     string // e.g. "200 OK"
    StatusCode int    // e.g. 200
    Proto      string // e.g. "HTTP/1.0"
    ProtoMajor int    // e.g. 1
    ProtoMinor int    // e.g. 0

    // Header maps header keys to values.  If the response had multiple
    // headers with the same key, they will be concatenated, with comma
    // delimiters.  (Section 4.2 of RFC 2616 requires that multiple headers
    // be semantically equivalent to a comma-delimited sequence.) Values
    // duplicated by other fields in this struct (e.g., ContentLength) are
    // omitted from Header.
    //
    // Keys in the map are canonicalized (see CanonicalHeaderKey).
    Header Header

    // Body represents the response body.
    //
    // The http Client and Transport guarantee that Body is always
    // non-nil, even on responses without a body or responses with
    // a zero-lengthed body.
    //
    // The Body is automatically dechunked if the server replied
    // with a "chunked" Transfer-Encoding.
    Body io.ReadCloser

    // ContentLength records the length of the associated content.  The
    // value -1 indicates that the length is unknown.  Unless Request.Method
    // is "HEAD", values >= 0 indicate that the given number of bytes may
    // be read from Body.
    ContentLength int64

    // Contains transfer encodings from outer-most to inner-most. Value is
    // nil, means that "identity" encoding is used.
    TransferEncoding []string

    // Close records whether the header directed that the connection be
    // closed after reading Body.  The value is advice for clients: neither
    // ReadResponse nor Response.Write ever closes a connection.
    Close bool

    // Trailer maps trailer keys to values, in the same
    // format as the header.
    Trailer Header

    // The Request that was sent to obtain this Response.
    // Request's Body is nil (having already been consumed).
    // This is only populated for Client requests.
    Request *Request
}

type Server struct {
    Addr           string        // TCP address to listen on, ":http" if empty
    Handler        Handler       // handler to invoke, http.DefaultServeMux if nil
    ReadTimeout    time.Duration // maximum duration before timing out read of the request
    WriteTimeout   time.Duration // maximum duration before timing out write of the response
    MaxHeaderBytes int           // maximum size of request headers, DefaultMaxHeaderBytes if 0
    TLSConfig      *tls.Config   // optional TLS config, used by ListenAndServeTLS

    // TLSNextProto optionally specifies a function to take over
    // ownership of the provided TLS connection when an NPN
    // protocol upgrade has occurred.  The map key is the protocol
    // name negotiated. The Handler argument should be used to
    // handle HTTP requests and will initialize the Request's TLS
    // and RemoteAddr if not already set.  The connection is
    // automatically closed when the function returns.
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
}
// ---------------------------------------------------------
Three ways to use a Timer

// (A)
time.AfterFunc(5 * time.Minute, func() {
    fmt.Printf("expired")
}

// (B) create a Timer object
timer := time.NewTimer(5 * time.Minute)
<-timer.C
fmt.Printf("expired")

// (C) time.After() returns timer.C internally
<-time.After(5 * time.Minute)
fmt.Printf("expired")
// ---------------------------------------------------------
GoLang操作MySQL数据库

package main
 
import (
    "fmt"
    "database/sql"
    _"mysql"
)
 
type TestMysql struct {
    db *sql.DB
}
 
/* 初始化数据库引擎 */
func Init() (*TestMysql,error){
    test := new(TestMysql);
    db,err := sql.Open("mysql","test:test@tcp(127.0.0.1:3306)/abwork?charset=utf8");
    //第一个参数 ： 数据库引擎
    //第二个参数 : 数据库DSN配置。Go中没有统一DSN,都是数据库引擎自己定义的，因此不同引擎可能配置不同
    //本次演示采用http://code.google.com/p/go-mysql-driver
    if err!=nil {
        fmt.Println("database initialize error : ",err.Error());
        return nil,err;
    }
    test.db = db;
    return test,nil;
}
 
/* 测试数据库数据添加 */
func (test *TestMysql)Create(){
    if test.db==nil {
        return;
    }
    stmt,err := test.db.Prepare("insert into test(name,age)values(?,?)");
    if err!=nil {
        fmt.Println(err.Error());
        return;
    }
    defer stmt.Close();
    if result,err := stmt.Exec("张三",20);err==nil {
        if id,err := result.LastInsertId();err==nil {
            fmt.Println("insert id : ",id);
        }
    }
    if result,err := stmt.Exec("李四",30);err==nil {
        if id,err := result.LastInsertId();err==nil {
            fmt.Println("insert id : ",id);
        }
    }
    if result,err := stmt.Exec("王五",25);err==nil {
        if id,err := result.LastInsertId();err==nil {
            fmt.Println("insert id : ",id);
        }
    }
}
 
/* 测试数据库数据更新 */
func (test *TestMysql)Update(){
    if test.db==nil {
        return;
    }
    stmt,err := test.db.Prepare("update test set name=?,age=? where age=?");
    if err!=nil {
        fmt.Println(err.Error());
        return;
    }
    defer stmt.Close();
    if result,err := stmt.Exec("周七",40,25);err==nil {
        if c,err := result.RowsAffected();err==nil {
            fmt.Println("update count : ",c);
        }
    }
}
 
/* 测试数据库数据读取 */
func (test *TestMysql)Read(){
    if test.db==nil {
        return;
    }
    rows,err := test.db.Query("select id,name,age from test limit 0,5");
    if err!=nil {
        fmt.Println(err.Error());
        return;
    }
    defer rows.Close();
    fmt.Println("");
    cols,_ :=  rows.Columns();
    for i := range cols {
        fmt.Print(cols[i]);
        fmt.Print("\t");
    }
    fmt.Println("");
    var id int;
    var name string;
    var age int;
    for rows.Next(){
        if err := rows.Scan(&id,&name,&age);err==nil {
            fmt.Print(id);
            fmt.Print("\t");
            fmt.Print(name);
            fmt.Print("\t");
            fmt.Print(age);
            fmt.Print("\t\r\n");
        }
    }
}
 
/* 测试数据库删除 */
func (test *TestMysql)Delete(){
    if test.db==nil {
        return;
    }
    stmt,err := test.db.Prepare("delete from test where age=?");
    if err!=nil {
        fmt.Println(err.Error());
        return;
    }
    defer stmt.Close();
    if result,err := stmt.Exec(20);err==nil {
        if c,err :=  result.RowsAffected();err==nil{
            fmt.Println("remove count : ",c);
        }
    }
}
 
func (test *TestMysql)Close(){
    if test.db!=nil {
        test.db.Close();
    }
}
 
func main(){
    if test,err := Init();err==nil {
        test.Create();
        test.Update();
        test.Read();
        test.Delete();
        test.Read();
        test.Close();
    }
}
// ---------------------------------------------------------
Wait until all the background goroutine finish

var w sync.WaitGroup
w.Add(2)
go func() {
    // do something
    w.Done()
}
go func() {
    // do something
    w.Done()
}
w.Wait()
// ---------------------------------------------------------
Parse HTML

import (
    "code.google.com/p/go.net/html"
)

func parseHtml(r io.Reader) {
    d := html.NewTokenizer(r)
    for { 
        // token type
        tokenType := d.Next() 
        if tokenType == html.ErrorToken {
            return     
        }       
        token := d.Token()
        switch tokenType {
            case html.StartTagToken: // <tag>
                // type Token struct {
                //     Type     TokenType
                //     DataAtom atom.Atom
                //     Data     string
                //     Attr     []Attribute
                // }
                //
                // type Attribute struct {
                //     Namespace, Key, Val string
                // }
            case html.TextToken: // text between start and end tag
            case html.EndTagToken: // </tag>
            case html.SelfClosingTagToken: // <tag/>

        }
    }
}
// ---------------------------------------------------------
type Transport struct {

    // Proxy specifies a function to return a proxy for a given
    // Request. If the function returns a non-nil error, the
    // request is aborted with the provided error.
    // If Proxy is nil or returns a nil *URL, no proxy is used.
    Proxy func(*Request) (*url.URL, error)

    // Dial specifies the dial function for creating TCP
    // connections.
    // If Dial is nil, net.Dial is used.
    Dial func(network, addr string) (net.Conn, error)

    // TLSClientConfig specifies the TLS configuration to use with
    // tls.Client. If nil, the default configuration is used.
    TLSClientConfig *tls.Config

    // DisableKeepAlives, if true, prevents re-use of TCP connections
    // between different HTTP requests.
    DisableKeepAlives bool

    // DisableCompression, if true, prevents the Transport from
    // requesting compression with an "Accept-Encoding: gzip"
    // request header when the Request contains no existing
    // Accept-Encoding value. If the Transport requests gzip on
    // its own and gets a gzipped response, it's transparently
    // decoded in the Response.Body. However, if the user
    // explicitly requested gzip it is not automatically
    // uncompressed.
    DisableCompression bool

    // MaxIdleConnsPerHost, if non-zero, controls the maximum idle
    // (keep-alive) to keep per-host.  If zero,
    // DefaultMaxIdleConnsPerHost is used.
    MaxIdleConnsPerHost int

    // ResponseHeaderTimeout, if non-zero, specifies the amount of
    // time to wait for a server's response headers after fully
    // writing the request (including its body, if any). This
    // time does not include the time to read the response body.
    ResponseHeaderTimeout time.Duration
    // contains filtered or unexported fields
}

const (
    StatusContinue           = 100
    StatusSwitchingProtocols = 101

    StatusOK                   = 200
    StatusCreated              = 201
    StatusAccepted             = 202
    StatusNonAuthoritativeInfo = 203
    StatusNoContent            = 204
    StatusResetContent         = 205
    StatusPartialContent       = 206

    StatusMultipleChoices   = 300
    StatusMovedPermanently  = 301
    StatusFound             = 302
    StatusSeeOther          = 303
    StatusNotModified       = 304
    StatusUseProxy          = 305
    StatusTemporaryRedirect = 307

    StatusBadRequest                   = 400
    StatusUnauthorized                 = 401
    StatusPaymentRequired              = 402
    StatusForbidden                    = 403
    StatusNotFound                     = 404
    StatusMethodNotAllowed             = 405
    StatusNotAcceptable                = 406
    StatusProxyAuthRequired            = 407
    StatusRequestTimeout               = 408
    StatusConflict                     = 409
    StatusGone                         = 410
    StatusLengthRequired               = 411
    StatusPreconditionFailed           = 412
    StatusRequestEntityTooLarge        = 413
    StatusRequestURITooLong            = 414
    StatusUnsupportedMediaType         = 415
    StatusRequestedRangeNotSatisfiable = 416
    StatusExpectationFailed            = 417
    StatusTeapot                       = 418

    StatusInternalServerError     = 500
    StatusNotImplemented          = 501
    StatusBadGateway              = 502
    StatusServiceUnavailable      = 503
    StatusGatewayTimeout          = 504
    StatusHTTPVersionNotSupported = 505
)

type URL struct {
    Scheme   string
    Opaque   string    // encoded opaque data
    User     *Userinfo // username and password information
    Host     string    // host or host:port
    Path     string
    RawQuery string // encoded query values, without '?'
    Fragment string // fragment for references, without '#'
}
// ---------------------------------------------------------
Sleep forever

select{}

#Get random value
Get random value in 1..99

max := big.NewInt(100)
i, err := rand.Int(rand.Reader, max)

#HTTP handler func

Example handler func.

func someHandler(w http.ResponseWriter, r *http.Request) {
    // read form value
    value := r.FormValue("value")
    if r.Method == "POST" {
        // receive posted data
        body, err := ioutil.ReadAll(r.Body)

}

func main() {
    http.HandleFunc("/", someHandler)
    http.ListenAndServe(":8080", nil)
}
// ---------------------------------------------------------
// 原来golang下载文件这么简单

// goGetJpg
package main
 
import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "io"
    "net/http"
    "os"
)
 
func main() {
    x, _ := goquery.NewDocument("http://www.fengyun5.com/Sibao/600/1.html")
    urls, _ := x.Find("#content img").Attr("src")
    res, _ := http.Get(urls)
    file, _ := os.Create("xxx.jpg")
    io.Copy(file, res.Body)
    fmt.Println("下载完成！")
}

// ---------------------------------------------------------
