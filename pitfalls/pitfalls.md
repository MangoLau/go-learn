## 1 切片扩容
### 1. 小测验
```go
func Validation() []error {
    var errs []error

    _ = append(errs, errors.New("error 1"))
    _ = append(errs, errors.New("error 2"))
    _ = append(errs, errors.New("error 3"))

    return errs
}
```

### 2. 解析
内置 append() 函数在向切片中追加元素时，如果切片存储容量不足以存储新元素，则会把当前切片扩容并产生一个新的切片。

append() 函数每次追加元素都有可能触发切片扩容，极有可能返回一个新的切片，这正是 append() 函数声明中返回值为切片的原因。使用时应该总是接收该返回值。

此外，如果不处理 append() 函数的返回值，那么编译器将给出编译错误，但上面的代码中匿名变量恰巧绕过了编译器检查。

### 3. 建议
使用 append() 函数时，谨记 append() 可能会生成新的切片，并谨慎地处理返回值。

## 2 空切片

### 1. 小测验
ValidateName() 函数用于检查某个名字是否合法，如果不为空则认为合法，否则返回一个 error。类似的，还可以有很多检查项，比如检查性别、年龄等，我们统称为子检查项。Validations() 函数用于收集所有子检查项的错误信息，将错误信息汇总到一个切片中返回。

请问 Validations() 函数有什么问题？
```go
func ValidateName(name string) error {
    if name != "" {
        return nil
    }
    return errors.New("empty name")
}

func Validations(name string) []error {
    var errs []error

    errs = append(errs, ValidateName(name))

    return errs
}
```

### 2. 解析
向切片中追加一个 nil 值是完全不会报错的，如以下代码所示。
```go
slice := append(slice, nil)
```

经过追加后，长度也会加 1.

实际上 nil 是一个预定义的值，即空值，所以完全有理由向切片中追加。

上述题目就是典型的向切片中追加 nil（当名字为空时）的问题。单纯从技术上讲是没有问题的，但在实际场景中可能存在极大风险。

题目中的函数用于收集所有的错误信息，没有错误就不应该追加到切片中。因为后续极有可能会根据切片的长度来判断是否有错误发生，例如：
```go
func foo() {
    errs := validations("")

    if len(errs) > 0 {
        println(errs)
        os.Exit(1)
    }
}
```
在上面的代码中，如果向切片中追加一个 nil 元素，那么切片的长度则不再为 0，程序很可能因此而退出，更糟糕的是，这样的切片不会打印任何内容，这无疑又增加了定位的难度。

### 3. 建议
使用 append() 函数时，谨记 append 可能会追加 nil 值，应该尽量避免追加无意义的元素。

## 3 append 的本质
### 1. 小测验
```go
func AppendDemo() {
    x := make([]int, 0, 10)
    x = append(x, 1, 2, 3)
    y := append(x, 4)
    z := append(x, 5)
    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}
```

### 2. 解析
题目先申请长度为 0 但空间为 10 的切片 x，然后分三次向 x 切片中追加元素并分别使用 x、y、z 来接收 append 的返回值。最后打印 x、y、z 的元素。

当 append 向切片 x 追加元素时，在空间足够存放新元素的情况下，新元素将从 x[len(x)] 位置开始存放，append 会生成一个新的切片，但不会修改原切片 x。

。。。

人参为切片，而切片只是一个 struct 数据结构，参数传递时发生了值拷贝，所以 append 无法操作原切片。

### 3. 参考答案
x、y、z 的输出分别为：
```
[1 2 3]
[1 2 3 5]
[1 2 3 5]
```

### 4. 建议
一般情况下，使用 append 向切片追加新的元素时，都会用原切片变量接收返回值来获得更新：
```go
A = append(A, elems...)
```
如果使用新的变量接收返回值，则需要考虑 append 返回的切片是否跟原切片共享底层的数组。

## 4 循环变量引用
### 1. 小测验
```go
func foo() {
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
}
```

### 2. 解析
该题目考察循环变量的绑定问题。题目中的 out 是一个存储整型指针的切片，在循环中，每次向切片中追加一个 i 变量的地址。在 Go 中，循环变量 i 只分配一次地址，在 3 次循环中，i 中存储的值分别为 0、1、2、3，但是 i 的地址没有变化。

所以函数打印的实际上是 i 的最终值 3.

### 3. 参考答案
函数输出
```
Values: 3 3 3
```

### 4. 建议
如果需要以指针的形式存放循环变量，则可以显式地拷贝一次：
```go
for i := 0; i < 3; i++ {
    iCopy := i // Copy i into a new variable.
    out = append(out, &iCopy)
}
```
另一种解决方案是修改切片的类型，避免存储指针。

## 5 协程引用循环变量
### 1. 小测验
### 2. 解析
#### 1）循环变量是易变的
循环变量实际上只是一个普通的变量。

在 `for index, value := range xxx`语句中，每次循环时 index 和 value 都会被重新复制（并非生成新的变量）。

如果循环体中会启动协程（并且协程会使用循环变量），那么就需要格外注意了，因为很可能循环结束后协程才开始执行，此时所有协程的循环变量有可能已被改写（是否会改写取决于引用循环变量的方式）。

#### 2）循环变量需要绑定

### 3. 建议
当循环中引用循环变量时，如果需要启动并发并且引用循环变量，则需要格外留意变量是否已绑定。

## 6 recover 失效
### 1. recover 的使用误区
```go
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
```
`func IsPanic() bool`用来接收异常，返回值用来说明是否发生了异常。在`func UpdateTable()`函数中，使用 defer 判断最终应该提交还是回滚。

上面的代码看起来还算合理，但此处的 IsPanic() 再也不会返回 true，这不是 IsPanic() 函数的问题，而是其调用位置不对。

### 2. recover 失效的条件
上面的代码中 IsPanic() 失效了，其原因是违反了 recover 的一个限制，导致 recover() 失效（永远返回nil）。

一下三个条件会让 recover() 返回 nil：
- “panic”时指定的参数为 nil（一般 panic 语句如 panic("xxx failed...") ）；
- 当前协程没有发生 panic；
- recover 没有被 defer 函数直接调用。

前两条都比较容易理解，上述例子正式匹配了第 3 个条件。

recover() 必须被 defer 函数直接调用才有效，否则它永远返回 nil。在本例中，defer 函数先调用 IsPanic() 函数，在 IsPanic() 函数内的 recover() 是无法生效的，这是 recover() 生效的硬性限制。