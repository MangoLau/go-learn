## 逃逸分析
逃逸分析（Escape analysis）是指由编译器决定内存分配的位置，不需要程序员指定。

- 如果分配在栈中，则函数执行结束后可自动将内存回收；
- 如果分配在堆中，则函数执行结束后可交给 GC （垃圾回收）处理。

### 1. 逃逸策略
在函数中神器新的对象时，编译器会根据该对象是否被函数外部引用来决定是否逃逸：
- 如果函数外部没有饮用，则优先放到栈中；
- 如果函数外部存在引用，则必定放到堆中。

注意，对于仅在函数内部使用的变量，也有可能放到堆中，比如内存过大超过栈道存储能力。

### 2. 逃逸场景
#### 1）指针逃逸
Go 可以返回局部变量指针，这其实是一个典型的变量逃逸案例，
```go
type Student struct {
	Name string
	Age  int
}

func StudentRegister(name string, age int) *Student {
	s := new(Student) // 局部变量 s 逃逸到堆中

	s.Name = name
	s.Age = age

	return s
}

func main() {
	StudentRegister("Jim", 18)
}
```
函数 StudentRegister() 内部的 s 为局部变量，其值通过函数返回值返回，s 本身为一个指针，其指向的内存地址不会是栈而是堆，这就是典型的逃逸案例。

通过编译参数 `-gcflag=-m` 可以查看编译过程中的逃逸分析过程：
```shell
go build -gcflags=-m

# cache-manager
./cache-manager.go:8:6: can inline StudentRegister
./cache-manager.go:17:6: can inline main
./cache-manager.go:18:17: inlining call to StudentRegister
./cache-manager.go:8:22: leaking param: name
./cache-manager.go:9:10: new(Student) escapes to heap
./cache-manager.go:18:17: new(Student) does not escape
```

#### 2）栈空间不足逃逸
当栈空间不足以存放当前对象或无法判断当前切片长度时会将对象分配到堆中。

#### 3）动态逃逸类型
很多函数的参数为 interface 类型，比如 `fmt.Println(a ...interface{})`，编译期间很难确定其参数的具体类型，也会产生逃逸。

#### 4）闭包饮用对象逃逸

### 3. 小结
- 栈上分配内存比在堆中分配内存有更高的效率；
- 栈上分配的内存不需要 GC 处理；
- 堆上分配的内存使用完毕会交给 GC 处理；
- 逃逸分析的目的是决定分配地址是栈还是堆；
- 逃逸分析在编译阶段完成。