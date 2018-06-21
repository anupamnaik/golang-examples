package functions

import "fmt"

// Sqrt finds the square root of x.
func Sqrt(x float64) float64 {
	z, y := 1.0, 0.0
	for {
		z -= (z*z - x) / (2 * z)
		if z == y {
			break
		}
		fmt.Println(z)
		y = z
	}
	return z
}
