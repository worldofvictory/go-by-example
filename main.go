package main

import (
	"fmt"
	"time"
)

func vals() (int, int) {
	return 3, 7
}

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
	fmt.Println("##########################")
	//slice
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("##########################")
	s = append(s, "d")
	s = append(s, "e", "f")
	s[5] = "a"
	fmt.Println("get:", s[0])
	fmt.Println("apd:", s)

	n, l := vals()
	fmt.Println(n)
	fmt.Println(l)

	_, c := vals()
	fmt.Println(c)
}
