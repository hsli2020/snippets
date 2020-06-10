package main

import "fmt"

type Set struct {
	m map[string]bool
}

func NewSet() Set {
	m := make(map[string]bool)
	return Set{m: m}
}

func (s *Set) Contains(val string) bool {
	_, ok := s.m[val]
	return ok
}

func (s *Set) Add(val string) {
	s.m[val] = true
}

func (s *Set) Remove(val string) {
	delete(s.m, val)
}

func main() {
	s := NewSet()

	s.Add("foo")
	fmt.Printf("s has foo: %t. s has bar: %t\n", s.Contains("foo"), s.Contains("bar"))

	s.Remove("foo")
	fmt.Printf("s has foo: %t. s has bar: %t\n", s.Contains("foo"), s.Contains("bar"))
}
