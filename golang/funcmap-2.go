// https://gist.github.com/alex-leonhardt/8ed3f78545706d89d466434fb6870023
// golang text/template with a map[string]interface{} populated from mixed json data
package main

import (
	"encoding/json"
	"os"
	"reflect"
	"text/template"
)

func main() {
	var err error
	text := `
{
	"msg": "hello world",
	"msgs": ["hello", "gophers", "..."],
	"msgNum": [0, 1, 2, 3, 4, 5],
	"nested": [
		{"msg": "why"},
		{"msg": "did"},
		{"msg": "the"},
		{"msg": [0, 1, 1, 3, 2]},
		{"msg": "chicken"},
		{"msg": "cross"},
		99,
		["yolo", "yolo", "yolo"]
	]
}`
	m := make(map[string]interface{})
	if err = json.Unmarshal([]byte(text), &m); err != nil {
		panic(err)
	}

	tmpl := `
{{ range $k, $v := $.msgs }}Key:{{ $k }}, Value:{{ $v }}
{{ end }}
{{ range $_, $v := $.msgNum }}Values: {{ $v }}
{{ end }}
{{ $.nested }}
{{ range $_, $v := $.nested }}
	{{ if isInt $v }}
	v is int .. {{ $v }}
	{{ end }}
	{{- if isMap $v -}}
	{{- range $k, $v := $v -}}
		k={{ $k }}, v={{ $v }}
	{{- end -}}
	{{- end -}}
	{{- if isSlice $v -}}
		{{ range $_, $s := $v -}}
			{{ $s }}
		{{- end }}
	{{- end -}}
{{ end }}
`
  // https://www.calhoun.io/intro-to-templates-p3-functions/
  // https://blog.golang.org/laws-of-reflection
  // https://golang.org/pkg/reflect/
  // https://golang.org/pkg/text/template/#FuncMap
  tf := template.FuncMap{
		"isInt": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
				return true
			default:
				return false
			}
		},
		"isString": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.String:
				return true
			default:
				return false
			}
		},
		"isSlice": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Slice:
				return true
			default:
				return false
			}
		},
		"isArray": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Array:
				return true
			default:
				return false
			}
		},
		"isMap": func(i interface{}) bool {
			v := reflect.ValueOf(i)
			switch v.Kind() {
			case reflect.Map:
				return true
			default:
				return false
			}
		},
	}
	t := template.New("hello").Funcs(tf)
	tt, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	if err = tt.Execute(os.Stdout, &m); err != nil {
		panic(err)
	}
}

// @duzun commented on Apr 19
// Cool! I've extended a little bit:

func IsList(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	return v == reflect.Array || v == reflect.Slice
}

func IsNumber(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	switch v {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func IsInt(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	switch v {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint32, reflect.Uint64:
		return true
	default:
		return false
	}
}

func IsFloat(i interface{}) bool {
	v := reflect.ValueOf(i).Kind()
	return v == reflect.Float32 || v == reflect.Float64
}