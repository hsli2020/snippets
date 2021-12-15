package main

import (
	"fmt"
	"runtime"
)

// https://github.com/bigwhite/experiments/tree/master/trace-function-call-chain
// github.com/bigwhite/experiments/blob/master/trace-function-call-chain/trace1/trace.go

func trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	fmt.Printf("enter: %s\n", name)
	return func() { fmt.Printf("exit: %s\n", name) }
}

// github.com/bigwhite/experiments/blob/master/trace-function-call-chain/trace1/main.go
func A1() {
	defer trace()()
	B1()
}

func B1() {
	defer trace()()
	C1()
}

func C1() {
	defer trace()()
	D()
}

func D() {
	defer trace()()
}

func main() {
	A1()
}
