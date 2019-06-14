// 原文链接：https://blog.csdn.net/qq_25504271/article/details/79825366
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("反正就是发生错误了")
		}
	}()

	file, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("打开文件错误")
	}

	//警告信息 stdout 输出  stderr 输出错误
	Warning = log.New(os.Stdout, "警告信息:", log.Ldate|log.Lmicroseconds)

	//提示信息
	Info = log.New(io.MultiWriter(file, os.Stderr), "提示信息：", log.Lmicroseconds|log.Ldate)

	//错误信息
	//discard  所有的io 不会操作 但是会返回成功
	Error = log.New(ioutil.Discard, "错误信息", log.Lmicroseconds)
}

func main() {
	Warning.Println("test")
}
