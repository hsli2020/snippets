package main  // https://gist.github.com/System-Glitch/301e95975a2645b8ea57c47b0c7cfef4

// This is an example of a resilient service worker program written in Go.
//
// This program will run a worker, wait 5 seconds, and run it again.
// It exits when SIGINT or SIGTERM is received, while ensuring any ongoing work
// is finished before exiting.
//
// Unexpected panics are also handled: program won't crash if the worker panics.
// However, panics in goroutines started by the worker won't be handled and have
// to be dealt with manually.

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"
	"time"
)

var (
	sigChan     chan os.Signal // Used for shutdown
	workerChan  chan bool      // Used for worker return value (can be of any type)
	triggerChan chan bool      // Used to trigger a new worker run
	waitGroup   sync.WaitGroup
	errLogger   *log.Logger
)

const sleepDuration time.Duration = time.Duration(5) * time.Second

func runWorker() {
	defer recoverWorker()

	// Fake long process
	time.Sleep(sleepDuration)

	waitGroup.Done()
	workerChan <- true
}

func recoverWorker() {
	if err := recover(); err != nil {
		// Handle unexpected panic
		errLogger.Println(err)
		errLogger.Print(string(debug.Stack()))

		// Finish worker execution anyway
		waitGroup.Done()

		// Return false if service should stop on panic
		// Service will continue otherwise
		workerChan <- true
	}
}

func runTimer() {
	for {
		triggerChan <- true
		if !<-workerChan { // Exit if worker returned false
			sigChan <- syscall.SIGTERM
			return
		}
		time.Sleep(sleepDuration)
	}
}

func listen() {
	for {
		select {
		case <-sigChan:
			// Wait for worker to finish before exit
			waitGroup.Wait()
			return
		case <-triggerChan:
			waitGroup.Add(1)
			go runWorker()
		}
	}
}

func setup() {
	sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	workerChan = make(chan bool)
	triggerChan = make(chan bool)

	errLogger = log.New(os.Stderr, "", log.LstdFlags)
}

func main() {
	setup()

	go runTimer()
	fmt.Println("Service running...")
	listen()
	fmt.Println("Service stopped.")
}
