package main

import "fmt"

func deferFuncReturn() (result int) {
    i := 1

    defer func() {
        result++
    }()
    return i
}

func foo() int {
	var i int

	defer func() {
		i++
	}()

	return 1
}

func bar() int {
	var i int

	defer func() {
		i++
	}()

	return i
}

func change() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}

func main() {
	fmt.Println(deferFuncReturn())
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println(change())
}