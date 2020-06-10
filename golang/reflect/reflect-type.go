package main

import "fmt"
import "log"
import "strings"
import "reflect"

func Log(format string, v ...interface{}) {
    format = strings.TrimSuffix(format, "\n") + "\n"
    log.Printf(format, v...)
}

var registry map[string]reflect.Type

type T struct {
	i int
}

func register(i interface{}) {
	registry[fmt.Sprint(reflect.TypeOf(i))] = reflect.TypeOf(i)
}

func main() {
    //Log("Hello, %s", "somebody")
    //Log("Hello, %s", "someone")

	registry = make(map[string]reflect.Type)
	register(T{})
	fmt.Printf("%#v\n", reflect.New(registry["main.T"]).Interface())
}
