package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	username := "adamng"
	password := "theultimatepassword"

	loginURL := "http://localhost:8080/login"

	urlData := url.Values{}
	urlData.Set("Username", username)
	urlData.Set("Password", password)

	resp, err := http.PostForm(loginURL, urlData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Status : ", resp.Status)
	fmt.Println("Header: ", resp.Header)
	fmt.Println("Body: ", resp.Body)
}
