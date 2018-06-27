package main

import (
	"fmt"

	"github.com/anupamnaik/app/functions"
)

// main the main stand-alone app
func main() {
	fmt.Println("Starting application...")

	//WebApp()
	//sampler()
	functions.Parse()

	fmt.Println("Stop.")
}

func sampler() {
	//fmt.Printf("%d \n", functions.RandInt(100))

	// New person
	p := functions.Person{FirstName: "Anupam", LastName: "Naik"}
	describe(p)
	fmt.Printf("%s \n", p)
	fmt.Printf("%s \n", &p)

}

func describe(i interface{}) {
	switch v := i.(type) {
	case fmt.Stringer:
		fmt.Println("Has String() function")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
	fmt.Printf("(%v %T) \n", i, i)
}
