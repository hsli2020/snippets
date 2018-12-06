package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	fmt.Println("hello world")
	dateCmd := exec.Command("date")
	
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
	
	grepCmd := exec.Command("grep", "hello")
	
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	
	fmt.Println("> grep")
	fmt.Println(string(grepBytes))
	
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
