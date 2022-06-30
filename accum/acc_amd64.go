//go:build amd64

package accum

import (
	"reflect"
	"unsafe"
)

//go:noescape
func accumulateAsm(src unsafe.Pointer, nByte int) int

func AccumulateGeneric(src []byte) int {
	cnt := 0
	for _, srcByte := range src {
		cnt += int(srcByte)
	}
	return cnt
}

func AccumulateAsm(src []byte) int {
	if len(src) < 16 {
		return AccumulateGeneric(src)
	}

	nSrcByte := len(src)
	srcHeader := (*reflect.SliceHeader)(unsafe.Pointer(&src))

	return accumulateAsm(unsafe.Pointer(srcHeader.Data), nSrcByte-nSrcByte%16) +
		AccumulateGeneric(src[nSrcByte-nSrcByte%16:])
}
