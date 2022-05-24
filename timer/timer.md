# 一次性定时器（Timer）

## 快速开始
Timer 是一种单一事件的定时器，即经过指定的时间后触发一个事件，这个事件通过其本身提供的 channel 进行通知。之所以叫单一事件，是因为 Timer 值执行一次就结束，这也是 Timer 与 Ticker 最重要的区别。

### 使用场景
#### 1）设定超时时间
```go
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}
```

#### 2）延迟执行某个方法
```go
func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
	}
}
```

### 3. Timer 对外接口
#### 1）创建定时器
```go
func NewTimer(d Duration) *Timer
```

#### 2）停止定时器
```go
func (t *Timer) Stop() bool
```
其返回值代表定时器有没有超时。
- true：定时器超时前停止，后续不会再发送事件；
- false：定时器超时后停止。

实际上，停止计时器意味着通知系统守护协程移除该定时器。

#### 3）重置定时器
```go
func (t *Timer) Reset(d Duration) bool
```
实际上，重置定时器意味着通知系统守护协程移除该定时器，重新设定时间后，再把定时器交给守护协程。

### 4. 简单接口
#### 1）After()
有时我们就是想等待指定的时间，没有提前停止定时器的需求，也没有复用该定时器的需求，那么可以使用匿名的定时器。

使用 `func After(d Duration) <-chan Time`方法创建一个定时器，并返回定时器的管道：
```go
func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	log.Println(time.Now())
}
```

#### 2）AfterFunc()
原型：
```go
func AfterFunc(d Duration, f func()) *Timer
```
该方法在指定时间到来后会执行函数 f。
```go
func AfterFuncDemo() {
	log.Println("AfterFuncDemo strat: ", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})
	time.Sleep(2 * time.Second) // 等待协程退出
}
```
time.AfterFunc() 是异步执行的，所以需要函数最后“sleep”等待指定的协程退出，否则可能函数结束时协程还未执行。

### 5. 小结
- time.NewTimer(d)：创建一个 Timer；
- time.Stop()：停止当前 Timer；
- time.Reset(d)：重置当前 Timer。