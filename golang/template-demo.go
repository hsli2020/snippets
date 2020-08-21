// https://github.com/fgm/pkger_demo/tree/main/v2
// A short demo using Go embed to load templates

package main

import (
	"embed"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"
)

// To embed a random file, just go:embed it as a string or []byte
//go:embed hello.txt
var hello string

// To use an embed hierarchy, use go:embed with an embed.FS.
//go:embed templates
var templates embed.FS

// PageData is an example template data structure
type PageData struct {
	Path string
	Year int
}

func compileTemplates(dir string) (*template.Template, error) {
	const fun = "compileTemplates"
	tpl := template.New("")
	// Since filepath.Walk only handles filesystem directories, we use the new
	// and optimized fs.WalkDir introduced in Go 1.16, which takes an fs.FS.
	err := fs.WalkDir(templates, dir, func(path string, info fs.DirEntry, err error) error {
		// Skip non-templates.
		if info.IsDir() || !strings.HasSuffix(path, ".gohtml") {
			return nil
		}
		// Load file from embed virtual file, or use the shortcut
		// templates.ReadFile(path).
		f, _ := templates.Open(path)
		// Now read it.
		sl, _ := io.ReadAll(f)
		// It can now be parsed as a string.
		tpl.Parse(string(sl))
		return nil
	})
	return tpl, err
}

func main() {
	const (
		addr = ":8080"
		dir  = "templates"
	)

	// Only compile templates on startup.
	tpl, _ := compileTemplates(dir)

	// Now serve pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := PageData{Path: r.URL.Path, Year: time.Now().Local().Year()}
		tpl.ExecuteTemplate(w, "page", ctx)
	})
	log.Printf("Listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// templates/page.gohtml

{{ define "page" }}
{{- /*gotype: code.osinet.fr/fgm/pkger_demo.PageData*/ -}}
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>{{ .Path }} callback</title>
</head>
<body>
  <p>Called on {{ .Path }}</p>
  {{ template "layout/footer" . }}
</body>
</html>
{{ end }}


// templates/layout/footer.gohtml


{{define "layout/footer"}}
  {{- /*gotype: github.com/fgm/pkger_demo/v2.PageData*/ -}}
  <footer>
    &copy; {{ .Year }} Frederic G. MARAND for OSInet
  </footer>
{{end}}
