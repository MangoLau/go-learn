package gotest_test

import (
	"gotest"
	"testing"
)

// sub1 为子测试，只做加法测试
func sub1(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

// sub2 为子测试，只做加法测试
func sub2(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

// sub3 为子测试，只做加法测试
func sub3(t *testing.T) {
	var a = 1
	var b = 2
	var expected = 3

	actual := gotest.Add(a, b)
	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected: %d", a, b, actual, expected)
	}
}

func TestSub(t *testing.T) {
	// setup code

	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("A=3", sub3)

	// tear-down code
}
