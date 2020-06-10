package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := &Person{"Jonh", 27}
	buf, _ := xml.Marshal(person)
	body := bytes.NewBuffer(buf)

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// build a new request, but not doing the POST yet (strings.NewReader(body))
	req, err := http.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(body))
	if err != nil {	fmt.Println(err) }

	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	// now POST it
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {	fmt.Println(err) }

	response, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(response))
}

func MakeRequest() {
	message := map[string]interface{}{
		"hello":    "world",
		"life":     42,
		"embedded": map[string]string{"yes": "of course!"},
	}

	d, err := json.Marshal(message)
	if err != nil {	log.Fatalln(err) }

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(d))
	if err != nil {	log.Fatalln(err) }

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result)
}
