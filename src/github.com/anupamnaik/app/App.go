package main

import (
	"fmt"

	"github.com/anupamnaik/app/functions"
)

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("Starting application...")
	fmt.Println(functions.RandomNumberInt(100))

	a, b := functions.SwapInt(12, 34)
	fmt.Println(a, b)

	s1, s2 := functions.SwapString("two", "strings")
	fmt.Println(s1, s2)

	functions.Repeat(10000, func(i int) {
		fmt.Println("Counter ", i)
	})

	fmt.Println("End here.")
}
