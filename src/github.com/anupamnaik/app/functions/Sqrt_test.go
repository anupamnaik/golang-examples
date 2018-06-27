package functions

import (
	"math"
	"testing"
)

// TestSqrt test the sqrt function
func TestSqrt(t *testing.T) {
	expected := math.Sqrt(5)
	actual, _ := Sqrt(float64(5.0))

	if expected != actual {
		t.Error("test failed, value expected")
	}

	expectedError := ErrNegativeSqrt(float64(-5.0))
	_, err := Sqrt(float64(-5.0))

	if expectedError != err {
		t.Error("test failed, error expected")
	}
}
