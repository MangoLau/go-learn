package main

import "github.com/mangolau/instrument_trace"

func foo() {
	defer trace.Trace()()
	bar()
}

func bar() {
	defer trace.Trace()()
}

func main() {
	defer trace.Trace()()
	foo()
}
