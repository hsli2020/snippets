package main  // https://ewanvalentine.io/efficiency-with-go-channels/

import (
	"fmt"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {

	defer timeTrack(time.Now(), "Rules")

	finished := make(chan bool)

	go func() {
		ok := ruleOne()
		if ok {	finished <- true }
	}()

	go func() {
		ok := ruleTwo()
		if ok {	finished <- true }
	}()

	go func() {
		ok := ruleThree()
		if ok {	finished <- true }
	}()

	go func() {
		ok := ruleFour()
		if ok {	finished <- true }
	}()

	go func() {
		ok := ruleFive()
		if ok {	finished <- true }
	}()

	go func() {
		ok := ruleSix()
		if ok {	finished <- true }
	}()

	for {
		select {
		case isFinished := <-finished:
			fmt.Println("Done: ", isFinished)
			return
		}
	}
}

func ruleOne()   bool { time.Sleep(1000); return false; }
func ruleTwo()   bool { time.Sleep(1000); return false; }
func ruleThree() bool {	time.Sleep(1000); return false; }
func ruleFour()  bool { time.Sleep(1000); return false; }
func ruleFive()  bool { time.Sleep(1000); return false; }
func ruleSix()   bool { time.Sleep(1000); return true;  }
