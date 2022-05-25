package iota

const (
	bit0, mask0 = 1 << iota, 1<<iota - 1 // const 声明第 0 行，即 iota==0
	bit1, mask1                          // const 声明第 1 行，即 iota==1，表达式继承上面的语句
	_, _                                 // const 声明第 2 行，即 iota==2
	bit3, mask3                          // const 声明第 3 行，即 iota==3
)

