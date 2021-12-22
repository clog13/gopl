package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", ">-S-<", "separator")

func Echo() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if *n {
		fmt.Println("<n>")
	}
}

/**
go run main.go -h                                                                                          [19:24:44]

Usage of /tmp/go-build089265255/b001/exe/main:
  -n    omit trailing newline
  -s string
        separator (default " ")
exit status 2
*/
