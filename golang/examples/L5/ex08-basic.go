package main

import (
	"fmt"
)

type Magician interface {
	Magic()
}

type Delegate struct {
	magician Magician
}

func (delegate Delegate) Magic() {
	if _, ok := delegate.magician.(Magician); ok {
		delegate.magician.Magic()
	} else {
		fmt.Println("base magic")
	}
}

func (delegate Delegate) MoreMagic() {
	delegate.magician.Magic()
	delegate.magician.Magic()
}

type Foo struct {
	Delegate
}

func (foo Foo) Magic() {
	fmt.Println("foo magic")
}

func main() {
	f := Delegate{new(Foo)}
	f.Magic()
	f.MoreMagic()
}

// rescue.go
package main

import (
	"fmt"
)

func foo(i *int) {
	fmt.Println(*i)
}

func test() (err error) {
	defer func() {
		if recover() != nil {
			fmt.Println("exception occured")
		}
	}()

	for i := 0; i <= 10; i++ {
		foo(&i)
	}
	foo(nil) // Should be panic
	return

}

func main() {
	test()
}
