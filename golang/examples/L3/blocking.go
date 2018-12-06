package main

import "fmt"

func main() {
	messages := make(chan string)
	signals  := make(chan bool)
	
	msg := "hi"
	select {
		case messages <- msg :
			fmt.Println("send message", msg)
		case sig := <-signals:
			fmt.Println("received",sig)
		default:
			fmt.Println("no send message")
	}
}
