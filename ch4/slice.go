package ch4

import "fmt"

func Slic() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	data := []string{"one", "", "three"}
	// data = nonempty(data)
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func nonempty(strings []string) []string {
	res := strings[:0]
	for _, s := range strings {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}
