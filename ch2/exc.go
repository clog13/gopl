package ch2

import "fmt"

type f6 float64

func Exc() {
	// var a, b = 1, a + 2
	// c, d := 1, c+2
	// fmt.Printf("%d,%d", a, b)
	// fmt.Printf("%d,%d", c, d)

	a, b := 1, 2
	a, b = b, a
	c, d := a, b
	c, d = d, c

	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	*p++
	fmt.Println(x)

	fmt.Println(f() == f())

	pp := new([0]int)
	qq := new([0]int)
	fmt.Println(&pp)
	fmt.Println(&qq)

	v := 1
	z := 0
	v++
	// z = v++
	v = z

	ff := f6(3)
	fmt.Println(ff.String()) // 3^C
	fmt.Printf("%v\n", ff)   // 3^C 变量的自然形式(natural format)
	fmt.Printf("%s\n", ff)   // 3^C 字符串
	fmt.Println(ff)          // 3^C
	fmt.Printf("%g\n", ff)   // 3   浮点数
	fmt.Println(float64(ff)) // 3
}

func (c f6) String() string {
	return fmt.Sprintf("%g^C", c)
}

func f() *int {
	v := 1
	return &v
}
