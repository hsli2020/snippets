Do is an experiment in mimicking the Result type (known from languages like Rust and Haskell). 
The idea is to minimize the amount of error checking and let the developer focus on the happy 
path, without undermining the power of Go's errors. The code below requires Go 1.18+, as it 
makes use of generic type parameters.

// ========== do.go

package do

type Void struct{}

type Result[T any] interface {
	Unwrap() (T, error)
}

type defaultResult[T any] struct {
	val T
	err error
}

// Unwrap brings a Result back to Go-land.
// At some point, the developer will have to see whether the result has a concrete value,
// or an error has been returned at some point.

func (r *defaultResult[T]) Unwrap() (T, error) {
	return r.val, r.err
}

// Try takes a value and an error straight from a regular Go function execution
// The goal is to wrap those in a Result and use that in successive operations,
// without havign to perform an immediate error check

func Try[T any](val T, err error) Result[T] {
	return &defaultResult[T]{val, err}
}

// Then is a Result success continuation.
// It allows the developer to use a previously returned Result and use its value, in case 
// there was no error. If the Result contains an error, the operation would return immediately,
// thus not proceeding to any further operation.

func Then[In, Out any](r Result[In], successMapper func(r In) (Out, error)) Result[Out] {
	var zeroVal Out

	val, err := r.Unwrap()
	if r != nil {
		return &defaultResult[Out]{zeroVal, err}
	}

	out, err := successMapper(val)
	return &defaultResult[Out]{out, err}
}


// ========== example.go

package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/preslavrachev/do"
)

func main() {

	purchase := do.Try(getPurchase("purchase-1234"))

	// The actual execution will stop here. Every next operation will be a Nop, 
	// since there was an error in the first call
	purchaseJson := do.Then(purchase, serializePurchase)

	// At the end of a long series of operations, we are back to Go-land.
	// Unwrapping requires a normal go-style error checking, thus it is 100% idiomatic Go code
	jsonBytes, err := purchaseJson.Unwrap()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(jsonBytes)
}

// Purchase is a type added for illustration
type Purchase struct {
}

// getPurchase is a function returning an error for illustration purposes
func getPurchase(purchaseID string) (Purchase, error) {
	return Purchase{}, errors.New("some random error")
}

// serializePurchase is a function added for illustration purposes
func serializePurchase(p Purchase) ([]byte, error) {
	return json.Marshal(&p)
}
