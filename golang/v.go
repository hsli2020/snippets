package main

import "fmt"
import "strings"

func main() {
	var vals = []interface{}{ "abc", 123, 45.67 }

	s := ""
	for _, v := range vals {
		s += fmt.Sprintf("%v", v)  // convert any type to string
	}
	fmt.Println(s)

	//a := make([]string, 0)
	var a []string
	for _, v := range vals {
		a = append(a, fmt.Sprintf("%v", v))  // convert any type to string
	}
    fmt.Println(strings.Join(a, "-"))

	var arr = []int{ 10, 20, 30, 1, 2, 3 }
	var over10 = func(v int) bool { return v >= 10 }
	big := Filter(arr, over10)
	fmt.Println(big)
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
    var p []int // == nil
    for _, v := range s {
        if fn(v) {
            p = append(p, v)
        }
    }
    return p
}
