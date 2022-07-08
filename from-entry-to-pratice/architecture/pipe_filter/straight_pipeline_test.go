package pipefilter

import "testing"

func TestStraightPipeline(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()
	sp := NewStraightPipeline("pipeline", splitFilter, toIntFilter, sumFilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("expect 6, actual %d", ret)
	}
}
