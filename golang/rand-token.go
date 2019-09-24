package main

import (
	"crypto/rand"
	"fmt"
)

func generateToken() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func main() {
    // generate an access token
    fmt.Println("Your access token is:", generateToken())
}