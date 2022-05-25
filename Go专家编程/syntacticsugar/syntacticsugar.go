package main

import "fmt"

// rule := "Short variable declarations" // syntax error: non-declaration statement outside function body

func fun1() {
	i := 0
	i, j := 1, 2
	fmt.Printf("i = %d, j = %d\n", i, j)
}

func fun2() {
	i := 0
	fmt.Println(i)
}

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}

func main() {
	// fun1()
	fun2()
}