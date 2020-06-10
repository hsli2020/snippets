package main

import "log"
import "time"

func shouldNotExit() {
	for {
		time.Sleep(time.Second) // 模拟一个工作负载
		// 模拟一个未预料到的恐慌。
		if time.Now().UnixNano()&0x3 == 0 {
			panic("unexpected situation")
		}
	}
}

func NeverExit(name string, f func()) {
	defer func() {
		if v := recover(); v != nil { // 侦测到一个恐慌
			log.Printf("协程%s崩溃了，准备重启一个", name)
			go NeverExit(name, f) // 重启一个同功能协程
		}
	}()
	f()
}

func main() {
	log.SetFlags(0)
	go NeverExit("job#A", shouldNotExit)
	go NeverExit("job#B", shouldNotExit)
	select {} // 永久阻塞主线程
}
