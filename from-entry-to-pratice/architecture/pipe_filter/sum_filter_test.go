package pipefilter

import "testing"

func TestSumFilter(t *testing.T) {
	sf := NewSumFilter()
	ret, err := sf.Process([]int{1, 2, 3})
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("expected value is %d, actual is %d", 6, ret)
	}
}

func TestWrongInputForSumFilter(t *testing.T) {
	sf := NewSumFilter()
	_, err := sf.Process([]float32{1.1, 2.2, 3.1})

	if err == nil {
		t.Fatal("An error is expected.")
	}
}
