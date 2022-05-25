for-range 表达式作用于所有的集合类型，包括数组、切片、string、map，甚至还可以遍历 channel。

## 特性浏览
1. 作用于数组

range 作用于数组时，从下标 0 开始依次遍历数组元素，返回元素的下标和元素值。
```go
func RangeArray() {
	a := [3]int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("index: %d value: %d\n", i, v)
	}
}
```

2. 作用于切片

range 作用于切片时，与数组类似，返回元素的下标和元素值。
```go
func RangeSlice() {
	s := []int{1, 2, 3}
	for i, v := range s {
		fmt.Printf("index: %d value: %d\n", i, v)
	}
}
```

3. 作用于 string

range 作用于 string 时，仍然返回元素的下标和元素值，但由于 string 底层使用 Unicode 编码存储字符，字符可能占用 1～4 个字节（比如汉字），所以下标可能是不连续的，并且元素值是该字符对应的 Unicode 编码的首个字节的值。

另外需要注意的是，range 的第二个返回值类型为 rune 类型，它仅代表 Unicode 编码的 1 个字节

4. 作用于 map
range 作用于 map 时，返回每个元素的 key 和 value：
```go
func RangeMap() {
	m := map[string]string{"animal": "monkey", "fruit": "apple"}
	for k, v := range m {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}
```
map 的数据结构本事没有顺序的概念，它仅存储 key-value 对，所以 range 分别返回 key 和 value。

另外，如果遍历过程中修改 map（增加或删除元素），则 range 行为是不确定的，删除元素不可能被遍历到，新加的元素可能遍历不到，总之尽量不要在循环中修改 map。

5. 作用于 channel
range 作用于 channel 时，会返回 channel 中所有的元素。
```go
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
```

channel 包含两个元素，range 遍历完两个元素后会阻塞等待，直到定时器到期后关闭 channel 才结束遍历。

channel 中的元素没有下标的概念，所以其最多只能返回一个元素。
range 会阻塞等待 channel 中的数据，直到 channel 被关闭，同时，如果 range 作用于值为 nil 的 channel 时，则会永久阻塞。

## 小结
对于数组、切片、string 和 map 类型，如果只有一个循环变量接收 range 返回值，则相当于省略掉了第二个返回值。
```
for lst := range xxx
// 等价于
for lst, _ := range xxx
```