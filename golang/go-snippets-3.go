package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializes given interface into JSON string
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializes JSON string object to given interface
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/saurabmish/Coffee-Shop/data"
	"github.com/saurabmish/Coffee-Shop/handlers"
)

func main() {
	l := log.New(os.Stdout, "Coffee shop API service ", log.LstdFlags)
	v := data.NewValidation()

	coffeeHandler := handlers.NewProducts(l, v)
	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/coffee/get/all", coffeeHandler.RetrieveAll)
	getRouter.HandleFunc("/coffee/get/{id:[0-9]+}", coffeeHandler.RetrieveSingle)

	putRouter := serveMux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/coffee/modify/{id:[0-9]+}", coffeeHandler.Modify)
	putRouter.Use(coffeeHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/coffee/add", coffeeHandler.Add)
	postRouter.Use(coffeeHandler.MiddlewareProductValidation)

	deleteRouter := serveMux.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/coffee/remove/{id:[0-9]+}", coffeeHandler.Remove)

	// reliability pattern for server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// ensure that service will not be blocked
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// ensure graceful shutdown of server
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	signal := <-signalChannel
	l.Println("Received signal for graceful shutdown", signal)
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

