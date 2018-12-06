package main

// http://stackoverflow.com/questions/21251242/is-it-possible-to-call-overrided-method-from-parent-struct-in-golang

import "fmt"

type I interface {
    Foo()
}

type A struct {
    i I
}

func (a *A) Foo() {
    fmt.Println("A.Foo()")
}

func (a *A) Bar() {
    a.i.Foo()
}

type B struct {
    A
}

func (b *B) Foo() {
    fmt.Println("B.Foo()")
}

func main() {
    b := B{A: A{}}
    b.i = &b     // here i works like an attribute of b
    b.Bar()
}

//output: B.Foo()
