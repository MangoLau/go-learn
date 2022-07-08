package autogrowing

import "testing"

const NumOfElems = 100000
const times = 1000

func TestAutoGrow(t *testing.T) {
	for i := 0; i < times; i++ {
		s := []int{}
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, NumOfElems)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestOverSizeInit(t *testing.T) {
	for i := 0; i < times; i++ {
		s := make([]int, 0, NumOfElems*5)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := []int{}
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkProperInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElems)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkOverSizeInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElems*5)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}
