//# First, we define base.tmpl:

{{ define "base" }}
<html>
<head>
    {{ template "title" . }}
</head>
<body>
    {{ template "scripts" . }}
    {{ template "sidebar" . }}
    {{ template "content" . }}
    <footer>
    ...
    </footer>
</body>
</html>
{{ end }}

// We define empty blocks for optional content so we don't have to define
// a block in child templates that don't need them
{{ define "scripts" }}{{ end }}
{{ define "sidebar" }}{{ end }}

//# And index.tmpl, which effectively extends our base template.

{{ define "title"}}<title>Index Page</title>{{ end }}

// Notice the lack of the script block - we don't need it here.

{{ define "sidebar" }}
    // We have a two part sidebar that changes depending on the page
    {{ template "sidebar_index" }} 
    {{ template "sidebar_base" }}
{{ end }}

{{ define "content" }}
    {{ template "listings_table" . }}
{{ end }}

import (
    "fmt"
    "html/template"
    "net/http"
    "path/filepath"
)

var templates map[string]*template.Template

// Load templates on program initialisation
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templatesDir := config.Templates.Path

	layouts, err := filepath.Glob(templatesDir + "layouts/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

    // Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = template.Must(template.ParseFiles(files...))
	}
}

// renderTemplate is a wrapper around template.ExecuteTemplate.
func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	return tmpl.ExecuteTemplate(w, "base.tmpl", data)
}

// Cache the template to improve performance
import (
    "fmt"
    "html/template"
    "net/http"

    "github.com/oxtoacart/bpool"
)

var bufpool *bpool.BufferPool

// renderTemplate is a wrapper around template.ExecuteTemplate.
// It writes into a bytes.Buffer before writing to the http.ResponseWriter to catch
// any errors resulting from populating the template.
func renderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	// Create a buffer to temporarily write to and check if any errors were encounted.
	buf := bufpool.Get()
	defer bufpool.Put(buf)
    
	err := tmpl.ExecuteTemplate(buf, "base.tmpl", data)
	if err != nil {
		return err
	}

	// Set the header and write the buffer to the http.ResponseWriter
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}

func init() {
	...
	bufpool = bpool.NewBufferPool(64)
	...
}
