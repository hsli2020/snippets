package main

import (
	"bytes"
	"encoding/gob"
    "encoding/hex"
	"fmt"
	"log"
)

type User struct {
	ID   string
	In   string
	Out  string
}

func (u User) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(u)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

func main() {
    user := &User{ "123", "Input", "Output"}
    fmt.Printf("%s", hex.Dump(user.Serialize()))
}
