package tassertion

import (
	"math/rand"
	"runtime"
	"testing"
)

type SomeStruct struct {
	foo int
	bar int
}

//go:noinline
func barNoCheck(c any) int {
	return c.(SomeStruct).bar
}

//go:noinline
func barCheck(c any) int {
	if c, ok := c.(SomeStruct); ok {
		return c.bar
	}
	return 0
}

func sumNoCheck(arr []interface{}) int {
	var sum = 0
	for _, candidate := range arr {
		sum += barNoCheck(candidate)
	}
	return sum
}

func sumCheck(arr []interface{}) int {
	var sum = 0
	for _, candidate := range arr {
		sum += barCheck(candidate)
	}
	return sum
}

func BenchmarkSumNoCheck(b *testing.B) {
	arr := initTestData()
	for i := 0; i < b.N; i++ {
		sum := sumNoCheck(arr)
		runtime.KeepAlive(sum)
	}
}

func BenchmarkSumCheck(b *testing.B) {
	arr := initTestData()
	for i := 0; i < b.N; i++ {
		sum := sumCheck(arr)
		runtime.KeepAlive(sum)
	}
}

func initTestData() []any {
	arr := make([]any, 1000)
	for i := range arr {
		arr[i] = SomeStruct{foo: rand.Int(), bar: rand.Int()}
	}
	return arr
}
