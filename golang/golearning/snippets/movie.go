package main

import (
	"encoding/json"
	"log"
	"os"
)

type ApiWrapper interface {
	Query(string) SearchResult
}

type Movie struct {
	Title string
	Link  string
	Price float64
}

type SearchResult struct {
	Service string
	Movie   Movie
}

type Youtube struct{}
type Amazon struct{}

func (y Youtube) Query(query string) SearchResult {
	// this is all done dynamically in the real implementation
	// and the implementation for each service (youtube in this case)
	// is in it's own file, but in the same package namespace as all
	// other wrappers

	// logging example
	log.Printf("Searching %s in %v\n", query, "youtube")
	movie := Movie{"Inception", "http://youtube.com/inception", 2.99}
	searchResult := SearchResult{"youtube", movie}
	return searchResult
}

func (a Amazon) Query(query string) SearchResult {
	// this is all done dynamically in the real implementation
	// and the implementation for each service (youtube in this case)
	// is in it's own file, but in the same package namespace as all
	// other wrappers

	// logging example
	log.Printf("Searching %s in %s\n", query, "amazon")
	movie := Movie{"Inception", "http://amazon.com/inception", 2.99}
	searchResult := SearchResult{"amazon", movie}
	return searchResult
}

func main() {
	resultsMap := make(map[string]Movie)
	movieQuery := "inception"

	wrappers := []ApiWrapper{new(Youtube), new(Amazon)}
	ch := make(chan SearchResult, len(wrappers))

	for _, wrapper := range wrappers {
		go func(wrapper ApiWrapper) { ch <- wrapper.Query(movieQuery) }(wrapper)
	}

	for _ = range wrappers {
		result := <-ch
		resultsMap[result.Service] = result.Movie
	}

	if err := json.NewEncoder(os.Stdout).Encode(&resultsMap); err != nil {
		log.Println(err)
	}
}
