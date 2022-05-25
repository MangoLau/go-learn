每个 interface 类型代表一个特定的方法集，方法集中的方法称为接口。
```go
type Animal interface {
    Speak() string
}
```
Animal 就是一个接口类型，其包含一个 Speak() 方法。

### 1）interface 变量
就像任何其他类型一样，我们也可以声明 interface 类型的变量。比如：
```go
var animal Animal
```
上面的 animal 变量的值为 nil，它没有赋值

### 2）实现接口
任何类型只要实现了 interface 类型的所有方法，就可以声称该类型实现了这个接口，该类型的变量就可以存储到 interface 变量中。
```go
type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow"
}
```
结构体 Cat 实现了 Speak() 方法，结构体 Cat 的变量就可以存储到 animal 变量中：
```go
var animal Animal
var cat Cat
animal = cat
```

事实上，interface 变量可以存储任意实现了该接口类型的变量。

### 3）复合类型
interface 类型的变量在存储某个变量时会同时保存变量类型和变量值。

### 4）空 interface
空 interface 是一种非常特殊的 interface 类型，它没有指定任何方法集，如此一来，任意类型都可以声称实现了空接口，那么接口变量也就可以存储任意变量