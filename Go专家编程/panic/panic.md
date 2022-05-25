## 工作机制
### 1. panic() 函数
panic() 是一个内置函数：
```go
func panic(v interface{})
```
它接受一个任意类型的参数，参数将在程序崩溃时通过另一个内置函数 `print(args ...Type)` 打印出来。如果程序返回途中任意一个 defer 函数执行力 recover()，那么该参数也是 recover() 的返回值。

panic 可由程序员显式地通过该内置函数触发，Go 运行时遇到诸如内存越界之类的问题时也会触发。

### 2. 工作流程
如果某协程执行过程中产生了 panic，那么程序将立即转向执行 defer 函数，当前函数中的 defer 执行完毕后将继续处理上层函数的 defer，当协程中所有 defer 处理完后，程序退出。

在 panic 的执行过程中有几个要点要明确：
- panic 会递归执行协程中所有的 defer，与函数正常退出时的执行顺序一致；
- panic 不会处理其他协程中的 defer；
- 当前协程中的 defer 处理完成后，触发程序退出。

如果 panic 在执行过程中（defer 函数中）再次发生 panic，程序将立即终止当前 defer 函数的执行，然后继续接下来的 panic 流程，只是当前 defer 函数中 panic 后面的语句就没有机会执行了。在这种情况下，把 defer 函数当成普通函数理解即可。

如果在 panic 的执行过程中任意一个 defer 函数执行力 recover()，那么 panic 的处理流程就会终止。