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

	ok := ruleOne()
	if ok {	fmt.Println("Done!") }

	ok = ruleTwo()
	if ok {	fmt.Println("Done!") }

	ok = ruleThree()
	if ok {	fmt.Println("Done!") }

	ok = ruleFour()
	if ok {	fmt.Println("Done!") }

	ok = ruleFive()
	if ok {	fmt.Println("Done!") }

	ok = ruleSix()
	if ok {	fmt.Println("Done!") }
}

func ruleOne()   bool { time.Sleep(1000); return false; }
func ruleTwo()   bool { time.Sleep(1000); return false; }
func ruleThree() bool { time.Sleep(1000); return false; }
func ruleFour()  bool { time.Sleep(1000); return false; }
func ruleFive()  bool { time.Sleep(1000); return false; }
func ruleSix()   bool { time.Sleep(1000); return true; }
