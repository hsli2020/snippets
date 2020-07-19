package main

// This code comes from
// https://github.com/gorilla/securecookie/blob/master/securecookie.go

import (
	"crypto/rand"
	"fmt"
	"io"

	//"github.com/gorilla/securecookie"
)

// GenerateRandomKey creates a random key with the given length in bytes.
// On failure, returns nil.
//
// Note that keys created using `GenerateRandomKey()` are not automatically
// persisted. New keys will be created when the application is restarted, and
// previously issued cookies will not be able to be decoded.
//
// Callers should explicitly check for the possibility of a nil return, treat
// it as a failure of the system random number generator, and not continue.
func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}

func main() {
	//fmt.Printf("%x\n", securecookie.GenerateRandomKey(16))
	//fmt.Printf("%x\n", securecookie.GenerateRandomKey(24))
	//fmt.Printf("%x\n", securecookie.GenerateRandomKey(32))

	fmt.Printf("%x\n", GenerateRandomKey(16))
	fmt.Printf("%x\n", GenerateRandomKey(24))
	fmt.Printf("%x\n", GenerateRandomKey(32))
}
