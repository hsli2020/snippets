package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT)
	for {
		sig := <-sc
		switch sig {
		case syscall.SIGINT:
			println("SIGINT")
			return
		}
	}
}
