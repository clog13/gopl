package ch4

import "fmt"

func Arr() {
	q := [3]int{1, 2, 3}
	q = [3]int{3, 2, 1}

	fmt.Printf("%T\n", q)
	// {n, ..., PRE_CONT: n, ..., n}
	r := [...]int{99, 12, 12, 3, 45: -23, -1, -43, -324}
	fmt.Printf("%T\t\n", r)
	for _, i := range r {
		fmt.Println(i)
	}

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "💴", RMB: "￥"}
	fmt.Println(GBP, symbol[GBP])
}
