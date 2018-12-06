package main

import (
    "errors"
    "fmt"
    "reflect"
)

func resolve(result interface{}) error {
    val := reflect.ValueOf(result)
    if val.Kind() != reflect.Ptr || val.IsNil() {
        return errors.New("result param for Lookup must be a pointer")
    }

    for i := 0; i < val.Elem().NumField(); i++ {
        name := val.Elem().Type().Field(i).Tag.Get("my_ref")
        value := val.Elem().Field(i)

        if name == "firstname" {
            value.SetString("John")
        } else if name == "lastname" {
            value.SetString("Doe")
        } else if name == "age" {
            value.SetInt(42)
        } else if name == "nicknames" {
            nicks := [...]string{"foo", "bar", "ping", "pong"}

            nicknames := reflect.MakeSlice(value.Type(), len(nicks), len(nicks))
            value.Set(nicknames)
            for i, nickname := range nicks {
                value.Index(i).SetString(nickname)
            }
        }
    }

    return nil
}

type My struct {
    Firstname string   `my_ref:"firstname"`
    Lastname  string   `my_ref:"lastname"`
    Age       int      `my_ref:"age"`
    Nicknames []string `my_ref:"nicknames"`
}

func main() {
    my := &My{}
    resolve(my)
    fmt.Printf("%#v", my)
}
