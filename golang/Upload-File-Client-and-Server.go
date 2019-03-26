// curl -X POST -H "Content-Type: application/octet-stream" --data-binary '@filename' http://127.0.0.1:5050/upload

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
	"io/ioutil"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("./result")
	if err != nil {
		panic(err)
	}
	n, err := io.Copy(file, r.Body)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(fmt.Sprintf("%d bytes are recieved.\n", n)))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	go func() {
		time.Sleep(time.Second * 1)
		upload()
	}()

	http.ListenAndServe(":5050", nil)
}

func upload() {
	file, err := os.Open("./filename")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post("http://127.0.0.1:5050/upload", "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}
