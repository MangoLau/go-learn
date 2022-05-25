## 特性
### 1. 用法
#### 1）声明
声明一个空字符变量再赋值：
```go
var s1 string
s1 = "Hello World"
```
需要注意的是空字符只是长度为 0，但不是 nil。

使用短变量声明：
```go
s2 := "Hello World" // 初始化的字符串
```

#### 2）双引号与反引号的区别
区别在于对特殊字符的处理。

反引号包含的字符串所见即所得。
```go
s := `Hi,
this is "RainbowMango".`
```

使用双引号表示时，需要对特殊字符转义：
```go
s := "Hi, \nthis is \"RanbowMango\"."
```

#### 3）字符串拼接
字符串可以使用加号进行拼接：
```go
s = s + "a" + "b"
```

需要注意的是，字符串拼接时会触发内存分配及内存拷贝，单行语句拼接多个字符串只分配一次内存。比如：
```go
s = s + "a" + "b"
```
在拼接时，会先计算最终字符串的长度后再分配内存。

#### 4）类型转换
`[]byte` 转 `string`:
```go
func ByteToString() {
	b := []byte{'H', 'e', 'l', 'l', 'o'}
	s := string(b)
	fmt.Println(s)
}
```

`string` 转 `[]byte`:
```go
func StringToByte() {
	s := "Hello"
	b := []byte(s)
	fmt.Println(b)
}
```

### 2. 特点
#### 1）UTF 编码
在使用 for-range 遍历字符串时，每次迭代将返回字符 UTF-8 编码的首个字节的下标及字节值，这意味着下标可能不连续。
```go
func StringIteration() {
	s := "中国"
	for index, value := range s {
		fmt.Printf("index: %d, value: %c\n", index, value)
	}
}
```

#### 2）值不可修改
字符串可以为空，但值不会是 nil。
字符串不可以修改。
字符串变量可以接受新的字符串赋值，但不能通过下标方式修改字符串中的值。
```go
s := "Hello"
&s[0] = byte[104] // 非法
s = "hello" // 合法
```