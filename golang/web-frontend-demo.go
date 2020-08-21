// ========== main.go
package main

import (
	"net/http"

	"github.com/philippta/web-frontend-demo/html"
)

func main() {
	http.HandleFunc("/dashboard", dashboard)
	http.HandleFunc("/profile/show", profileShow)
	http.HandleFunc("/profile/edit", profileEdit)
	http.ListenAndServe(":8080", nil)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	p := html.DashboardParams{
		Title:   "Dashboard",
		Message: "Hello from dashboard",
	}
	html.Dashboard(w, p)
}

func profileShow(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileShowParams{
		Title:   "Profile Show",
		Message: "Hello from profile show",
	}
	html.ProfileShow(w, p)
}

func profileEdit(w http.ResponseWriter, r *http.Request) {
	p := html.ProfileEditParams{
		Title:   "Profile Edit",
		Message: "Hello from profile edit",
	}
	html.ProfileEdit(w, p)
}

// ========== html/html.go
package html

import (
	"embed"
	"io"
	"text/template"
)

//go:embed *
var files embed.FS

var (
	dashboard   = parse("dashboard.html")
	profileShow = parse("profile/show.html")
	profileEdit = parse("profile/edit.html")
)

type DashboardParams struct {
	Title   string
	Message string
}

func Dashboard(w io.Writer, p DashboardParams) error {
	return dashboard.Execute(w, p)
}

type ProfileShowParams struct {
	Title   string
	Message string
}

func ProfileShow(w io.Writer, p ProfileShowParams) error {
	return profileShow.Execute(w, p)
}

type ProfileEditParams struct {
	Title   string
	Message string
}

func ProfileEdit(w io.Writer, p ProfileEditParams) error {
	return profileEdit.Execute(w, p)
}

func parse(file string) *template.Template {
	return template.Must(
		template.New("layout.html").ParseFS(files, "layout.html", file))
}

// ========== html/dashboard.html
{{define "content"}}
<p>
    Dashboard:
    {{.Message}}
</p>
{{end}}

// ========== html/layout.html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    {{block "content" .}}{{end}}
</body>
</html>

// ========== html/profile/edit.html
{{define "content"}}
<p>
    Profile Edit:
    {{.Message}}
</p>
{{end}}

// ========== html/profile/show.html
{{define "content"}}
<p>
    Profile Show:
    {{.Message}}
</p>
{{end}}
