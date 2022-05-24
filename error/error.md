## 基础 error
### 1. error 接口
error 是一种内建的接口类型，内建意味着不要“import”任何包就可以直接使用，使用起来就像 int、string 一样自然。

### 2. 创建 error
标准库提供了两种创建 error 的方法：
- errors.New()
- fmt.Errorf()

性能对比
`fmt.Errorf()`适用于需要格式化输出错误字符串的场景，如果不需要格式化字符串，则建议直接使用`errors.New()`。

### 3. 自定义 error
任何实现 error 接口的类型都可以称为 error。

### 4. 异常处理
针对 error 而言，异常处理包括如何检查错误、如何传递错误。

#### 1）检查 error
最常见的检查 error 的方式是与 nil 值进行比较：
```go
if err != nil {
    // somethin went wrong
}
```

有时也会与一些预定义的 error 进行比较：
```go
// 标准库 os 包中定义了一些常见的错误
// Errpermission = errors.New("permission denied")

if err == os.ErrPermission {
    // permission denied
}
```

由于任何实现了 error 接口的类型均可以作为 error 来处理，所以往往也会使用类型断言来检查 error：
```go
func AssertError(err error) {
    if e, ok := err.(*os.PahtError); ok {
        fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v", e.Op, e.Path, e.Err)
    }
}
```
上面代码中的断言，如果 err 是 os.PathError 类型，则可以使用 e 来访问 os.PathError 中的成员。在下面的 Example 测试中，err1 是 os.PathError 类型，断言为真，而 err2 的断言为假。
```go
// go test ./errors -run=ExampleAssertError
func ExampleAssertError() {
    err1 := &os.PathError {
        Op: "write",
        Path: "/root/demo.txt",
        Err: os.ErrPermission,
    }
    AssertError(err1)

    err2 := fmt.Errorf("not an os.PathError")
    AssertError(err2)

    // Output:
    // it's an os.PathError, operation: write, path: /root/demo.txt, msg: permission denied
}
```

#### 2）传递 error
在一个函数中收到一个 error，往往需要附加一些上下文信息再把 error 继续向上层抛。
最常见的添加附加上下文信息的方法是使用 fmt.Errorf():
```go
if err != nil {
    return fmt.Errorf("decompress %v: %v", name, err)
}
```
这种抛出 error 有一个糟糕的问题，那就是原 error 信息和附加的信息被糅合到一起了。比如下面的函数，就会把 os.ErrPermission 和附加信息糅合到一起：
```go
func WriteFile(fileName string) {
    if fileName == "a.txt" {
        return fmt.Errorf("write file error: %v", os.ErrPermission)
    }

    return nil
}
```
这样，在下面的 Example 测试中，就无法判断 err 是否是 os.ErrPermission 值了：
```go
// go test ./errors -run=ExampleWriteFile
func ExampleWriteFile() {
    err := WriteFile("a.txt")
    if err == os.ErrPermission {
        fmt.Printf("permission denied")
    }

    // Output:
    //
}
```
为了解决这个问题，我们可以自定义 error 类型，就像 os.PathError 那样，上下文信息与原 error 信息分开存放：
```go
type PathError struct {
    Op string // 上下文
    Path string // 上下文
    Err error // 原 error
}
```
这样，对于一个 os.PathError 类型的 error，我们可以检测它到底是不是一个权限不足的错误：
```go
if e, ok := err.(*os.PathError); ok && e.Err == os.ErrPermission {
    fmt.Printf("permission denied")
}
```

### 5. 小结

## 链式 error
Go 1.13 中针对 error 的优化，包括：
- 新的 error 类型 wrapError；
- 增强了 fmt.Errorf() 以便通过 %w 动词创建 wrapError；
- 引入了 errors.Unwrap()  以便拆解 wrapError；
- 引入了 errors.Is() 用于检查 error 链条中是否包含指定的错误值；
- 引入了 errors.As() 用于检查 error 链条中是否包含指定的错误类型。

### 1. wrapError

### 2. fmt.Errorf()

使用 fmt.Errorf() 生成 wrapError 有两个限制：
- 每次生成 wrapError 时只能使用一次 %w 动词；
- %w 动词只能匹配实现了 error 接口的参数。

### 3. errors.Unwrap()
如果参数 err 没有实现 Unwrap() 函数，则说明是基础 error，直接返回 nil，否则调用原 err 实现的 Unwrap() 函数并返回。

### 4. errors.Is()
errors.Is()用于检查特定的 error 链中是否包含指定的 error 值。
```go
func ExampleIs() {
    err1 := fmt.Errorf("write file error: %w", os.ErrPermission)
    err2 := fmt.Errorf("write file error: %w", err1)
    
    if errors.Is(err2, os.ErrPermission) {
        fmt.Printf("permission denied")
    }
    // Output:
    // permission denied
}
```

### 5. errors.As()
errors.As() 用于从一个 error 链中查找是否有指定类型出现，如有，则把 error 转换成该类型。

### 6. 小结