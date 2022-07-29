// https://github.com/bugbytes-io/htmx-go-demo

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("Go app...")

	// handler function #1 - returns the index.html template, with film data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	// handler function #2 - returns the template block with the newly added film, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		// tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
	}

	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTMX & Go - Demo</title>
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-KK94CHFLLe+nY2dmCWGMq91rCGa5gtU4mk92HdvYe+M/SXH301p5ILy+dN9+nJOZ" crossorigin="anonymous">
    <script src="https://unpkg.com/htmx.org@1.9.2" integrity="sha384-L6OqL9pRWyyFU3+/bjdSri+iIphTN/bvYyM37tICVyOJkWZLpP2vGn6VUEXgzg6h" crossorigin="anonymous"></script>
</head>

<body class="container">

    <div class="row mt-4 g-4">
        <div class="col-8">
            <h1 class="mb-4">Film List</h1>
        
            <ul class="list-group fs-5 me-5" id="film-list">
                {{ range .Films }}
                {{ block "film-list-element" .}}
                    <li class="list-group-item bg-primary text-white">{{ .Title }} - {{ .Director }}</li>
                {{ end }}
                {{ end }}
            </ul>
        </div>

        <div class="col-4">
            <h1 class="mb-4">Add Film</h1>

            <form hx-post="/add-film/" hx-target="#film-list" hx-swap="beforeend" hx-indicator="#spinner">
                <div class="mb-2">
                    <label for="film-title">Title</label>
                    <input type="text" name="title" id="film-title" class="form-control" />
                </div>
                <div class="mb-3">
                    <label for="film-director">Director</label>
                    <input type="text" name="director" id="film-director" class="form-control" />
                </div>

                <button type="submit" class="btn btn-primary">
                    <span class="spinner-border spinner-border-sm htmx-indicator" id="spinner" role="status" aria-hidden="true"></span>
                    Submit
                </button>
            </form>
        </div>

    </div>

</body>
</html>
