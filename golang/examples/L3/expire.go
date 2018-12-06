package main

import  (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	
	<-timer1.C
	fmt.Println("time 1 exp")
	
	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("time 2 exp")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("time 2 stop")
	}
	fmt.Println("END");
}
