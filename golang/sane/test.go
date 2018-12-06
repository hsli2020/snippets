package main

import (
	"fmt"

	"github.com/bloom42/sane-go"
)

type D struct {
	A string
}

type C struct {
	A int64 `sane:"a"`
	D []D   `sane:"d"`
}

type S struct {
	A string  `sane:"a"`
	B []int64 `sane:"b"`
	C C       `sane:"c"`
}

func main() {
	str1 := `
a = "a"
b = [1, 2]
c = { a = 1, d = [ { a = "3.3" }, { a = "xxx" } ] }
`
	var s S

	err := sane.Unmarshal([]byte(str1), &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n\n", s)

	b, err := sane.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}
