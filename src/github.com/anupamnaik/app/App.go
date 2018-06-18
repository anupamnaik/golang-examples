package main

import (
	"fmt"

	"github.com/anupamnaik/app/functions"
)

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("Strating application...")
	fmt.Println(functions.RandomNumberInt(100))

	a, b := functions.SwapInt(12, 34)
	fmt.Println(a, b)

	fmt.Println("End here.")
}
