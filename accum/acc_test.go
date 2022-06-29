package accum

import (
	"math/rand"
	"runtime"
	"testing"
)

func BenchmarkAccumGeneric(b *testing.B) {
	arr := make([]byte, 1000)
	for i := range arr {
		arr[i] = byte(rand.Int())
	}

	for i := 0; i < b.N; i++ {
		sum := AccumulateGeneric(arr)
		runtime.KeepAlive(sum)
	}
}

func BenchmarkAccumAsm(b *testing.B) {
	arr := make([]byte, 1000)
	for i := range arr {
		arr[i] = byte(rand.Int())
	}

	for i := 0; i < b.N; i++ {
		sum := AccumulateAsm(arr)
		runtime.KeepAlive(sum)
	}
}

func TestAccumulate(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := make([]byte, rand.Int31n(10_000))
		for i := range arr {
			arr[i] = byte(rand.Int())
		}

		sumGeneric := AccumulateGeneric(arr)
		sumAmd64 := AccumulateAsm(arr)

		if sumGeneric != sumAmd64 {
			t.Logf("Expected %d, got %d", sumGeneric, sumAmd64)
			t.Fail()
		}
	}
}
