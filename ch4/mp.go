package ch4

import "fmt"

func Mp() {
	ages := make(map[string]int)
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages["bob"])
	ages["bob"]++
	fmt.Println(ages["bob"])
}
