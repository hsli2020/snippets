package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var godaemon = flag.Bool("d", false, "run app as a daemon with -d .")
var port = flag.String("p", ":8080", "set server port, -p=:8899 or -p :8899")
var close = flag.Bool("c", false, "close server")
var pid = flag.String("i", "0", "server pid")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *close {
		if *pid == "0" {
			fmt.Println("you must set -i, this is server pid")
			os.Exit(0)
		}
		cmd := exec.Command("kill", "-9", *pid)
		cmd.Start()
		os.Exit(0)
	}

	if *godaemon {
		args := os.Args[1:]
		for i := 0; i < len(args); i++ {
			if args[i] == "-d" {
				args = append(args[:i], args[i+1:]...)
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Printf("Server running...\nClose server: kill -9 %d\n", cmd.Process.Pid)
		fmt.Printf("http://localhost%s \n", *port)
		*godaemon = false
		os.Exit(0)
	}
}

func main() {
	fmt.Printf("http://localhost%s \n", *port)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	http.Handle("/", http.FileServer(http.Dir(dir)))
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
