##  约法三章
defer 不仅可以用于资源释放，也可以用于流程控制和异常处理，但 defer 关键字只能作用于函数或函数调用。

defer 关键字后接一个匿名函数：
```go
defer func() {
    fmt.Print("Hello World!")
}()
```

defer 关键字后接一个函数调用：
```go
file, err := os.Open(name)
if err != nil {
    return nil, err
}
defer file.Close()
```

### 1. 使用场景
#### 1）释放资源
```go
m.mutex.Lock()
defer m.mutex.Unlock()
```
defer 常用语关闭文件句柄、数据库连接、停止定时器 Ticker 及关闭管道等资源清理场景。

#### 2）流程控制
```go
var wg wait.Group
defer wg.Wait()
// ...
```
defer 也常用于控制函数执行顺序，比如配合 wait.Group 实现等待协程退出。

#### 3）异常处理
```go
defer func() { recover() }() // Actually eat panics
```
defer 也常用于处理异常，与 recover 配合可以消除 panic。另外，recover 只能用于 defer 函数中。

### 2. 行为规则
#### 1）规则一：延迟函数的参数在 defer 语句出现时就已经确定了
```go
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

defer 语句中的 fmt.Println()参数 i 值在 defer 出现时就已经确定了，实际上是复制了一份。后面对变量 i 的修改不会影响 fmt.Println() 函数的执行，仍然打印 0。

注意，对于指针类型参数，此规则仍然适用，只不过延迟函数的参数是一个地址值，在这种情况下，defer 后面的语句对变量的修改可能会影响延迟函数。

#### 2）规则二：延迟函数按后进先出的顺序执行，即先出现的 defer 最后执行
每申请到一个用完需要释放的资源时，立即定义一个 defer 来释放资源是一个很好的习惯。

#### 3）规则三：延迟函数可能操作主函数的具名返回值
定义 defer 的函数（下称主函数）可能有返回值，返回值可能有名字（具名返回值），也可能没有名字（匿名返回值），延迟函数可能会影响返回值。

（1）函数返回过程。

关键字 return 不是一个原子操作，实际上 return 只代表汇编指令 ret，即跳转程序执行。比如语句 `return i`，实际上分两步执行，即先将 i 值存入栈中作为返回值，然后执行跳转，而 defer 的执行时机正是在跳转前，所以说 defer 执行时还是有机会操作返回值的。
```go
func deferFuncReturn() (result int) {
    i := 1

    defer func() {
        result++
    }()
    return i
}
```
该函数的 return 语句可以拆分成下面两行：
```go
result = i
return
```
而延迟函数的执行正是在 return 之前，即加入 defer 后的执行过程如下：
```go
result = i
result++
return
```
所以上面的函数实际返回 i++ 值。

（2）主函数拥有匿名返回值，返回字面值。

一个主函数拥有一个匿名返回值，返回时使用字面值，比如返回 1、2、Hello 这样的值，这种情况下 defer 语句是无法操作返回值的。
一个返回字面值的函数如下：
```go
func foo() int {
	var i int

	defer func() {
		i++
	}()

	return 1
}
```
上面的 return 语句直接把 1 写入栈中作为返回值，延迟函数无法操作该返回值，所以就无法影响返回值。

（3）主函数拥有匿名返回值，返回变量。

一个主函数拥有一个匿名返回值，返回本地或全局变量，这种情况下 defer 语句可以引用返回值，但不会改变返回值。
```go
func bar() int {
	var i int

	defer func() {
		i++
	}()

	return i
}
```
上面的函数返回一个局部变量，同时 defer 函数也会操作这个局部变量。对于匿名返回值来说，可以假定仍然有一个变量存储返回值，假定返回值变量为 anony，上面的返回语句可以拆分为以下过程：
```go
anony = i
i++
return
```

（4）主函数拥有具名返回值

主函数声明语句中带名字的返回值会被初始化为一个局部变量，函数内部可以像使用局部变量一样使用该返回值。如果 defer 语句操作该返回值，则可能改变返回结果。
```go
func change() (ret int) {
	defer func() {
		ret++
	}()
	return 0
}
```
上面的函数拆解出来如下所示：
```go
ret = 0
ret++
return
```
函数真正返回前，在 defer 中对返回值做了 +1 操作，所以函数最终返回 1.

## 实现原理
### 1. 数据结构
### 2. defer 的创建和执行
- deferproc()：负责把 defer 函数处理成 _defer 实例，并存入 goroutine 中的链表；
- deferreturn()：负责把 defer 从 goroutine 链表中的 defer 实例取出并执行。

### 3. 小结
- defer 定义的延迟函数参数在 defer 语句出现时就已经确定了；
- defer 定义的顺序与实际的执行顺序相反；
- return 不是原子操作，执行过程是：保存返回值（若有）-> 执行 defer（若有） -> 执行 ret 跳转；
- 申请资源后立即使用 defer 关闭资源是一个好习惯。

## 性能优化
### 1. 堆 defer
编译器将 defer 语句编译成一个 deferproc() 函数调用，然后运行时执行 deferproc 函数，deferproc 函数会根据 defer 语句生成一个 _defer（运行时内部数据结构名）实例并插入 goroutine 的 _defer 链表头部。同时编译器还会在函数尾部插入 deferreturn 函数，deferreturn 函数会逐个去除 _defer 实例并执行。

运行时协程 g 的数据结构如下：
```go
type g struct {
    ...
    _defer *_defer
    ...
}
```

堆 defer 的特点是新创建的 defer 节点存储在堆中，deferproc 函数会将被延迟的函数组成一个 _defer 实例并复制到 _defer 节点中，deferreturn 函数消费完 _defer 后，再将节点销毁。

与其后出现的两种 defer 类型相比较，堆 defer 的痛点主要在于频繁的堆内存分配及释放，性能稍差。

### 2. 栈 defer
栈 defer 正是为了提高堆 defer 的内存使用效率而引入的，编译器将尽量将 defer 语句编译成一个 deferprocStack() 函数调用，deferprocStack() 的工作机制与 deferproc()  类似，区别在于编译器会直接在栈上预留 _defer 的存储空间，deferprocStack() 不需要再分配空间。deferprocStack() 仍然需要将 _defer 插入协程 g 的 _defer 链表中。

此时，运行时 _defer 的数据结构中引入了成员变量 heap 标记是否存储于堆中：
```go
type _defer struct {
    ...
    heap bool
    ...
}
```
在函数结尾处，编译器仍然会插入 deferreturn() 函数，该函数的执行过程与堆 defer 类似。所不同的是，执行结束后不需要再释放内存了。

编译器会尽可能地把 defer 语句编译成栈类型，但由于栈空间也有限，并不能把所有的 defer 都存储在栈中，所以还需要保留堆 defer。

### 3. 开放编码 defer
以下场景下 defer 语句不能被处理成开放编码类型：
- 编译时禁用了编译器优化，即 `-gcflags="-N -1"`；
- defer 出现在循环语句中；
- 单个函数中出现了 8 个以上，或者 return 语句的个数和 defer 语句的个数乘积超过了 15。

编程 Tips：
- 单个函数中如果存在过多的 defer，那么可以考虑拆分函数；
- 单个函数中如果存在过多的 return 语句，那么需要控制 defer 的使用数量；
- 在循环中使用 defer 语句需要谨慎。