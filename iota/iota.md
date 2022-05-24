iota 常用语 const 表达式中，其值是从 0 开始的，const 声明块中每增加一行，iota 值自增 1。

iota 代表了 const 声明块的行索引（下标从 0 开始）。
const 声明还有一个特点，即如果为常量指定了一个表达式，但后续的常量没有表达式，则继承上面的表达式。

```go
const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // const 声明第 0 行，即 iota==0
	bit1, mask1                          // const 声明第 1 行，即 iota==1，表达式继承上面的语句
	_, _                                 // const 声明第 2 行，即 iota==2
	bit3, mask3                          // const 声明第 3 行，即 iota==3
)
```

- 第 0 行的表达式展开即 `bit0, mask0 = 1 << 0, 1<<0 - 1`，所以 `bit==1,mask0==0`；
- 第 1 行没有指定表达式继承第一行，即`bit1, mask1 = 1 << 1, 1<<1 - 1`，所以 `bit==2,mask0==1`；
- 第 2 行没有定义常量；
- 第 3 行没有指定表达式继承第一行，即`bit3, mask3 = 1 << 3, 1<<3 - 1`，所以 `bit==8,mask0==7`。