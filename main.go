package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	t := time.Now().Weekday()
	switch t {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	var a [5]int
	fmt.Println("emp:", a)

	a[0] = 1
	fmt.Println("emp:", a)
	fmt.Println("set:", a)
	fmt.Println("get:", a[0])

	fmt.Println("len:", len(a))
}
