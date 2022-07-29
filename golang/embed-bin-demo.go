// https://www.abilityrush.com/how-to-embed-any-file-in-golang-binary/
package main

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//go:embed test
var s []byte

//go:embed images/golang.png
var embededImage []byte

//go:embed images/*
var imageList embed.FS

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write(s)
	})

	http.HandleFunc("/image", func(rw http.ResponseWriter, r *http.Request) {

		imageData, err := ioutil.ReadFile("images/golang.png")
		if err != nil {
			rw.Write([]byte("Some error occured - " + err.Error()))
			return
		}

		rw.Header().Write(bytes.NewBufferString("Content-Type: image/png"))
		rw.Header().Write(bytes.NewBufferString("Content-Length: " + strconv.Itoa(len(imageData))))
		rw.Write(imageData)
	})

	http.HandleFunc("/image_embeded", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Write(bytes.NewBufferString("Content-Type: image/png"))
		rw.Header().Write(bytes.NewBufferString("Content-Length: " + strconv.Itoa(len(embededImage))))
		rw.Write(embededImage)
	})

	http.HandleFunc("/get_image", func(rw http.ResponseWriter, r *http.Request) {

		fname := r.URL.Query().Get("name")

		data, err := imageList.ReadFile("images/" + fname)
		if err != nil {
			rw.Write([]byte("Error occured : " + err.Error()))
			return
		}

		rw.Header().Write(bytes.NewBufferString("Content-Type: image/png"))
		rw.Header().Write(bytes.NewBufferString("Content-Length: " + strconv.Itoa(len(data))))
		rw.Write(data)
	})

	fmt.Println("started")

	http.ListenAndServe(":7000", nil)
}
