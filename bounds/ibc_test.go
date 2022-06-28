package bounds

import (
	"math/rand"
	"runtime"
	"testing"
)

//go:noinline
func first1000Sum_v1(arr []int) int {
	z := 0
	for i := 1000; i > 0; i-- {
		z += arr[i]
	}

	return z
}

//go:noinline
func first1000Sum_v2(arr []int) int {
	z := arr[1000]
	for i := 999; i > 0; i-- {
		z += arr[i]
	}

	return z
}

func BenchmarkSum_v1(b *testing.B) {
	arr := make([]int, 1500)
	for i := 0; i < 1000; i++ {
		arr[i] = rand.Int()
	}

	for i := 0; i < b.N; i++ {
		s := first1000Sum_v1(arr)
		runtime.KeepAlive(arr)
		runtime.KeepAlive(s)
	}
}

func BenchmarkSum_v2(b *testing.B) {
	arr := make([]int, 1500)
	for i := 0; i < 1000; i++ {
		arr[i] = rand.Int()
	}

	for i := 0; i < b.N; i++ {
		s := first1000Sum_v2(arr)
		runtime.KeepAlive(arr)
		runtime.KeepAlive(s)
	}
}
