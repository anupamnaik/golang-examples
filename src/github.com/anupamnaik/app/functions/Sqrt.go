package functions

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt error when sqrt of negative number is attempted
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt finds the square root of x.
func Sqrt(x float64) (float64, error) {
	if err := ErrNegativeSqrt(x); x < 0 {
		return 0, err
	}
	z, y := 1.0, 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-y) < 1e-8 {
			break
		}
		y = z
	}
	return z, nil
}
