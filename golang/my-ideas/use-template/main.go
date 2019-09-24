package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type PageData struct {
	Title string
	Data  map[string]interface{}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	Index()
	println("\n#####\n")
	About()
}

func Index() {
	var page = PageData{
		Title: "Index",
		Data:  make(map[string]interface{}),
	}
	page.Data["Message"] = "Welcome to Homepage"
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", page)
}

func About() {
	var page = PageData{
		Title: "About",
		Data:  make(map[string]interface{}),
	}
	page.Data["Message"] = "Introduce Ourself"
	tpl.ExecuteTemplate(os.Stdout, "about.gohtml", page)
}
