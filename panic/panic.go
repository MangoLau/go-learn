package main

import "fmt"

func foo() {
	defer fmt.Print("A")
	defer fmt.Print("B")

	fmt.Print("C")
	panic("demo")
	defer fmt.Print("C")
}

func PanicDemo1() {
	defer func() {
		recover()
	}()

	foo()
}

func PanicDemo2() {
	defer func() {
		recover()
	}()

	defer func() {
		fmt.Print("1")
	}()

	foo()
}

func PanicDemo3() {
	defer func() {
		fmt.Print("demo")
	}()

	go foo()
}

func PanicDemo4() {
	defer func() {
		recover()
	}()

	defer fmt.Print("A")

	defer func() {
		fmt.Print("B")
		panic("panic in defer")
		fmt.Print("C")
	}()

	panic("panic")
	fmt.Print("D")
}

func main() {
	PanicDemo4()
}
