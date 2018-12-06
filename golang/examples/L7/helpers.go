package helpers

import (
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"
	"time"
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -(MaxInt - 1)
)

func StructToBSONMap(st interface{}) (m map[string]interface{}) {

	s := reflect.ValueOf(st).Elem()
	typeOfT := s.Type()

	m = make(map[string]interface{})

	for i := 0; i < s.NumField(); i++ {

		field := s.Field(i)
		typeField := typeOfT.Field(i)

		fieldName := strings.Split(typeField.Tag.Get("bson"), ",")[0]

		if fieldName == "" {
			fieldName = typeField.Name
		}

		m[fieldName] = field.Interface()
	}

	return
}

func ArrayOfBytes(i int, b byte) (p []byte) {

	for i != 0 {

		p = append(p, b)
		i--
	}
	return
}

func FitBytesInto(d []byte, i int) []byte {

	if len(d) < i {

		dif := i - len(d)

		return append(ArrayOfBytes(dif, 0), d...)
	}

	return d[:i]
}

func StripByte(d []byte, b byte) []byte {

	for i, bb := range d {

		if bb != b {
			return d[i:]
		}
	}

	return nil
}

func IsNil(v interface{}) bool {
	return reflect.ValueOf(v).IsNil()
}

func DecodeJSON(r io.Reader, t interface{}) (err error) {

	err = json.NewDecoder(r).Decode(t)
	return
}

func SHA1(data []byte) string {

	hash := sha1.New()
	hash.Write(data)
	return SHAString(hash.Sum(nil))
}

func SHA256(data []byte) []byte {

	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

func SHAString(data []byte) string {
	return fmt.Sprintf("%x", data)
}

// From http://devpy.wordpress.com/2013/10/24/create-random-string-in-golang/
func RandomString(n int) string {

	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func RandomInt(a, b int) int {

	var bytes = make([]byte, 1)
	rand.Read(bytes)

	per := float32(bytes[0]) / 256.0
	dif := Max(a, b) - Min(a, b)

	return Min(a, b) + int(per*float32(dif))
}

func Max(a, b int) int {

	if a >= b {

		return a
	}

	return b
}

func Min(a, b int) int {

	if a <= b {

		return a
	}

	return b
}

func EncodeBase64(data []byte) []byte {

	base64data := []byte{}
	base64.StdEncoding.Encode(base64data, data)
	return base64data
}

func DecodeBase64(base64data []byte) (data []byte) {

	base64.StdEncoding.Decode(data, base64data)
	return
}

func EncodeBigsBase64(is ...*big.Int) []byte {

	arr := []byte{}
	for _, i := range is {
		arr = append(arr, i.Bytes()...)
	}
	return EncodeBase64(arr)
}

func DecodeBigsBase64(d []byte, i int) []*big.Int {

	arr := make([]*big.Int, i)
	is := DecodeBase64(d)
	l := len(is) / i

	for i, _ := range is {

		arr[i] = big.NewInt(0).SetBytes(is[l*i : l*(i+1)])
	}

	return arr
}

func Timeout(i time.Duration) chan bool {

	t := make(chan bool)
	go func() {
		time.Sleep(i)
		t <- true
	}()

	return t
}
