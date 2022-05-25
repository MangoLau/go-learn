package main

import "fmt"

func AppendDemo() {
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

func foo() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
}

func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task: %s\n", task)
		}()
	}
}

func Process2(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func(t string) {
			fmt.Printf("Worker start process task: %s\n", t)
		}(task)
	}
}

func Double(a int) int {
	return a * 2
}

func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
		return true
	}

	return false
}

func UpdateTable() {
	// 在 defer 中决定提交还是会滚
	defer func ()  {
		if IsPanic() {
			// Rollback transaction
		} else {
			// Comit transaction
		}
	}()

	// Database update operation...
}

func main() {
	// AppendDemo()
	foo()
}
