// go test -v -cpu=4 -run=none -bench=. -benchtime=10s -benchmem praseip_test.go
package praseip_test

import (
	"net"
	"testing"
)

func fastParseIP(s string) (ip [4]byte, ok bool) {
	var i, n int
	for _, b := range []byte(s) {
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n = n*10 + int(b-'0')
		case '.':
			if n >= 256 {
				return
			}
			ip[i] = byte(n)
			n = 0
			i++
		default:
			return
		}
	}
	if i != 3 || n >= 256 {
		return
	}
	ip[i] = byte(n)
	ok = true

	return
}

func FastParseIP(s string) net.IP {
	ip, ok := fastParseIP(s)
	if !ok {
		return nil
	}
	return ip[:]
}

func TestFastParseIP(t *testing.T) {
	println(FastParseIP("1.22.44.255").String())
}

func BenchmarkNetParseIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		net.ParseIP("8.8.8.8")
	}
}

func BenchmarkFastParseIP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastParseIP("8.8.8.8")
	}
}

/*
BenchmarkNetParseIP-4           251827638               48.20 ns/op           16 B/op          1 allocs/op
BenchmarkFastParseIP-4          1000000000               8.072 ns/op           0 B/op          0 allocs/op
*/
