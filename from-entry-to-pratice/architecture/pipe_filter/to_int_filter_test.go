package pipefilter

import (
	"reflect"
	"testing"
)

func TestToIntFilter(t *testing.T) {
	input := []string{"1", "2", "3"}
	expect := []int{1, 2, 3}
	tif := NewToIntFilter()
	ret, err := tif.Process(input)
	if err != nil {
		t.Fatal(err)
	}
	output, ok := ret.([]int)
	if !ok {
		t.Fatalf("Response type is %T, but the expected type is []int", output)
	}
	if !reflect.DeepEqual([]int{1, 2, 3}, output) {
		t.Errorf("Expected value is %v, but actual is %v", expect, output)
	}
}

func TestWrongInputForTIF(t *testing.T) {
	tif := NewToIntFilter()
	_, err := tif.Process([]string{"1", "2.2", "3"})
	if err == nil {
		t.Fatal("An error is expected for wrong input")
	}
	_, err = tif.Process([]int{1, 2, 3})
	if err == nil {
		t.Fatal("An error is expected for wrong input")
	}
}
