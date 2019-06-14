package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// create data structures to hold JSON data this needs to match the fields in
// the JSON feed see sample data fetch at bottom of this file
// create an individual entry data structure
// this does not need to hold every field, just the ones we want
type Entry struct {
	Title     string
	Author    string
	URL       string
	Permalink string
}

// the feed is the full JSON data structure,this sets up the array of Entry types (defined above)
type Feed struct {
	Data struct {
		Children []struct {
			Data Entry
		}
	}
}

func main() {
	url := "https://www.reddit.com/r/golang/new.json"

	resp, err := http.Get(url) // fetch url
	if err != nil {
		log.Fatalln("Error fetching:", err)
	}
	defer resp.Body.Close() // defer response close

	if resp.StatusCode != http.StatusOK { // confirm we received an OK status
		log.Fatalln("Error Status not OK:", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body) // read the entire body of the response
	if err != nil {
		log.Fatalln("Error reading body:", err)
	}

	// create an empty instance of Feed struct, this is what gets filled in when unmarshaling JSON
	var entries Feed
	if err := json.Unmarshal(body, &entries); err != nil {
		log.Fatalln("Error decoing JSON", err)
	}

	for _, ed := range entries.Data.Children { // loop through the children and create entry objects
		entry := ed.Data
		log.Println(">>>")
		log.Println("Title   :", entry.Title)
		log.Println("Author  :", entry.Author)
		log.Println("URL     :", entry.URL)
		log.Printf("Comments: http://reddit.com%s \n", entry.Permalink)
	}
}
