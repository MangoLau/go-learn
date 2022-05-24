package main

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
