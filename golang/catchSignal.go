package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	log.SetOutput(os.Stdout)

	server := &http.Server{
		Addr: ":9000",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		}),
	}

	go catchSignal(server)

	log.Println("starting server")

	log.Println("Error from listen and serve", server.ListenAndServe()) //ListenAndServe will block the code here.

}

func catchSignal(server *http.Server) {

	terminateSignals := make(chan os.Signal, 1)
	reloadSignals := make(chan os.Signal, 1)

	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM) //NOTE:: syscall.SIGKILL we cannot catch kill -9 as its force kill signal.

	signal.Notify(reloadSignals, syscall.SIGUSR1)

	for { //We are looping here because config reload can happen multiple times.
		select {
		case s := <-terminateSignals:
			log.Println("Got one of stop signals, shutting down server gracefully, SIGNAL NAME :", s)
			log.Println("Error from shutdown", server.Shutdown(context.Background()))
			break //break is not necessary to add here as if server is closed our main function will end.
		case <-reloadSignals:
			log.Println("Got reload signal, will reload")
			//Some config reload code here.
		}
	}

}
