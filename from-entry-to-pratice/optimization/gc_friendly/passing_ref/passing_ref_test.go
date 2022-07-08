package passingref_test

import (
	"testing"
)

const NumOfElems = 1000

type Content struct {
	Detail [10000]int
}

func withValue(arr [NumOfElems]Content) int {
	// fmt.Println(arr[2].Detail[2])
	return 0
}

func withReference(arr *[NumOfElems]Content) int {
	// fmt.Println(arr[2].Detail[2])
	return 0
}

func TestFn(t *testing.T) {
	var arr [NumOfElems]Content
	withValue(arr)
	withReference(&arr)
}

func BenchmarkPassingArrayWtihValue(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withValue(arr)
	}
	b.StopTimer()
}

func BenchmarkPassingArrayWithRef(b *testing.B) {
	var arr [NumOfElems]Content
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		withReference(&arr)
	}
	b.StopTimer()
}
