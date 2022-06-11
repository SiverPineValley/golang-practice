package Package

import (
	common "golang-practice/src/Common"
	"strconv"
	"testing"
)

func BenchmarkLenForString(b *testing.B) {
	b.StopTimer()
	v := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		v[i] = strconv.Itoa(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		common.Len(v[i%1000000])
	}
}
