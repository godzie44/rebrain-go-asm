package main

import (
	"math/rand"
	"testing"
)

func AccumulateGeneric(src []byte) int {
	cnt := 0
	for _, srcByte := range src {
		cnt += int(srcByte)
	}
	return cnt
}

func TestAccumulate(t *testing.T) {
	for i := 0; i < 10; i++ {
		arr := make([]byte, 1008)
		for i := range arr {
			arr[i] = byte(rand.Int())
		}

		sumGeneric := AccumulateGeneric(arr)
		sumAvo := int(AccumulateAvo(arr))

		if sumGeneric != sumAvo {
			t.Logf("Expected %d, got %d", sumGeneric, sumAvo)
			t.Fail()
		}
	}
}
