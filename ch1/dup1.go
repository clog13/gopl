package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func Dup1() {
	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	counts[input.Text()]++
	// }

	// fmt.Println("--------------")
	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }
	f, err := os.Open("ch1/dup1.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
	} else {
		input := bufio.NewScanner(f)
		for input.Scan() {
			fmt.Println(input.Text())
		}
	}
	f.Close()
}
