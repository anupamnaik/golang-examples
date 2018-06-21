package functions

import (
	"fmt"
	"time"
)

// WhenSaturday method showing the use of switch-case.
func WhenSaturday() {
	fmt.Println("When's Saturday?")
	switch today := time.Now().Weekday(); time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

// Pointer method showing the use of pointers.
func Pointer() {
	i := 6
	p := &i
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(*p)

	*p = *p * *p
	fmt.Println(*p)
	fmt.Println(i)
}
