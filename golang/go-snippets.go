golang ����Ƭ��

�ļ����ж�ȡ.go

f, err := os.Open("a.txt")
defer f.Close()
if nil == err {
    buff := bufio.NewReader(f)
    for {
        line, err := buff.ReadString('\n')
        if err != nil || io.EOF == err{
            break
        }
        fmt.Println(line)
    }
}

��chanʵ��쳲���������

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

����chan����ʱ(��ʱ)

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
}

�����������ö���������

package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	data := []int{1, 5, 6, 3, 2, 7, 8, 9}
	Sort(data)
	fmt.Println(data)
}

ͨ���������ʽ������ѡ���ĸ�����

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
    var op func(p, q Point) Point
    if add {
        op = Point.Add
    } else {
        op = Point.Sub
    }
    for i := range path {
        // Call either path[i].Add(offset) or path[i].Sub(offset).
        path[i] = op(path[i], offset)
    }
}

ch���channel��buffer��С��1�����Իύ���Ϊ�ջ�Ϊ��������ֻ��
һ��case���Խ�����ȥ������i����������ż�����������ӡ0 2 4 6 8��

ch := make(chan int, 1)
for i := 0; i < 10; i++ {
    select {
    case x := <-ch:
        fmt.Println(x) // "0" "2" "4" "6" "8"
    case ch <- i:
    }
}

�������ļ�ʵ��

func worker(start chan bool) {
	heartbeat := time.Tick(30 * time.Second)
	for {
		select {
			// �� do some stuff
		case <- heartbeat:
			//�� do heartbeat stuff
		}
	}
}

ɾ��ĳ��slice��ĳ��Ԫ��

for i := range s {
    if equal(s[i], element) {
		s = append(s[:i], s[i+1:]...)
    }
}

������תslice

func main() {
    a := [3]int{1, 2, 3}
    b := a[:]
    fmt.Println(b)
}
func main(){
    a := [3]int{1, 2, 3}
    b := (&a)[:]
}

����ת��

func forward(src net.Conn, network, address string, timeout time.Duration) {
	defer src.Close()
	dst, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		log.Printf("dial err: %s", err)
		return
	}
	defer dst.Close()

	cpErr := make(chan error)

	go cp(cpErr, src, dst)
	go cp(cpErr, dst, src)

	select {
	case err = <-cpErr:
		if err != nil {
			log.Printf("copy err: %v", err)
		}
	}

	log.Printf("disconnect: %s", src.RemoteAddr())
}

func cp(c chan error, w io.Writer, r io.Reader) {
	_, err := io.Copy(w, r)
	c <- err
	fmt.Println("cp end")
}

��ȡ��ǰĿ¼�͸���Ŀ¼

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

���浱ǰ�����pid

func savePid() {
	pidFilename := ROOT + "/pid/" +filepath.Base(os.Args[0]) + ".pid"
	pid := os.Getpid()

	ioutil.WriteFile(pidFilename, []byte(strconv.Itoa(pid)), 0755)
}

����struct{}��channel

ch:=make(chan struct{})
ch <- struct{}{}
<-ch

break��label

OuterLoop:
	for i = 0; i < n; i++ {
		for j = 0; j < m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}

Go����ʵ��http�����ļ�������

package main

import (
	"net/http"
	"os"
	"strings"
)
func shareDir(dirName string,port string,ch chan bool){
	h := http.FileServer(http.Dir(dirName))
	err := http.ListenAndServe(":"+port,h)
	if err != nil {
		println("ListenAndServe : ",err.Error())
		ch <- false
	}
}
func main(){
	ch := make(chanbool)
	port := "8000"//Default port
	if len(os.Args)>1 {
		port = strings.Join(os.Args[1:2],"")
	}
	go shareDir(".",port,ch)
	println("Listening on port ",port,"...")
	bresult := <-ch
	if false == bresult {
		println("Listening on port ",port," failed")
	}
}

ѡ�������ǲ��ȶ��ġ��㷨���Ӷ���O(n ^2 )��

package main
import (
    "fmt"
)

type SortInterface interface {
    sort()
}
type Sortor struct {
    name string
}

func main() {
    arry := []int{6, 1, 3, 5, 8, 4, 2, 0, 9, 7}
    learnsort := Sortor{name: "ѡ������--��С����--���ȶ�--n*n---"}
    learnsort.sort(arry)
    fmt.Println(learnsort.name, arry)
}

func (sorter Sortor) sort(arry []int) {
    arrylength := len(arry)    for i := 0; i < arrylength; i++ {
        min := i        for j := i + 1; j < arrylength; j++ {            if arry[j] < arry[min] {
                min = j
            }
        }
        t := arry[i]
        arry[i] = arry[min]
        arry[min] = t
    }
}

http://blog.csdn.net/guoer9973/article/details/51924715

����ļ���Ŀ¼�Ƿ����

// ����� filename ָ�����ļ���Ŀ¼�����򷵻� true�����򷵻� false
func fileExist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}

�ض���golang��panic��Ϣ

package main

import (
    "log"
    "os"
    "syscall"
)

