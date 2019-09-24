package main

import (
	"encoding/json"
	"fmt"
)

type P struct {
	N int
}

func main() {
	n1 := `{"N": 5}`
	p := P{}
	json.Unmarshal([]byte(n1), &p)
	n2 := ""
	json.Unmarshal([]byte(n2), &p)
	fmt.Println(p.N)
}
