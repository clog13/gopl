package ch5

import "fmt"

func Anon() {
	f := squares()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
	fmt.Println(f(4))

	fmt.Printf("%*s</%s>\n", 2, "^", "?")
}

func squares() func(y int) int {
	x := 0
	return func(y int) int {
		t := y
		fmt.Print(t)
		x++
		return x * x
	}
}
