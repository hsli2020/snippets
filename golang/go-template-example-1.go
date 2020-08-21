package main

import (
    rice "github.com/GeertJohan/go.rice"
    "gitlab.com/christianhellsten/go-utils/log"
    "html/template"
    "net/http"
    "os"
    "path/filepath"
    "strings"
)

type renderTemplateFunc func(w http.ResponseWriter, tmpl string, p interface{})

var renderTemplate renderTemplateFunc
var templates = template.New("").Funcs(templateMap)
var templateBox *rice.Box

func loadTemplates() {
    if config.debug {
        renderTemplate = renderTemplateDev
    } else {
        renderTemplate = renderTemplateProd
        newTemplate := func(path string, _ os.FileInfo, _ error) error {
            if path == "" {
                return nil
            }
            templateString, err := templateBox.String(path)
            if err != nil {
                log.Fatal("Unable to parse: path=%s, err=%s", path, err)
            }
            templates.New(filepath.Join("tmpl", path)).Parse(templateString)
            return nil
        }
        // Load and parse templates from binary or disk
        templateBox = rice.MustFindBox("tmpl")
        templateBox.Walk("", newTemplate)
    }
}

var (
    templateMap = template.FuncMap{
        "Upper": func(s string) string {
            return strings.ToUpper(s)
        },
    }
)

func renderTemplateProd(w http.ResponseWriter, tmpl string, p interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func renderTemplateDev(w http.ResponseWriter, tmpl string, p interface{}) {
    t, _ := template.ParseFiles(tmpl)
    t.Execute(w, p)
}