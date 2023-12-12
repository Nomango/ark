package ark

import (
	crand "crypto/rand"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// NewRandomString generate random string by specify chars.
func NewRandomString(n int, alphabets ...byte) string {
	return string(NewRandomBytes(n, alphabets...))
}

// NewRandomBytes generate random []byte by specify chars.
func NewRandomBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var bytes = make([]byte, n)
	var randByMath bool
	if num, err := crand.Read(bytes); num != n || err != nil {
		randByMath = true
	}
	for i, b := range bytes {
		if randByMath {
			bytes[i] = alphabets[rand.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}
