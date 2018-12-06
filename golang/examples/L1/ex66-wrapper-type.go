// There is a problem that comes up in C++ a lot (in my experience) which has no
// good C++ answer, but which Go has a great answer for (OMG, Tim said something
// nice about Go!).  The problem is: wrappers.
//
// I often find myself wanting to implement a particular interface (abstract base
// class in C++) by delegating all method calls to another implementation of that
// interface, but intercepting just one or two of them.  I can't use inheritance
// because I don't know the concrete type of the wrappee.  So I implement my own
// wrapper type, and I implement every one of the interface's methods as simple
// calls to the wrappee.  This works, but it can be INCREDIBLY tedious to do,
// especially if the interface has a lot of methods.  It works right up until
// someone adds a new method to the base interface -- then my build breaks and I
// have to implement it too.
//
// Go's answer is very elegant - embed the interface type to "inherit" methods,
// "override" mthods you care about, and init from an instance of that interface.
//
// thockin@google.com
// http://play.golang.org/p/ssz2AKIj_y

package main

import "fmt"
import "math/rand"

// Base interface I want to wrap.
type Frobber interface {
	FrobGently()
	FrobAggressively()
	FrobWithPrejudice()
}

// One implementation.
type wetFrobber struct{}

func (wet wetFrobber) FrobGently() {
	fmt.Println("wet.FrobGently()")
}

func (wet wetFrobber) FrobAggressively() {
	fmt.Println("wet.FrobAggressively()")
}

func (wet wetFrobber) FrobWithPrejudice() {
	fmt.Println("wet.FrobWithPrejudice()")
}

// Another implementation.
type dryFrobber struct{}

func (dry dryFrobber) FrobGently() {
	fmt.Println("dry.FrobGently()")
}

func (dry dryFrobber) FrobAggressively() {
	fmt.Println("dry.FrobAggressively()")
}

func (dry dryFrobber) FrobWithPrejudice() {
	fmt.Println("dry.FrobWithPrejudice()")
}

// Choose a random implementation.
func newFrobber() Frobber {
	// Yes, I know playground is not actually random.
	if rand.Intn(2) == 0 {
		return dryFrobber{}
	}
	return wetFrobber{}
}

// My wrapper.
type FrobWrapper struct {
	Frobber
}

// Override just one method.
func (wrap FrobWrapper) FrobWithPrejudice() {
	fmt.Printf("OMG: ")
	wrap.Frobber.FrobWithPrejudice()
}

func main() {
	var f Frobber = FrobWrapper{newFrobber()}
	f.FrobGently()
	f.FrobAggressively()
	f.FrobWithPrejudice()
}
