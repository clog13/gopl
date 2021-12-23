package ch3

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func Sring() {
	str := "\xfd"
	fmt.Println(str)
	str = "\377" // � 不是一个正常的utf8
	fmt.Println(str)

	doc := `fdsaf
    fdsafdsf
	fdsafsdfsad   
	
	fdsafdsafds
    af
	dsafsdaf` + "`" + `<>`
	fmt.Println(doc)

	w1 := "世界"
	w2 := "\xe4\xb8\x96\xe7\x95\x8c"
	w3 := "\u4e16\u754c"
	w4 := "\U00004e16\U0000754c"
	fmt.Printf("%s,%s,%s,%s\n", w1, w2, w3, w4)

	s := "Hello, 世界" // "\xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	var buf bytes.Buffer
	buf.WriteByte('[')
	for i := 0; i < 10; i++ {
		buf.WriteRune('a')
		buf.WriteString("\u4e16")
	}
	fmt.Printf("%s\n", &buf)
}
