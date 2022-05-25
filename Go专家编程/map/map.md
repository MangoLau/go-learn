## 特性浏览
### 1 操作方式
#### 1）初始化
字面量初始化
```go
func MapInitByLiteral() {
    m := map[string]int {
        "apple":  2,
        "banana": 3,
    }

    for k, v := range m {
        fmt.Printf("%s-%d\n", k, v)
    }
}
```

内置函数 `make()` 初始化
```go
func MapInitByMake() {
    m := make(map[string]int, 10)
    m["apple"] = 2
    m["banana"] = 3

    for k, v := range m {
        fmt.Printf("%s-%d\n", k, v)
    }
}
```

#### 2)增删改查
```go
func MapCURD() {
	m := make(map[string]string, 10)
	m["apple"] = "red"     // 添加。如果键“apple”不存在，则 map 会创建一个新的键值对并存储，等同于添加新的元素
	m["apple"] = "green"   // 修改
    // 在 map 为 nil 或指定的键不存在的情况下，delete()也不会报错，相当于空操作。
	delete(m, "apple")     // 删除
	v, exist := m["apple"] // 查询
	if exist {
		fmt.Printf("apple-%s\n", v)
	}
}
```

### 2 危险操作
#### 1）并发操作
map 操作不是原子的，这意味着多个协程同时操作 map 时有可能产生读写冲突，读写冲突会触发 panic 从而导致程序退出。

#### 2）空 map
未初始化的 map 的值为 nil，在向值为 nil 的 map 添加元素时会触发 panic，在使用时需要避免。

值为 nil 的 map，长度与空 map 一致
```go
func EmptyMap() {
	var m1 map[string]int
	m2 := make(map[string]int)

	fmt.Printf("len(m1) = %d\n", len(m1)) // 0
	fmt.Printf("len(m2) = %d\n", len(m2)) // 0
}
```

### 3 小结
初始化 map 时推荐使用内置函数 make() 并指定预估的容量。

修改键值对时，需要先查询指定键是否存在，否则 map 将创建新的键值对。

查询键值对时，最好检查键是否存在，避免操作零值。

避免并发读写 map，如果需要并发读写，则可以使用额外的锁（互斥锁、读写锁），也可以考虑使用标准库 sync 包中的 sync.Map。

