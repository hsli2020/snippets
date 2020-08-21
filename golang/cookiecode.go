package cookiecode  // https://github.com/4thabang/go-cookiecode

import (
	"fmt"
	"log"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

/*
	TODO: how will the users add their encryption keys
	TODO: find api for hash and block key generation (optional)
	TODO: better godocs comments
*/

// Encrypt holds the values for our 'HashKey' and 'BlockKey' as a struct.
// These values must be within the Advanced Encryption Standard.
type Encrypt struct {
	HashKey  []byte // AES 256-bit
	BlockKey []byte // AES 128-bit
}

// Keys allows us to store our encryption keys securely for re-use.
func (e *Encrypt) Keys(v map[string]string) (*securecookie.SecureCookie, error) {
	// TODO: We need to allow the user to grab their environment variables when needed.
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}

	_ = &Encrypt{
		HashKey:  []byte(v["HASH_KEY"]),  // AES 256-bit
		BlockKey: []byte(v["BLOCK_KEY"]), // AES 128-bit
	}

	secure := securecookie.New(e.HashKey, e.BlockKey)

	return secure, nil
}

// EncodeType is a struct that houses the value of our key which will be
// determined by the user at runtime.
type EncodeType struct {
	Value string
	Key   string
}

// Encoder is an interface that allows us to implement our Encode function
// in order to fully encode our cookie values.
type Encoder interface {
	Encode(v map[string]string) (string, error)
}

// Encode allows us to encode our cookie in order to keep it secure, safe and unexposed.
func (e *Encoder) Encode(v map[string]string) (string, error) {
	var k Encrypt

	secure, err := k.Keys(v)
	if err != nil {
		return "", err
	}

	et := &EncodeType{
		Value: v["value"],
		Key:   v["key"],
	}

	// cookiecode.Encode(value)
	encode, err := secure.Encode(et.Key, et.Value)
	if err != nil {
		return "", err
	}

	return encode, nil
}

func (e *EncodeType) encode(v map[string]string) (*EncodeType, error) {
	et := &EncodeType{
		Value: v["value"],
		Key:   v["key"],
	}
	return et, nil
}

// Decode allows us to decode our cookie in order to consume it safely, awaay from prying eyes.
func Decode(key, cookie string) (map[string]string, error) {
	v := make(map[string]string)

	var k Encrypt

	secure, err := k.Keys(v)
	if err != nil {
		fmt.Println(err)
	}

	var value map[string]string

	// cookiecode.Decode(cookie.Name, cookie.Value)
	err = secure.Decode(key, cookie, &value)
	if err != nil {
		fmt.Println(err)
	}

	return value, nil
}
