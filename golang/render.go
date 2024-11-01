// ========== ./internal/render/render.go

package render

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/shapito27/go-web-app/internal/config"
	"github.com/shapito27/go-web-app/internal/models"
)

// data we pass to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

var functions = template.FuncMap{}

var appConfig *config.AppConfig

var pathToTemplate = "./templates"

// NewRenderer stores config for render package
func NewRenderer(ac *config.AppConfig) {
	appConfig = ac
}

// add default data to all template data
func addDefaultData(templateData *models.TemplateData, r *http.Request) *models.TemplateData {
	templateData.Flash = appConfig.Session.PopString(r.Context(), "flash")
	templateData.Error = appConfig.Session.PopString(r.Context(), "error")
	templateData.Warning = appConfig.Session.PopString(r.Context(), "warning")
	templateData.CSRFToken = nosurf.Token(r)

	return templateData
}

// Template to render templates
func Template(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData) error {
	var templates map[string]*template.Template

	if appConfig.UseCache {
		// getting templates list from config
		templates = appConfig.TemplatesCache
	} else {
		var err error
		templates, err = GetTemplatesCache()
		if err != nil {
			fmt.Println("Error getting template", err)
			return err
		}
	}

	t, ok := templates[tmpl+".page.tmpl"]
	if !ok {
		fmt.Println("Template doesn't exist")
		return errors.New("Template doesn't exist")
	}

	data = addDefaultData(data, r)

	err := t.Execute(w, data)

	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

//collect all templates then merge them with layout
func GetTemplatesCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplate))
	if err != nil {
		return myCache, err
	}

	matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplate))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplate))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

	// getting templates list
	templates, err := render.GetTemplatesCache()

	// pass config to render package
	render.NewRenderer(&app)

// Home page handler
func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home", &models.TemplateData{})
}

// About page handler
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about", &models.TemplateData{})
}

func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact", &models.TemplateData{})
}

func (rep *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
    // ...
	render.Template(w, r, "reservation-summary", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}
