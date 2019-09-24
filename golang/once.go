package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		go once.Do(oneTime)
	}
	time.Sleep(3 * time.Second)
}

func oneTime() {
	fmt.Println("one time")
}
