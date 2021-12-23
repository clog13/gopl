package ch3

import "fmt"

func Integrate() {
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var i int
	var ui uint
	fmt.Printf("int 8: %d\n", &i8)  // 8bit, 1byte(字节)
	fmt.Printf("int16: %d\n", &i16) // 16bit, 2byte(字节)
	fmt.Printf("int32: %d\n", &i32) // 32bit, 4byte(字节)
	fmt.Printf("int64: %d\n", &i64) // 64bit, 8byte(字节)
	fmt.Printf("int  : %d\n", &i)   // 32 or 64bit, 4 or 8byte(字节)
	fmt.Printf("int u: %d\n", &ui)

	a := 1 &^ 3
	fmt.Println(a)
}
