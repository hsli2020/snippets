package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var usage = `Usage: %s [options] 
Options are:
    -n requests     Number of requests to perform
    -c concurrency  Number of multiple requests to make at a time
    -s timeout      Seconds to max. wait for each response
    -m method       Method name
`

var (
	requests    int
	concurrency int
	timeout     int
	method      string
	url         string
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}

	flag.IntVar(&requests, "n", 1000, "")
	flag.IntVar(&concurrency, "c", 100, "")
	flag.IntVar(&timeout, "s", 10, "")
	flag.StringVar(&method, "m", "GET", "")
	flag.Parse()

	if flag.NArg() != 1 {
		exit("Invalid url.")
	}

	method = strings.ToUpper(method)
	url = flag.Args()[0]

	if method != "GET" {
		exit("Invalid method.")
	}

	if requests < 1 || concurrency < 1 {
		exit("-n and -c cannot be smaller than 1.")
	}

	if requests < concurrency {
		exit("-n cannot be less than -c.")
	}

	w := Work{
		Requests:    requests,
		Concurrency: concurrency,
		Timeout:     timeout,
		Method:      method,
		Url:         url,
	}

	w.Run()
}

func exit(msg string) {
	flag.Usage()
	fmt.Fprintln(os.Stderr, "\n[Error] "+msg)
	os.Exit(1)
}

type Work struct {
	Requests    int
	Concurrency int
	Timeout     int
	Method      string
	Url         string
	results     chan *Result
	start       time.Time
	end         time.Time
}

type Result struct {
	Duration time.Duration
}

func (w *Work) Run() {
	w.results = make(chan *Result, w.Requests)
	w.start = time.Now()
	w.runWorkers()
	w.end = time.Now()

	w.print()
}

func (w *Work) runWorkers() {
	var wg sync.WaitGroup

	wg.Add(w.Concurrency)

	for i := 0; i < w.Concurrency; i++ {
		go func() {
			defer wg.Done()
			w.runWorker(w.Requests / w.Concurrency)
		}()
	}

	wg.Wait()
	close(w.results)
}

func (w *Work) runWorker(num int) {
	client := &http.Client{
		Timeout: time.Duration(w.Timeout) * time.Second,
	}

	for i := 0; i < num; i++ {
		w.sendRequest(client)
	}
}

func (w *Work) sendRequest(client *http.Client) {
	req, err := http.NewRequest(w.Method, w.Url, nil)

	if err != nil {
		log.Fatal(err.Error())
	}

	start := time.Now()
	client.Do(req)
	end := time.Now()

	w.results <- &Result{
		Duration: end.Sub(start),
	}
}

func (w *Work) print() {
	sum := 0.0
	num := float64(len(w.results))

	for result := range w.results {
		sum += result.Duration.Seconds()
	}

	rps := int(num / w.end.Sub(w.start).Seconds())
	tpr := sum / num * 1000

	fmt.Printf("Requests per second:\t%d [#/sec]\n", rps)
	fmt.Printf("Time per request:\t%.3f [ms]\n", tpr)
}
