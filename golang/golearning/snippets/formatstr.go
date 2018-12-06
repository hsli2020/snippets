package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func WithMap(format string, m map[string]string) (string, error) {
	tpl, err := template.New("format").Parse(format)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, m)
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

func main() {
	m := map[string]string{
		"data": "world",
		"end":  "!",
	}
	f, err := WithMap("Hello {{.data}}{{.end}}", m)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}

package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func WithMap(format string, m map[string]string) (string, error) {
	tpl := template.New("format").Funcs(template.FuncMap{
		"S": func(key string) string {
			if val, ok := m[key]; ok {
				return val
			}
			return key // could format this however you want
		},
	})
	tpl, err := tpl.Parse(format)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, m)
	if err != nil {
		return "", err
	}
	return buf.String(), nil

}

func main() {
	m := map[string]string{
		"data": "world",
		"end":  "!",
	}
	f, err := WithMap(`Hello {{S "data"}}{{S "end"}} {{S "missing"}}`, m)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}

package main

import "fmt"

func main() {
	m := map[string]string{
		"data": "world",
		"end":  "!",
	}
	fmt.Printf(`Hello %[2]s %[1]s %[3]s`, m["data"], m["end"], m["huh"]) 
}

package main

import (
    "fmt"
)

func main() {
    msg := "%s has %d messages."

    name := "Steve"
    num := 50

    result := fmt.Sprintf(msg, name, num)

    fmt.Println(result)
}

// To substitute more variables ( higher than 3), it is recommended to use 
// the text/template package.

package main

import (
    "fmt"
    "os"
    "text/template"
)

type Data struct {
    Name string // has to be uppercase/exportable/public
    Num  uint   // for the interpolation/substitution to work
}

func main() {
    msg := "{{.Name}} has {{.Num}} messages.\n"

    substitute := Data{"Steve", 50}

    tmpl, err := template.New("msg").Parse(msg)

    err = tmpl.Execute(os.Stdout, substitute)

    if err != nil {
        fmt.Println(err)
    }
}

package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {
	data := map[string]interface{}{
		"Name":     "Bob",
		"UserName": "bob92",
		"Roles":    []string{"dbteam", "uiteam", "tester"},
	}

	t := template.Must(template.New("email").Parse(emailTmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}
	s := buf.String()
	fmt.Println(s)

}

const emailTmpl = `Hi {{.Name}}!

Your account is ready, your user name is: {{.UserName}}

You have the following roles assigned:
{{range $i, $r := .Roles}}{{if ne $i 0}}, {{end}}{{.}}{{end}}
`
