// stack.go
// Package collection implements a generic stack.
package collection

// The zero value for Stack is an empty stack ready to use.
type Stack struct {
    data []interface{}
}

// Push adds x to the top of the stack.
func (s *Stack) Push(x interface{}) {
    s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// It’s a run-time error to call Pop on an empty stack.
func (s *Stack) Pop() interface{} {
    i := len(s.data) - 1
    res := s.data[i]
    s.data[i] = nil  // to avoid memory leak
    s.data = s.data[:i]
    return res
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
    return len(s.data)
}

// example_test.go
package collection_test

import (
    collection "."
    "fmt"
)

func Example() {
    var s collection.Stack
    s.Push("world")
    s.Push("hello, ")
    for s.Size() > 0 {
        fmt.Print(s.Pop())
    }
    fmt.Println()
    // Output: hello, world
}

/*
Go 	                        Approximate Java equivalent
var v1 int 	                int v1 = 0;
var v2 *int 	            Integer v2 = null;
var v3 string 	            String v3 = "";
var v4 [10]int 	            int[] v4 = new int[10]; // v4 is a value in Go.
var v5 []int 	            int[] v5 = null;
var v6 *struct { a int }   	C v6 = null; // Given: class C { int a; }
var v7 map[string]int 	    HashMap<String,Integer> v7 = null;
var v8 func(a int) int 	    F v8 = null; // interface F { int f(int a); }
*/

// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
    ch := make(chan struct{})
    go func() {
        time.Sleep(delay)
        fmt.Println(text)
        close(ch)
    }()
    return ch
}

wait := Publish("important news", 2 * time.Minute)
// Do some more work.
<-wait // blocks until the text has been published

