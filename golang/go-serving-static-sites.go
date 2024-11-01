package main

// http://www.alexedwards.net/blog/serving-static-sites-with-go

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			errorHandler(w, r, 404)
			//http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		//http.Error(w, http.StatusText(500), 500)
		errorHandler(w, r, 500)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		//http.Error(w, http.StatusText(500), 500)
		errorHandler(w, r, 500)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    //w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/html") // why do I have to do this?
    if status == http.StatusNotFound {
        //fmt.Fprint(w, "custom 404")

		fp := filepath.Join("templates", "error404.html")

		tmpl, _ := template.ParseFiles(fp)
		tmpl.Execute(w, nil)
    }

    if status == 500 {
		fp := filepath.Join("templates", "error500.html")

		tmpl, _ := template.ParseFiles(fp)
		tmpl.Execute(w, nil)
    }
}