// redirectStderr to the file passed in
func redirectStderr(f *os.File) {
    err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
    if err != nil {
        log.Fatalf("Failed to redirect stderr to file: %v", err)
    }
}

import (
	"code.google.com/p/log4go"
	"os"
	"syscall"
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle = kernel32.MustFindProc("SetStdHandle")
)

func SetStdHandle(stdhandle int32, handle syscall.Handle) error {
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

var f *os.File

func redirect_err() {
	var err error
	f, err = os.Create(`panic.txt`)
	if err != nil {
		log4go.Error("os.Create failed: %v", err)
	}
	err = SetStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		log4go.Error("SetStdHandle failed: %v", err)
	}
}

go ִ�������ȡstderr

    var stderr io.ReadCloser
    cmd := exec.Command("ping", "localhost")
    stderr, err = cmd.StderrPipe()
	if err != nil {
		glog.Errorln("stderr pipe error: ", err)
	}
	// start cmd
	err = cmd.Start()
	if err != nil {
		glog.Infoln("start cmd error: ", err)
		return err
	}

	glog.Infoln("start ", cmd.Args)
	if stderr != nil {
		reader := bufio.NewReader(stderr)
		go func() {
			for {
				line, err := reader.ReadString(byte('\n'))
				if err != nil || io.EOF == err {
					glog.Infoln("cmd end")
					stderr.Close()
					break
				}
				glog.Infoln("stderr: ", line)
			}
		}()
	}
	err:=cmd.Wait()
	if err!=nil {
	    return err
	}

�ж�accept�ǲ�����ʱ����

for {
    conn,err:=ln.Accept()
    if err, ok := err.(net.Error); ok && err.Temporary() {
        continue
    }
}

���ٱ�����ǰ�ļ��к��ļ�

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	upperDirPattern := "./*"
	matches, err := filepath.Glob(upperDirPattern)
	if err != nil {
		panic(err)
	}
	for _, file := range matches {
		fmt.Println(file)
	}
}

ͨ��http package����user:passwd��Ϣ

    ���������� curl -u ������
    client := &http.Client{}
    req, err := http.NewRequest("GET", <url>, nil)
    req.SetBasicAuth(<username>, <userpasswd>)
    if err != nil {
        log.Fatal(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    content, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

golangͨ�����佫δ֪����ת��Ϊarray

func interfaceSlice(slice interface{}) []interface{} {
    s := reflect.ValueOf(slice)
    if s.Kind() != reflect.Slice {
        panic("InterfaceSlice() given a non-slice type")
    }

    ret := make([]interface{}, s.Len())

    for i := 0; i < s.Len(); i++ {
        ret[i] = s.Index(i).Interface()
    }

    return ret
}

Mux ��ȡGet Query����

    vals := r.URL.Query()
    oriDriver, ok := vals["driver"]

Mux ������ʴ���

    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    log.Println(http.ListenAndServe(":8000", handlers.CORS(headersOk, originsOk, methodsOk)(r)))

�ݹ�����ļ�

func readAPK(path string, apk map[string]int) {
    // fmt.Println("����", path)
    files, _ := ioutil.ReadDir(path)
    for _, file := range files {
        // fmt.Println(file.Name(), file.IsDir())
        if file.IsDir() {
            readAPK(path+"/"+file.Name(), apk)
        } else {
            if strings.Compare(file.Name(), "APK.log") == 0 {
                ap := make(map[string]string)
                body, err := ioutil.ReadFile(path + "/" + file.Name())
                if err != nil {
                    fmt.Printf("[%s]��ȡʧ��[%s]\n", file.Name(), err.Error())
                    return
                }

                err = json.Unmarshal(body, &ap)
                if err != nil {
                    fmt.Printf("[%s]����ʧ��[%s]\n", file.Name(), err.Error())
                    return
                }

                for a := range ap {
                    at := strings.Split(ap[a], "|+|")
                    for _, atemp := range at {
                        info := strings.Split(atemp, "|-|")
                        if len(info) > 1 {
                            header := strings.Split(info[1], " ")
                            for _, h := range header {
                                if strings.Contains(h, "Referer:") {
                                    u, err := url.ParseRequestURI(h[8:])
                                    if err != nil {
                                        fmt.Println(err.Error())
                                    } else {
                                        apk[u.Host]++
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}

����Mongo�Ѿ����ڵ�����

type App struct {
    Id string `json:"id" bson:"_id,omitempty"`
    User_id string `bson:"user_id"`
    Name string `bson:"name"`
    Domain string `bson:"domain"`
    Business_line string `bson:"business_line"`
}

�ص�:
1.ʹ��bson���η�
2.����flagֱ��ʹ�ÿո�ָ�
3._id��omitemptyʹ��,�ָͬʱ���ܴ��ڿո�

��ʱ��

for {
	now := time.Now()
	next := now.Add(time.Minute * 10)
	next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	log.Printf("�´βɼ�ʱ��Ϊ[%s]\n", next.Format("200601021504"))

	select {
	case <-t.C:
		err := sync.Gather()
		if err != nil {
			log.Println(err)
		}
	}
}
