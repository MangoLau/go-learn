package gotest_test

import "gotest"

// 检测但行输出
func ExampleSayHello() {
	gotest.SayHello()
	// OutPut: Hello World
}

// 检测多行输出
func ExampleSayGoodbye() {
	gotest.SayGoodbye()
	// OutPut:
	// Hello,
	// goodbye
}

func ExamplePrintNames() {
	gotest.PrintNames()
	// Unordered outPut:
	// Jim
	// Bob
	// Tom
	// Sue
}
