# 语法糖
用于表示编程语言中特定类型的语法这些语法对语言对功能没有影响，但是更方便程序员使用。

语法糖也称为糖语法，这些语法不仅不会影响语言的功能，编译后的结果跟不使用语法糖也一样。语法糖有可能让代码的编写变得简单，也有可能让代码的可读性更高，但也有可能让使用者“步入陷阱”。

## 1. 简短变量声明符
可以使用关键字 var 或直接使用简短变量声明符（:=）声明变量。后者使用得更频繁一些，尤其是在接收函数返回值的场景中，不必使用 var 声明一个变量再用变量接收函数返回值，使用“:=”可以“一步到位”。比如：
```go
a := foo1()
a, b := foo2()
```

### 规则
#### 1. 规则一：多变量赋值可能会重新声明
如果两个变量中的一个再次出现在“:=”的左侧就会被重新声明，像下面这样：
```go
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)
```
offset 被重新声明。

重新声明并没有什么问题，它并没有引入新的变量，只是把变量的值改变了，但要明白，这是 Go 提供的一个语法糖。
- 当“:=”左侧存在新变量时（如 field2），已声明的变量（如 offset）会被重新声明，不会有其他额外副作用。
- 当“:=”左侧没有新变量是不允许的，编译会提示`no new variables on left side of :=`

我们所说的重新声明不会引入问题要满足一个前提，那就是变量声明要在同一个作用域中出现。如果出现在不同的作用域中，则很可能创建了新的同名变量，同一函数不同作用域的同名变量往往不是预期做法，很容易引入陷阱。

#### 2. 规则二：不能用于函数外部
简短变量声明符只能用于函数中，使用“:=”来声明和初始化全局变量是行不通的。
比如，像下面这样：
```go
package sugar
import fmt

rule := "Short variable declarations" // syntax error: non-declaration statement outside function body
```
`syntax error: non-declaration statement outside function body`表示非声明语句不能出现在函数外部。可以理解成“:=”实际上会拆分成两个语句，即声明和赋值。赋值语句不能出现在函数外部。

#### 3. 变量作用域问题
如果使用“:=”过于随意，那么有可能在多个作用域中声明了同名变量而不自知，从而引发故障。
```go
func Redeclare() {
    field, err := nextField() // 1 号 err

    if field == 1 {
        field, err := nextField() // 2 号 err
        newField, err := nextField() // 3 号 err
        ...
    }
    ...
}
```
注意上面声明的三个 err 变量。
- 2 号 err 与 1 号 err 不属于同一个作用域，“:=”声明了新的变量，所以 2 号 err 与 1 号 err 是两个变量。
- 2 号 err 与 3 号 err 属于同一个作用域，“:=” 重新声明了 err 但没有创建新的变量，所以 2 号 err 与 3 号 err 是同一个变量。

如果误把 2 号 err 与 1 号 err 混淆，就很容易产生意想不到的错误。

# 可变参函数
可变参函数是指函数的某个参数可有可无，即这个参数的个数可以是 0 或多个。声明可变参数函数的方式是在参数类型前加上“...”前缀。

比如 fmt 包中的 Println：
```go
func Println(a ...interface{})
```

## 1. 函数特征
```go
func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}
```
Greeting 函数负责给指定的人打招呼，其参数 who 为可变参数。
这个函数几乎把可变参函数的的特征全部表现出来了：
- 可变参数必须在参数列表的尾部，即最后一个（如放前面会引起编译时歧义）；
- 可变参数在函数内部是作为切片来解析的（可以使用 range 遍历）；
- 可变参数可以不填，不填时函数内部当成 nil 切片处理；
- 可变参数必须是相同的类型（如果需要是不同的类型则可以定义为 interface{} 类型）。

## 2. 使用举例
### 1）不传值
调用可变参数时，可变参部分可以不传值的，例如：
```go
func ExampleGreetingWithoutParameter() {
    sugar.Greeting("nobody")
    // Output:
    // Nobody to say hi.
}
```
这里没有传递第二个参数。可变参数不传递值时默认为 nil。

### 2）传递多个参数
调用可变参函数时，可变参数部分可以传递多个值，例如：
```go
func ExampleGreetingWithParameter() {
    sugar.Greeting("hello:", "Joe", "Anna", "Eileen")
    // Output:
    // hello: Joe
    // hello: Anna
    // hello: Eileen
}
```
可变参数可以有多个，多个参数将生成一个切片传入，函数内部按照切片来处理。

### 3）传递切片
调用可变参函数时，可变参数部分可以直接传递一个切片。参数部分需要使用 slice... 来表示切片，例如：
```go
func ExampleGreetingWithSlice() {
    guest := []string{"Joe", "Anna", "Eileen"}
    sugar.Greeting("hello:", guest...)
    // Output:
    // hello: Joe
    // hello: Anna
    // hello: Eileen
}
```
此时需要注意的是，切片传入时不会生成新的切片，也就是说函数内部使用的切片与传入的切片共享相同的存储空间。说得再直白一点就是，如果函数内部修改了切片，则可能影响外部调用的函数。

## 3. 小结
- 可变参数必须要位于函数列表尾部；
- 可变参数是被当作切片来处理的；
- 函数调用时，可变参数可以不填；
- 函数调用时，可变参数可以填入切片。