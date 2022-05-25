package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

type Cat struct {
	Animal // 由于内嵌结构体 Animal，从而产生一个隐式的同名字段。
}

type Dog struct {
	a Animal // 显式指定，与其他类型没有区别。
}

func EmbeddedFoo() {
	c := Cat{}

	// 可以直接访问 Animal 的字段和方法
	c.SetName("A")
	fmt.Printf("Name: %s\n", c.Name)

	// 也可以通过隐式声明的 Animal 字段来访问。
	c.Animal.SetName("a")
	fmt.Printf("Name: %s\n", c.Name)

	c.Name = "b"
	fmt.Printf("Name: %s\n", c.Name)
}

// 方法
type Student struct {
	Name string
}

// 作用于 Student 的拷贝对象，修改不会反映到原对象
func (s Student) SetName(name string) {
	s.Name = name
}

// 作用于 Student 的指针对象，修改会反映到原对象
func (s *Student) UpdateName(name string) {
	s.Name = name
}

func Receiver() {
	s := Student{}
	s.SetName("Rainbow")
	fmt.Printf("Name: %s\n", s.Name) // empty
	s.UpdateName("Rainbow")
	fmt.Printf("Name: %s\n", s.Name) // Rainbow
}

// Tag
type TypeMeta struct {
	Kind       string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	APIVersion string `json:"apiversion,omitempty" protobuf:"bytes,2,opt,name=apiversion"`
}

func PrintTag() {
	t := TypeMeta{}
	ty := reflect.TypeOf(t)

	for i := 0; i < ty.NumField(); i++ {
		fmt.Printf("Field: %s, Tag: %s\n", ty.Field(i).Name, ty.Field(i).Tag.Get("json"))
	}
}
