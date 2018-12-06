package view

import (
	"encoding/gob" "encoding/json" "fmt" "os" "strings"
	"html/template" "net/http" "net/url" "path/filepath"
)
/*
	"Template": {
		"Root": "base",
		"Children": [
			"partial/menu",
			"partial/footer"
		]
	},
	"View": {
		"BaseURI": "/",
		"Extension": "tmpl",
		"Folder": "template",
		"Name": "blank",
		"Caching": true
	}

	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
*/
var (
	childTemplates     []string
	rootTemplate       string

	templateCollection = make(map[string]*template.Template)
	pluginCollection   = make(template.FuncMap)

	viewInfo           View   // viewCfg  ViewConfig  / config Config
)

// View attributes
type View struct {
	BaseURI   string   // => config
	Extension string   // => config
	Folder    string   // => config
	Caching   bool     // => config

	Name      string
	Vars      map[string]interface{}

	request   *http.Request
}

// Configure sets the view information
func Configure(vi View) { viewInfo = vi }

// LoadTemplates will set the root and child templates
func LoadTemplates(rootTemp string, childTemps []string) {
	rootTemplate = rootTemp
	childTemplates = childTemps
}

// LoadPlugins will combine all template.FuncMaps into one map and then set the
// plugins for the templates. If a func already exists, it is rewritten, there is no error
func LoadPlugins(fms ...template.FuncMap) {
	// Final FuncMap
	fm := make(template.FuncMap)

	// Loop through the maps
	for _, m := range fms {
		// Loop through each key and value
		for k, v := range m {
			fm[k] = v
		}
	}

	// Load the plugins
	pluginCollection = fm
}

// PrependBaseURI prepends the base URI to the string
func (v *View) PrependBaseURI(s string) string {
	return v.BaseURI + s
}

// New returns a new view
func New(req *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})
	v.Vars["AuthLevel"] = "anon"

	v.BaseURI = viewInfo.BaseURI
	v.Extension = viewInfo.Extension
	v.Folder = viewInfo.Folder
	v.Name = viewInfo.Name

	// Make sure BaseURI is available in the templates
	v.Vars["BaseURI"] = v.BaseURI

	// This is required for the view to access the request
	v.request = req

	// Get session
	sess := session.Instance(v.request)

	// Set the AuthLevel to auth if the user is logged in
	if sess.Values["id"] != nil {
		v.Vars["AuthLevel"] = "auth"
	}

	return v
}

// AssetTimePath returns a URL with the proper base uri and timestamp appended.
// Works for CSS and JS assets. Determines if local or on the web
func (v *View) AssetTimePath(s string) (string, error) {
	if strings.HasPrefix(s, "//") {
		return s, nil
	}

	s = strings.TrimLeft(s, "/")
	abs, err := filepath.Abs(s)

	if err != nil {
		return "", err
	}

	time, err2 := FileTime(abs)
	if err2 != nil {
		return "", err2
	}

	return v.PrependBaseURI(s + "?" + time), nil
}

// RenderSingle renders a template to the writer (NO LAYOUT)
func (v *View) RenderSingle(w http.ResponseWriter) {  // RenderSimple

	// Get the plugin collection
	pc := pluginCollection

	templateList := []string{v.Name}

	// Loop through each template and test the full path
	for i, name := range templateList {
		// Get the absolute path of the root template
		path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
		if err != nil {
			http.Error(w, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		templateList[i] = path
	}

	// Determine if there is an error in the template syntax
	templates, err := template.New(v.Name).Funcs(pc).ParseFiles(templateList...)

	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tc := templates

	// Display the content to the screen
	err = tc.Funcs(pc).ExecuteTemplate(w, v.Name+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Render renders a template to the writer
func (v *View) Render(w http.ResponseWriter) {

	// Get the template collection from cache
	tc, ok := templateCollection[v.Name]

	// Get the plugin collection
	pc := pluginCollection

	// If the template collection is not cached or caching is disabled
	if !ok || !viewInfo.Caching {

		// List of template names
		var templateList []string
		templateList = append(templateList, rootTemplate)
		templateList = append(templateList, v.Name)
		templateList = append(templateList, childTemplates...)

		// Loop through each template and test the full path
		for i, name := range templateList {
			// Get the absolute path of the root template
			path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
			if err != nil {
				http.Error(w, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
				return
			}
			templateList[i] = path
		}

		// Determine if there is an error in the template syntax
		templates, err := template.New(v.Name).Funcs(pc).ParseFiles(templateList...)

		if err != nil {
			http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Cache the template collection
		templateCollection[v.Name] = templates

		// Save the template collection
		tc = templates
	}

	// Display the content to the screen
	err := tc.Funcs(pc).ExecuteTemplate(w, rootTemplate+"."+v.Extension, v.Vars)

	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}

// Validate returns true if all the required form values are passed
func Validate(req *http.Request, required []string) (bool, string) {
	for _, v := range required {
		if req.FormValue(v) == "" {
			return false, v
		}
	}

	return true, ""
}

// Repopulate updates the dst map so the form fields can be refilled
func Repopulate(list []string, src url.Values, dst map[string]interface{}) {
	for _, v := range list {
		dst[v] = src.Get(v)
	}
}

// FileTime returns the modification time of the file
func FileTime(name string) (string, error) {
	fi, err := os.Stat(name)
	if err != nil {
		return "", err
	}
	mtime := fi.ModTime().Unix()
	return fmt.Sprintf("%v", mtime), nil
}
