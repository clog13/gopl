package mch9

import (
	"bufio"
	"fmt"
	"strings"
)

func Nscan() {
	sc := NewScanner(bufio.NewReader(NewReader("ab.cd ef.gh ij.kl mn.op")))
	sc.Split(ScanDot)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}

	tsc := bufio.NewScanner(bufio.NewReader(strings.NewReader("ab.cd ef.gh ij.kl mn.op")))
	tsc.Split(bufio.ScanWords)
	for tsc.Scan() {
		fmt.Println(tsc.Text())
	}
}
