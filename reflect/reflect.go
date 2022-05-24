package main

import (
	"fmt"
	"reflect"
)

func foo() {
	var A interface{}
	A = 100

	v := reflect.ValueOf(A)
	B := v.Interface()

	if A == B {
		fmt.Printf("They are same!\n")
	}
}

func modifyValue() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println("x : ", v.Elem().Interface())
}

func main() {
	var x float64 = 3.4
	t := reflect.TypeOf(x) // t is reflect.Type
	fmt.Println("type:", t)

	v := reflect.ValueOf(x) // v is reflect.Value
	fmt.Println("value:", v)

	modifyValue()
}
