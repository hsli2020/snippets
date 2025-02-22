package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.POST("/", indexSubmission)
	r.Handler("GET", "/favicon.ico", http.NotFoundHandler())
	r.ServeFiles("/public/*filepath", http.Dir("./public"))
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := getCookie(w, r)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:]) //only send over images
}

func indexSubmission(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := getCookie(w, r)

	mf, fh, err := r.FormFile("nf")
	if err != nil {
		fmt.Println(err)
	}
	defer mf.Close()
	// create sha for file name
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
	// create new file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "pics", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	// copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)
	// add filename to this user's cookie
	c = appendValue(w, c, fname)

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:]) //only send over images
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
/*
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>INDEX</title>
</head>
<body>

<h1>Your Pictures:</h1>
{{range .}}
<img src="/public/pics/{{.}}" width="400px">
{{end}}

<form method="post" enctype="multipart/form-data">
    <input type="file" name="nf">
    <input type="submit">
</form>

</body>
</html>
*/