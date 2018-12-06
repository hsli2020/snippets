package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

func main() {
    secret := "mysecret"
    data := "data"
    fmt.Printf("Secret: %s\nData:   %s\n", secret, data)

    // Create a new HMAC by defining the hash type and the key (as byte array)
    h := hmac.New(sha256.New, []byte(secret))

    // Write Data to it
    h.Write([]byte(data))

    // Get result and encode as hexadecimal string
    sha := hex.EncodeToString(h.Sum(nil))

    fmt.Println("Result: " + sha)
	// or
    fmt.Printf("Result: %x\n", h.Sum(nil))

	// sha256.Sum256(...)
    fmt.Printf("Result: %x\n", sha256.Sum256([]byte("data")))
}
