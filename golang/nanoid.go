/*
Copyright 2022 jaevor.
License can be found in the LICENSE file.

Original reference: https://github.com/ai/nanoid
*/

package nanoid

import (
	cryptoRand "crypto/rand"
	"errors"
	"math"
	"math/bits"
	mathRand "math/rand"
	"sync"
)

// Default characters (A-Za-z0-9_-).
var defaultAlphabet = []byte("useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict")

type generator = func() string

/*
Creates a new generator for Nano IDs.

📝 Recommended (standard) length is 21

Returns error if length is not between 2 and 255 (inclusive).

Concurrency safe.
*/
func Standard(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	// Multiplying by 128 and using an offset so
	// that the bytes only have to be refilled
	// every 129th nanoid. This is more efficient.
	// b holds the random crypto bytes.
	b := make([]byte, length*128)
	size := len(b)
	offset := 0

	cryptoRand.Read(b)

	// Since the default alphabet is ASCII,
	// we don't have to use runes here. ASCII max is
	// 128, so byte will be perfect.
	// id := make([]rune, length)
	id := make([]byte, length)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		// If all the bytes in the slice
		// have been used, refill.
		if offset == size {
			cryptoRand.Read(b)
			offset = 0
		}

		for i := 0; i < length; i++ {
			// Index using the offset.
			id[i] = defaultAlphabet[b[i+offset]&63]
		}

		// Extend the offset.
		offset += length

		return string(id)
	}, nil
}

/*
Create a non-secure Nano ID generator.
Non-secure is faster than secure because it uses pseudorandom numbers.

Returns error if length is not between 2 and 255 (inclusive).

⚠ Remember to seed using rand.Seed().

Concurrency safe.
*/
func StandardNonSecure(length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	// b holds pseudorandom bytes.
	b := make([]byte, length*128)
	size := len(b)
	offset := 0

	mathRand.Read(b)

	// Reuse.
	id := make([]byte, length)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		if offset == size {
			// Refill b.
			mathRand.Read(b)
			offset = 0
		}

		for i := 0; i < length; i++ {
			/*
				"It is incorrect to use bytes exceeding the alphabet size.
				The following mask reduces the random byte in the 0-255 value
				range to the 0-63 value range. Therefore, adding hacks such
				as empty string fallback or magic numbers is unneccessary because
				the bitmask trims bytes down to the alphabet size (64).""
			*/
			id[i] = defaultAlphabet[b[i+offset]&63]
		}

		offset += length

		return string(id)
	}, nil
}

/*
Create a Nano ID generator that uses a custom alphabet.

Concurrency safe.
*/
func Custom(alphabet string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	setLen := len(alphabet)
	runicSet := []rune(alphabet)

	// Because the custom character-set is not guaranteed to have
	// 64 chars to utilise, we have to calculate a suitable mask.
	// The following calculations are 1:1 to the original implementation.
	clz := bits.LeadingZeros32((uint32(setLen) - 1) | 1)
	mask := (2 << (31 - clz)) - 1
	w := (1.6 * float64(mask*length)) / float64(setLen)
	step := int(math.Ceil(w))

	// Will be reusing the same rune and byte slices.
	id := make([]rune, length)
	b := make([]byte, step)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		for u := 0; ; {
			cryptoRand.Read(b)

			for i := 0; i < step; i++ {
				idx := b[i] & byte(mask)

				if idx < byte(setLen) {
					// id.WriteRune(runicSet[idx])
					id[u] = runicSet[idx]
					u++
					if u == length {
						return string(id)
					}
				}
			}
		}
	}, nil
}

/*
Create a non-secure Nano ID generator that uses a custom alphabet.
Non-secure is faster than secure because it uses pseudorandom numbers.

Returns error if length is not between 2 and 255 (inclusive).

⚠️ Remember to seed using rand.Seed().

Concurrency safe.
*/
func CustomNonSecure(alphabet string, length int) (generator, error) {
	if invalidLength(length) {
		return nil, errInvalidLength
	}

	runicSet := []rune(alphabet)
	setLen := len(runicSet)

	id := make([]rune, length)

	var mu sync.Mutex

	return func() string {
		mu.Lock()
		defer mu.Unlock()

		for i := 0; i < length; i++ {
			id[i] = runicSet[mathRand.Intn(setLen)]
		}

		return string(id)
	}, nil
}

var errInvalidLength = errors.New("length must be between 2 and 255 (inclusive)")

func invalidLength(length int) bool {
	return length < 2 || length > 255
}
