/*
Package yuyuid provides an implementation of Universally Unique Identifier (UUID) version 3, 4 and 5.

Example usage:
	package main

	import (
		"fmt"

		"github.com/komuw/yuyuid"
	)

	func main() {
		UUID4 := yuyuid.UUID4()
		fmt.Println("UUID4 is", UUID4)

		UUID5 := yuyuid.UUID5(yuyuid.NamespaceDNS, "SomeName")
		fmt.Println("UUID5", UUID5)
	}
*/
package yuyuid

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

const (
	reservedNcs       byte = 0x80 //Reserved for NCS compatibility
	rfc4122           byte = 0x40 //Specified in RFC 4122
	reservedMicrosoft byte = 0x20 //Reserved for Microsoft compatibility
	reservedFuture    byte = 0x00 // Reserved for future definition.
)

//The following standard UUIDs are for use with UUID3() or UUID5().
var (
	NamespaceDNS  = UUID{107, 167, 184, 16, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NamespaceURL  = UUID{107, 167, 184, 17, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NamespaceOID  = UUID{107, 167, 184, 18, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
	NamespaceX500 = UUID{107, 167, 184, 20, 157, 173, 17, 209, 128, 180, 0, 192, 79, 212, 48, 200}
)

// UUID represents a UUID
type UUID [16]byte

func (u UUID) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func (u *UUID) setVariant(variant byte) {
	switch variant {
	case reservedNcs:
		u[8] &= 0x7F
	case rfc4122:
		u[8] &= 0x3F
		u[8] |= 0x80
	case reservedMicrosoft:
		u[8] &= 0x1F
		u[8] |= 0xC0
	case reservedFuture:
		u[8] &= 0x1F
		u[8] |= 0xE0
	}
}

func (u *UUID) setVersion(version byte) {
	u[6] = (u[6] & 0x0F) | (version << 4)
}

// UUID3 generates a version 3 UUID
func UUID3(namespace UUID, name string) UUID {
	var uuid UUID
	var version byte = 3
	hasher := md5.New()
	hasher.Write(namespace[:])
	hasher.Write([]byte(name))
	sum := hasher.Sum(nil)
	copy(uuid[:], sum[:len(uuid)])

	uuid.setVariant(rfc4122)
	uuid.setVersion(version)
	return uuid
}

// UUID4 generates a version 4 UUID
func UUID4() UUID {

	var uuid UUID
	var version byte = 4

	// Read is a helper function that calls io.ReadFull.
	_, err := rand.Read(uuid[:])
	if err != nil {
		panic(err)
	}

	uuid.setVariant(rfc4122)
	uuid.setVersion(version)
	return uuid
}

// UUID5 generates a version 5 UUID
func UUID5(namespace UUID, name string) UUID {
	var uuid UUID
	var version byte = 5
	hasher := sha1.New()
	hasher.Write(namespace[:])
	hasher.Write([]byte(name))
	sum := hasher.Sum(nil)
	copy(uuid[:], sum[:len(uuid)])

	uuid.setVariant(rfc4122)
	uuid.setVersion(version)
	return uuid
}
