package main

import (
    "fmt"
    "log"
    "path/filepath"
    "os"
    "time"
    "strings"
    "reflect"
    "errors"
/*
    "bufio"
    "bytes"
    "io"
    "math"
    "net"
    "net/http"
*/
)

func main() {
}

// array_keys() in php
func Keys(v interface{}) ([]string, error) {
    rv := reflect.ValueOf(v)
    if rv.Kind() != reflect.Map {
        return nil, errors.New("not a map")
    }
    t := rv.Type()
    if t.Key().Kind() != reflect.String {
        return nil, errors.New("not string key")
    }
    var result []string
    for _, kv := range rv.MapKeys() {
        result = append(result, kv.String())
    }
    return result, nil
}

func stringsDemo() {
    //pr := fmt.Println

    pr(strings.Title("helloGoLang"))
    pr(strings.Fields("Welcome To The Dollhouse!"))
    pr(strings.Fields("  foo bar  baz   "))
    pr(strings.EqualFold("Go", "go"))
    pr(strings.Count("cheese", "e"))

    pr(strings.Repeat("-", 80))
	pr(strings.Split("a man a plan a canal panama", "a "))
	pr(strings.Split(" xyz ", ""))
	pr(strings.Split("", "Bernardo O'Higgins"))

    pr(strings.Split("a,b,c", ","))
    pr(strings.SplitAfter("a,b,c", ","))

    rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}

	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

    r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
    pr(r.Replace("This is <b>HTML</b>!"))

    pr(strings.Fields("  foo bar  baz   "))
    pr(strings.Split("  foo bar  baz   ", " "))
}
