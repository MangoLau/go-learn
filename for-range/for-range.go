package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用传统 for 循环遍历切片
func ForExpression() {
	s := []int{1, 2, 3}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

// 使用 for-range 遍历切片
func ForRangeExpression() {
	s := []int{1, 2, 3}

	for i := range s {
		fmt.Println(s[i])
	}
}

func PrintSlice() {
	s := []int{1, 2, 3}
	var wg sync.WaitGroup

	wg.Add(len(s))
	for _, v := range s {
		go func() {
			fmt.Println(v) // 输出：3 行 3
			wg.Done()
		}()
	}
	wg.Wait()
}

func RangeTimmer() {
	t := time.NewTimer(time.Second)

	for _ = range t.C {
		fmt.Println("hi")
	}
}

func RangeDemo() {
	s := []int{1, 2, 3}
	for i := range s {
		s = append(s, i)
	}
}

// range 作用于数组时，从下标 0 开始依次遍历数组元素，返回元素的下标和元素值。
func RangeArray() {
	a := [3]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("index: %d value: %d\n", i, v)
	}
}

func RangeSlice() {
	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Printf("index: %d value: %d\n", i, v)
	}
}

func RangeMap() {
	m := map[string]string{"animal": "monkey", "fruit": "apple"}
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func RangeChannel() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"

	time.AfterFunc(time.Microsecond, func() {
		close(c)
	})

	for e := range c {
		fmt.Printf("element: %s\n", e)
	}
}

func main() {
	// PrintSlice()
	// RangeTimmer() // 打印后阻塞
	// RangeDemo() // 在遍历开始前已经决定了循环次数，所以迭代过程中向切片追加元素不会导致无休止执行，函数可以正常退出。
	RangeChannel()
}
