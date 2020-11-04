package random

import (
	"math/rand"
	"time"
)

// allowed characters to generate
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

// randomized seed used for generator
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// String returns a pseudorandomly generated string
func String(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
