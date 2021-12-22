package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo1() {
	fmt.Println(strings.Join(os.Args, " "))
}

// go run echo1.go 1 a xyz
// /tmp/go-build234788996/b001/exe/echo1 1 a xyz
