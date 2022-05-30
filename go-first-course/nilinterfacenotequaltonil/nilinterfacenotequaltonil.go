package main

import (
	"errors"
	"fmt"
)

type QuackableAnimal interface {
	Quack()
}

type Duck struct{}

func (Duck) Quack() {
	println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
	println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
	println("bird quack!")
}

func AnimalQuackInForest(a QuackableAnimal) {
	a.Quack()
}

// ----
type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

// ----

type T struct {
	n int
	s string
}

func (T) M1() {}
func (T) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

// ----

func main() {
	// var err error = 1
	// fmt.Println(err)

	// var err error
	// err = errors.New("error1")
	// fmt.Printf("%T\n", err)

	// animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
	// for _, animal := range animals {
	// 	AnimalQuackInForest(animal)
	// }

	// err := returnsError()
	// if err != nil {
	// 	fmt.Printf("error occur: %+v\n", err)
	// }
	// fmt.Println("ok")

	var t = T{
		n: 17,
		s: "hello, interface",
	}
	var ei interface{}
	ei = t

	var i NonEmptyInterface
	i = t
	fmt.Println(ei)
	fmt.Println(i)
}
