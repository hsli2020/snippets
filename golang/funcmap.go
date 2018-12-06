package main

import (
    "fmt"
    "errors"
    "reflect"
)

var (
    ErrParamsNotAdapted = errors.New("The number of params is not adapted.")
)

type Funcs map[string]reflect.Value

func NewFuncs(size int) Funcs {
    return make(Funcs, size)
}

func (f Funcs) Bind(name string, fn interface{}) (err error) {
    defer func() {
        if e := recover(); e != nil {
            err = errors.New(name + " is not callable.")
        }
    }()
    v := reflect.ValueOf(fn)
    v.Type().NumIn()
    f[name] = v
    return
}

func (f Funcs) Call(name string, params ... interface{}) (result []reflect.Value, err error) {
    if _, ok := f[name]; !ok {
        err = errors.New(name + " does not exist.")
        return
    }
    if len(params) != f[name].Type().NumIn() {
        err = ErrParamsNotAdapted
        return
    }
    in := make([]reflect.Value, len(params))
    for k, param := range params {
        in[k] = reflect.ValueOf(param)
    }
    result = f[name].Call(in)
    return
}

var (
    testcases = map[string]interface{}{
        "hello": func() {print("hello")},
        "foobar": func(a, b, c int) int {return a+b+c},
        "errstring": "Can not call this as a function",
        "errnumeric": 123456789,
    }
    funcs = NewFuncs(100)
)

func TestBind() {
    for k, v := range testcases {
        err := funcs.Bind(k, v)
        if k[:3] == "err" {
            if err == nil {
                fmt.Printf("Bind %s: %s\n", k, "an error should be paniced.")
            }
        } else {
            if err != nil {
                fmt.Printf("Bind %s: %s\n", k, err)
            }
        }
    }
}

func TestCall() {
    if _, err := funcs.Call("foobar"); err == nil {
        fmt.Printf("Call %s: %s\n", "foobar", "an error should be paniced.")
    }
    if _, err := funcs.Call("foobar", 0, 1, 2); err != nil {
        fmt.Printf("Call %s: %s\n", "foobar", err)
    }
    if _, err := funcs.Call("errstring", 0, 1, 2); err == nil {
        fmt.Printf("Call %s: %s\n", "errstring", "an error should be paniced.")
    }
}

func main() {
	TestBind()
	TestCall()
}
