package pipefilter

import (
	"reflect"
	"testing"
)

func TestSpliteFilter(t *testing.T) {
	expected := []string{"1", "2", "3"}
	sf := NewSplitFilter(",")
	output, err := sf.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	parts, ok := output.([]string)
	if !ok {
		t.Fatalf("Response type is %T, but the expected type is string", parts)
	}
	if !reflect.DeepEqual(expected, parts) {
		t.Errorf("Expected value is {\"1\",\"2\",\"3\"}, but actual is %v", parts)
	}
}

func TestWrongInput(t *testing.T) {
	sf := NewSplitFilter(",")
	_, err := sf.Process(123)
	if err == nil {
		t.Fatal("An error is expected.")
	}
}
