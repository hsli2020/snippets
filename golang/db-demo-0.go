package main  // main.go

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Comment struct {
	User string
	Text string
}

func main() {
	db, err := sql.Open("postgres", "user=dbdemo password=dbdemo dbname=dbdemo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS Comments (
			"User"    TEXT,
			"Comment" TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			ShowErrorPage(w, http.StatusMethodNotAllowed, "Invalid method", nil)
			return
		}

		rows, err := db.Query(`SELECT "User", "Comment" FROM Comments`)
		if err != nil {
			ShowErrorPage(w, http.StatusInternalServerError, "Unable to access DB", err)
			return
		}

		comments := []Comment{}
		for rows.Next() {
			var comment Comment
			err := rows.Scan(&comment.User, &comment.Text)
			if err != nil {
				ShowErrorPage(w, http.StatusInternalServerError, "Unable to load data", err)
				return
			}
			comments = append(comments, comment)
		}

		if err := rows.Err(); err != nil {
			ShowErrorPage(w, http.StatusInternalServerError, "Failed to load data from DB", err)
			return
		}

		ShowCommentsPage(w, comments)
	})

	http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			ShowErrorPage(w, http.StatusMethodNotAllowed, "Invalid method", nil)
			return
		}

		if err := r.ParseForm(); err != nil {
			ShowErrorPage(w, http.StatusBadRequest, "Unable to parse data", err)
			return
		}

		user := r.Form.Get("user")
		comment := r.Form.Get("comment")

		_, err = db.Exec(`INSERT INTO Comments ("User", "Comment") VALUES ($1, $2)`, user, comment)
		if err != nil {
			ShowErrorPage(w, http.StatusInternalServerError, "Unable to add data", err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	log.Println("Started listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


package main  // present.go

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func ShowCommentsPage(w http.ResponseWriter, comments []Comment) {
	w.WriteHeader(http.StatusOK)
	terr := commentsTemplate.Execute(w, map[string]interface{}{
		"Comments": comments,
	})
	if terr != nil {
		log.Println(terr)
	}
}

func ShowErrorPage(w http.ResponseWriter, statuscode int, title string, err error) {
	w.WriteHeader(statuscode)
	terr := errorTemplate.Execute(w, map[string]interface{}{
		"Title": title,
		"Error": err,
	})
	if terr != nil {
		log.Println(terr)
	}
}

var commentsTemplate = template.Must(template.New(``).Parse(`
<!DOCTYPE html>
<html><body>
	<h1>Comment Site</h1>
	<form action="/comment" method="POST">
		<input type="text" name="user"/>
		<input type="text" name="comment"/>
		<input type="submit" />
	</form>
	<h2>Comments</h2>
	<div>
	{{ range .Comments }}
		<div class="comment">
			{{ .User }} - {{ .Text }}
		</div>
	{{ end }}
	</div>
</body></html>
`))

var errorTemplate = template.Must(template.New(``).Parse(`
<!DOCTYPE html>
<html>
<head>
	{{if .Redirect}}<meta http-equiv="refresh" content="0; url={{.Redirect}}">{{end}}
</head>
<body>
	<h1>{{.Title}}</h1>
	{{if .Error}}<p>{{.Error}}</p>{{end}}
</body>
</html>
`))
