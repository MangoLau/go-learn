## 单元测试
- 测试文件名必须以 _test.go 结尾；
- 测试函数名必须以 TestXxx 开始；
- 在命令行下使用 go test 即可启动测试。

## 性能测试
- 文件名必须以 _test.go 结尾；
- 函数名必须以 BenchmarkXxx 开始；
- 使用 `go test -bench=.` 命令即可开始性能测试。

## 示例测试
- 示例测试函数名需要以 Example 开头；
- 检测单行输出格式为 `// Output: <期望字符串>`；
- 检测多行输出格式为 `// Output: \n <期望字符串> \n <期望字符串>`，每个期望值字符串占一行；
- 检测无序输出格式为 `// Unordered output: \n <期望字符串> \n <期望字符串>`，每个期望值字符串占一行；
- 测试字符串时会自动忽略字符串前后的空白字符；
- 如果测试函数中没有 Output 标识，则该测试函数不会被执行；
- 执行测试可以使用 `go test`，此时该目录下的其他测试文件也会一并执行；
- 执行测试可以使用 `go test <xxx_test.go>`，此时仅执行特定文件中的测试函数。

## 进阶测试
### 子测试
- 子测试适用于单元测试和性能测试；
- 子测试可以控制并发；
- 子测试提供一种类似 table-driven 风格的测试；
- 子测试可以共享 setup 和 tear-down。

### Main 测试
```go
// TestMain 用于主动执行各种测试，可以在测试前后做 setup 和 tear-down 操作
func TestMain(m *testing.M) {
    println("TestMain setup.")

    retCode := m.Run() // 执行测试，包括单元测试、性能测试和示例测试

    println("TestMain tear-down.")

    os.Exit(retCode)
}
```

如果所有测试通过，则 `m.Run()`返回 0，如果`m.Run()`返回 1，则代表测试失败。

单元测试函数需要传递一个 testing.T 类型的参数，而性能测试函数需要传递一个 testing.B 类型的参数，该参数可用于控制测试的流程，比如标记测试失败等。