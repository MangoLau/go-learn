# 接口：为什么 nil 接口不等于 nil？
## 接口的静态特性与动态特性
接口的静态特性体现在接口类型变量具有静态类型，比如 `var err error` 中的变量 err 的静态类型为 error。拥有静态类型，那就意味着编译器会在编译阶段对所有接口类型变量的赋值操作进行类型检查，编译器会检查右值的类型是否实现了该接口方法集合中的所有方法。如果不满足，就会报错：
```go
var err error = 1 // cannot use 1 (constant of type int) as type error in variable declaration: int does not implement error (missing Error method)
```

而接口的动态特性，就体现在接口类型变量在运行时还存储了右值的真实类型信息，这个右值的真实类型被称为接口类型变量动态类型。
```go
var err error
err = errors.New("error1")
fmt.Printf("%T\n", err) // *errors.errorString
```
这个示例通过 errors.New 构建了一个错误值，赋值给了 error 接口类型变量 err，并通过 fmt.Printf 函数输出接口类型变量 err 的动态类型为 *errors.errorString。

接口“动静皆备”的特性的好处：

首先，接口类型变量在程序运行时可以被赋值为不同的动态类型变量，每次赋值后，接口类型变量中存储的动态类型信息都会发生变化，这让 Go 语言可以像动态语言那样拥有使用 Duck Typing（鸭子类型）的灵活性。所谓鸭子类型，就是值某类型所表现出的特性（比如是否可以作为某接口类型的右值），不是由其基因（比如 C++ 中的父类）决定的，而是由类型所表现出来的行为（比如类型拥有的方法）决定的。

```go
type QuackableAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInForest(a QuackableAnimal) {
	a.Quack()
}

func main() {
	animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
	for _, animal := range animals {
		AnimalQuackInForest(animal)
	}
}
```

这个例子中，我们用接口类型 QuackableAnimal 来代表具有“会叫”这一特征的动物，而 Duck、Bird 和 Dog 类型各自都具有这样的特征，于是我们可以将这三个类型的变量赋值给 QuackableAnimal 类型变量，只是因为他们表现出了 QuackableAnimal 所要求的特征罢了。

不过，与动态语言不同的是，Go 接口还可以保证“动态特性”使用时的安全性。比如，编译器在编译期就可以捕捉到将 int 类型变量传给 QuackableAnimal 接口类型变量这样的明显错误，绝不会让这样的错误遗漏到运行时才被发现。

### nil error 值 != nil
```go
type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func main() {
    err := returnsError()
    if err != nil {
        fmt.Printf("error occur: %+v\n", err)
    }
    fmt.Println("ok")
}
```

实际输出：
```
error occur: <nil>
```

### 接口类型变量的内部表示
```go
// $GOROOT/src/runtime/runtime2.go
type iface struct {
    tab *itab
    data unsafe.Pointer
}

type eface struct {
    _type *_type
    data unsafa.Pointer
}
```

我们看到，在运行时层面，接口类型变量有两种内部表示： iface 和 eface，这两种表示分别用于不同的接口类型变量：
- eface 用于表示没有方法的空接口（empty interface）类型变量，也就是 interface{} 类型的变量；
- iface 用于表示其余拥有方法的接口 interface 类型变量。

这两个结构的共同点是它们都有两个指针字段，并且第二个指针字段的功能相同，都是只想当前赋值给该接口类型变量的动态类型变量的值。

它们的不同点在于 eface 表示的空接口类型并没有方法列表，因此它的第一个指针字段指向一个 _type 类型结构，这个结构为该类型变量的动态类型信息。

而 iface 除了要存储动态类型信息之外，还要存储接口本身的信息（接口的类型信息、方法列表信息等）以及动态类型所实现的方法的信息，因此 iface 的第一个字段指向一个 itab 类型结构。

### 接口类型装箱（boxing）原理
装箱（boxing）是编程语言领域的一个基础概念，一般是指把一个值类型转换成引用类型。

在 Go 语言中，将任意类型赋值给一个接口类型变量也是装箱操作。接口类型的装箱实际就是创建一个 eface 或 iface 的过程。