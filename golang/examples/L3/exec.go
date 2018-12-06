package main

import ( 
	"fmt"
	"syscall"
	"os/exec"
	"os"
)

func main() {
	fmt.Println("hello world")
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	
	args := []string{"ls", "-a", "-l", "-l"}
	
	env := os.Environ()
	fmt.Println(env)
	
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}