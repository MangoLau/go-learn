package concatstring

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const Numbers = 100

func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < Numbers; j++ {
			s = fmt.Sprintf("%v%v", s, i)
		}
	}
}

func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for j := 0; j < Numbers; j++ {
			s += strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkBytesBuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for j := 0; j < Numbers; j++ {
			buf.WriteString(strconv.Itoa(j))
		}
		_ = buf.String()
	}
	b.StopTimer()
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < Numbers; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
