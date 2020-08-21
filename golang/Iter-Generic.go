package main

import (
	"constraints"
	"fmt"
	"strings"
)

type Iter[T any] struct {
	next func() *T
}

func IterSlice[T any](slice []T) *Iter[T] {
	next := func() *T {
		if len(slice) < 1 {
			return nil
		}
		tmp := &slice[0]
		slice = slice[1:]
		return tmp
	}
	return &Iter[T]{next: next}
}

func IterOf[T any](elems ...T) *Iter[T] {
	return IterSlice(elems)
}

func (iter *Iter[I]) ForEach(f func(I)) {
	for ptr := iter.next(); ptr != nil; ptr = iter.next() {
		f(*ptr)
	}
}

func (iter *Iter[I]) Map(f func(I) I) *Iter[I] {
	next := func() *I {
		ptr := iter.next()
		if ptr == nil {
			return nil
		}
		*ptr = f(*ptr)
		return ptr
	}
	return &Iter[I]{next: next}
}

func (iter *Iter[I]) Filter(f func(I) bool) *Iter[I] {
	next := func() *I {
		for {
			ptr := iter.next()
			if ptr == nil {
				return nil
			}
			if f(*ptr) {
				return ptr
			}
		}
	}
	return &Iter[I]{next: next}
}

func (iter *Iter[I]) Any(f func(I) bool) bool {
	for ptr := iter.next(); ptr != nil; ptr = iter.next() {
		if f(*ptr) {
			return true
		}
	}
	return false
}

func (iter *Iter[I]) All(f func(I) bool) bool {
	for ptr := iter.next(); ptr != nil; ptr = iter.next() {
		if !f(*ptr) {
			return false
		}
	}
	return true
}

func (iter *Iter[I]) Collect() []I {
	slice := []I{}
	for ptr := iter.next(); ptr != nil; ptr = iter.next() {
		slice = append(slice, *ptr)
	}
	return slice
}

func Print[T any](elem T) {
	fmt.Println(elem)
}

func SqrInt[I constraints.Signed](i I) I {
	return i * i
}

func main() {
	s := IterOf("hello,", "world.", "!!").
		Map(strings.Title).
		Collect()

	fmt.Println(s)

	IterOf(2, 3, 5).
		Map(SqrInt[int]).
		ForEach(Print[int])

	isEven := func(i int) bool { return i%2 == 0 }

	fmt.Println(IterOf(2, 3, 4).Any(isEven))
	fmt.Println(IterOf(2, 3, 4).All(isEven))

	fmt.Println(IterOf(2, 3, 4).Filter(isEven).Collect())
}
