package main

import "fmt"

func RecoverDemo1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic("demo")
	fmt.Println("B")
}

func RecoverDemo2() {
	defer func() {
		fmt.Println("C")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic("demo")
	fmt.Println("B")
}

func RecoverDemo3() {
	defer func() {
		func() { // recover 在 defer 嵌套函数中无效
			if err := recover(); err != nil {
				fmt.Println("A")
			}
		}()
	}()

	panic("demo")
	fmt.Println("B")
}

func RecoverDemo4() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("B")
		}
	}()

	panic("demo")
	fmt.Println("C")
}

func RecoverDemo5() {
	foo := func() int {
		defer func() {
			recover()
		}()

		panic("demo")

		return 10
	}

	ret := foo()
	fmt.Println(ret)
}

func RecoverDemo6() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A")
		}
	}()

	panic(nil)
	fmt.Println("B")
}

func main() {
	// RecoverDemo1()
	// RecoverDemo2()
	// RecoverDemo3()
	// RecoverDemo4()
	// RecoverDemo5()
	RecoverDemo6()
}
