package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// GET
		t, _ := template.ParseFiles("upload.gtpl")

		t.Execute(w, nil)

	} else if r.Method == "POST" {
		// Post
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)

	} else {
		fmt.Println("Unknown HTTP " + r.Method + "  Method")
	}
}

func main() {
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":9090", nil) // setting listening port
}

/*
<html>
<head>
<title>Test Upload a File</title>
</head>
<body>
  <form enctype="multipart/form-data" action="http://localhost:9090/upload" method="post">
    <input type="file" name="uploadfile" />
    <input type="submit" value="upload file" />
  </form>
</body>
</html>
*/
