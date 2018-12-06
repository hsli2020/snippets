package main

import (
	"fmt"
	"bytes"
	"strconv"
	"runtime"
	"time"
)

func main() {
	go foo()
	fmt.Println(GetGID())
	time.Sleep(1*time.Second)
}

func foo() {
	fmt.Println("Foo", GetGID())
}

// 获取 goroutine 的协程 id
func GetGID() uint64 {
    b := make([]byte, 64)
    b = b[:runtime.Stack(b, false)]
	fmt.Println(string(b))
    b = bytes.TrimPrefix(b, []byte("goroutine "))
    b = b[:bytes.IndexByte(b, ' ')]
    n, _ := strconv.ParseUint(string(b), 10, 64)
    return n
}
