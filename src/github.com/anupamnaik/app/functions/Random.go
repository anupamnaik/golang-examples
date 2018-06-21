package functions

import (
	"math/rand"
	"time"
)

// RandInt generate a random number less than given max.
func RandInt(max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max)
}
