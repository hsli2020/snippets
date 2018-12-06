package main

import (
    "encoding/json"
    "fmt"
)

var my_json string = `{
    "an_array":[
		"with_a string",
		{
			"and":"some_more",
			"different":["nested", "types"]
		}
    ]
}`

func WTHisThisJSON(f interface{}) {
    switch vf := f.(type) {
		case map[string]interface{}:
			fmt.Println("is a map:")
			for k, v := range vf {
				switch vv := v.(type) {
				case string:
					fmt.Printf("%v: is string - %q\n", k, vv)
				case int:
					fmt.Printf("%v: is int - %q\n", k, vv)
				default:
					fmt.Printf("%v: ", k)
					WTHisThisJSON(v)
				}
			}
		case []interface{}:
			fmt.Println("is an array:")
			for k, v := range vf {
				switch vv := v.(type) {
				case string:
					fmt.Printf("%v: is string - %q\n", k, vv)
				case int:
					fmt.Printf("%v: is int - %q\n", k, vv)
				default:
					fmt.Printf("%v: ", k)
					WTHisThisJSON(v)
				}
			}
    }
}

func main() {
    fmt.Println("JSON:\n", my_json, "\n")

    var f interface{}
    err := json.Unmarshal([]byte(my_json), &f)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("JSON: ")
        WTHisThisJSON(f)
    }
}
