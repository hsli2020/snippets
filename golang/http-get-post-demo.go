// Making a GET Request

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Print("Request Failed")
	}
	defer resp.Body.Close()

	// Log the request body
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Print(bodyString)
}

// Making a GET Request with Params

type Posts []struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	params := url.Values{}
	params.Add("userId", "1")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?" + params.Encode())
	if err != nil {
		log.Print("Request Failed")
	}
	defer resp.Body.Close()

	// Log the request body
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal result
	posts := Posts{}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Print(err)
	}
	log.Printf("The title of the first post is %s", posts[0].Title)
}

// Making a POST Request

type Post struct {
	Userid string `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	params := url.Values{}
	params.Add("title", "foo")
	params.Add("body", "bar")
	params.Add("userId", "1")

	resp, err := http.PostForm("https://jsonplaceholder.typicode.com/posts", params)
	if err != nil {
		log.Print("Request Failed")
	}
	defer resp.Body.Close()

	// Log the request body
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal result
	post := Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Print(err)
	}

	log.Printf("Post added with ID %d", post.ID)
}
