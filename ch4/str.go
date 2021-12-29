package ch4

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID int
}

func Emp() {
	fmt.Println(byID(12).ID)
	byID(12).ID = 21
}

func byID(id int) *Employee {
	var EOF = errors.New("EOF")
	// fmt.Errorf(EOF.Error())
	fmt.Printf("%T\n", EOF)

	emo := Employee{12}
	var res *Employee = &emo
	return res

}
