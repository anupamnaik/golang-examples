package functions

import "fmt"

// Repeat repeat the given function n times.
func Repeat(n int, repeatFunction func(c int)) {

	fmt.Println("Repeater...")

	for i := 0; i < n; i++ {
		go repeatFunction(i)
	}
}
