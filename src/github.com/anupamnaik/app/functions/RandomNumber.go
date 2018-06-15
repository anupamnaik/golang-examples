package functions

import "math/rand"

// RandomNumberInt generate a random number less than given limit.
func RandomNumberInt(limit int) int {
	return rand.Intn(limit)
}
