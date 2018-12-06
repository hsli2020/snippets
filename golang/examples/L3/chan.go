package main

import ("fmt")

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) 
	// 假如注释掉这一块的话~ fatal error: all goroutines are asleep - deadlock!

	//~ goroutine 1 [chan receive]:
	//~ main.main()
    //~ /home/wwwroot/chan.go:11 +0x98
	//~ exit status 2
	
	for elem := range queue {
		fmt.Println(elem)
	}
}
