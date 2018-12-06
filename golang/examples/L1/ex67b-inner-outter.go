package main

// http://stackoverflow.com/questions/21251242/is-it-possible-to-call-overrided-method-from-parent-struct-in-golang

import "fmt"

type A struct {
}

func (a *A) Foo() {
    fmt.Println("A.Foo()")
}

func (a *A) Bar() {
    a.Foo()
}

type B struct {
    A
}

func (b *B) Foo() {
    fmt.Println("B.Foo()")
}

func main() {
    b := B{A: A{}}
    b.Bar()
}

//output: A.Foo()
