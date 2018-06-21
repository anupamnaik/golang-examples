package main

import (
	"fmt"

	"github.com/anupamnaik/app/functions"
)

func main() {
	fmt.Println("Starting application...")
	fmt.Println(functions.RandInt(100))
	functions.Sqrt(25)

	fmt.Println("End here.")
}
