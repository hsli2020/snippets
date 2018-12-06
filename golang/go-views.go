package views

import (
	"html/template"
	"net/http"
)

func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

type ViewData struct {
	Flashes map[string]string
	Data    interface{}
}

var flashRotator int = 0

func flashes() map[string]string {
	flashRotator = flashRotator + 1
	if flashRotator%3 == 0 {
		return map[string]string{
			"warning": "You are about to exceed your plan limts!",
		}
	} else {
		return map[string]string{}
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	vd := ViewData{
		Flashes: flashes(),
		Data:    data,
	}
	return v.Template.ExecuteTemplate(w, v.Layout, vd)
}

//func (v *View) Render(w http.ResponseWriter, data interface{}) error {
//	return v.Template.ExecuteTemplate(w, v.Layout, data)
//}
