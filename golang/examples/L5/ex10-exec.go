package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, _ := exec.Command("ls", "-la").CombinedOutput()
	fmt.Println(string(out))
}

package main

import (
	"os/exec"
	"log"
	"bytes"
	"fmt"
)

func main() {
	cmd := exec.Command("ls", "-l")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}

package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// start
	cmd := exec.Command("sleep", "5")
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// wait or timeout
	donec := make(chan error, 1)
	go func() {
		donec <- cmd.Wait()
	}()
	select {
	case <-time.After(3 * time.Second):
		cmd.Process.Kill()
		fmt.Println("timeout")
	case <-donec:
		fmt.Println("done")
	}
}
