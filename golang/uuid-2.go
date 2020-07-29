package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Uuid [16]byte

func init() {
	gob.Register(Uuid{})
}

func New() (Uuid, error) {
	var u Uuid
	_, err := io.ReadFull(rand.Reader, u[:])
	if err != nil {
		return u, err
	}
	u[6] &= 0x0F // clear version
	u[6] |= 0x40 // set version to 4 (random uuid)
	u[8] &= 0x3F // clear variant
	u[8] |= 0x80 // set to IETF variant
	return u, nil
}

func (u Uuid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		u[0:4], u[4:6], u[6:8], u[8:10], u[10:16])
}

func (u Uuid) Bytes() []byte {
	return u[:]
}

func (u Uuid) MarshalJSON() ([]byte, error) {
	b, _ := json.Marshal(u.String())
	return b, nil
}

func (u *Uuid) UnmarshalJSON(b []byte) error {
	if u == nil {
		return fmt.Errorf("Uuid receiver nil")
	}
	var uuid string
	err := json.Unmarshal(b, &uuid)
	if err != nil {
		return err
	}
	uid, err := ParseUuid(uuid)
	if err != nil {
		return err
	}
	*u = uid
	return nil
}

func (u Uuid) Equals(uu Uuid) bool {
	return bytes.Equal(u[:], uu[:])
}

var empty = Uuid{}

func (u Uuid) IsEmpty() bool {
	return u.Equals(empty)
}

func UuidFromBytes(input []byte) (Uuid, error) {
	var u Uuid
	if len(input) != 16 {
		return u, errors.New("UUIDs must be exactly 16 bytes long")
	}

	copy(u[:], input)
	return u, nil
}

func ParseUuid(input string) (Uuid, error) {
	var u Uuid
	j := 0
	for _, r := range input {
		switch {
		case r == '-' && j&1 == 0:
			continue
		case r >= '0' && r <= '9' && j < 32:
			u[j/2] |= byte(r-'0') << uint(4-j&1*4)
		case r >= 'a' && r <= 'f' && j < 32:
			u[j/2] |= byte(r-'a'+10) << uint(4-j&1*4)
		case r >= 'A' && r <= 'F' && j < 32:
			u[j/2] |= byte(r-'A'+10) << uint(4-j&1*4)
		default:
			return Uuid{}, fmt.Errorf("invalid UUID %q", input)
		}
		j += 1
	}
	if j != 32 {
		return Uuid{}, fmt.Errorf("invalid UUID %q", input)
	}
	return u, nil
}
