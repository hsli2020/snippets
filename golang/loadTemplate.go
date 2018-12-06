var (
   templates *template.Template
)

func loadTemplate() {
    funcMap := template.FuncMap{
        "safe":func(s string) template.HTML {
            return template.HTML(s)
        },
    }
    var err error
    templates, err = utils.BuildTemplate("/theme/path/", funcMap)
    if err != nil {
        log.Printf("Can't read template file %v,", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
        //lookup the theme your want to use
    templ = templates.Lookup("theme.html")
    err := templ.Execute(w, data)
    if err != nil {
        log.Println(err)
    }
 }

func main() {
   loadTemplate()
}

func BuildTemplate(dir string, funcMap template.FuncMap) (*template.Template, error) {
    fs, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Printf("Can't read template folder: %s\n", dir)
        return nil, err
    }
    files := make([]string, len(fs))
    for i, f := range (fs) {
        files[i] = path.Join(dir, f.Name())
    }
    return template.Must(template.New("Template").Funcs(funcMap).ParseFiles(files...)), nil
}

////////////////////////////////////////////////////////////////////////////////
package main

import (
    "fmt"
    "html/template"
    "os"
    "path/filepath"
)

func consumer(p string, i os.FileInfo, e error) error {
    t := template.New(p)
    fmt.Println(t.Name())
    return nil
}

func main() {
    filepath.Walk("/path/to/template/root", filepath.WalkFunc(consumer))
}
////////////////////////////////////////////////////////////////////////////////
