// Example to parse querystring to struct
package main

import (
	"log"
	"net/url"

	"github.com/gorilla/schema"
)

type URLParams struct {
	Code  string `schema:"code"`
	State string `schema:"state"`
}

func main() {
	var (
		params  URLParams
		decoder = schema.NewDecoder()
	)
	p := "https://www.redirect-url.com?code=CODE&state=RANDOM_ID"

	u, _ := url.Parse(p)

	err := decoder.Decode(&params, u.Query())
	if err != nil {
		log.Println("Error in Decode parameters : ", err)
	} else {
		log.Printf("Decoded parameters : %#v\n", params)
	}
}
