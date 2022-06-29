//go:generate go run main.go -out avo_amd64.s -stubs avo_amd64.go

package main

import (
	. "github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/operand"
)

func main() {
	TEXT("AccumulateAvo", NOSPLIT, "func(b []byte) uint64")
	Doc("Accumulate vector")

	n := Load(Param("b").Len(), GP64())
	ptr := Load(Param("b").Base(), GP64())

	x0 := XMM()
	x1 := XMM()
	PXOR(x0, x0)
	PXOR(x1, x1)

	x2 := XMM()

	Label("loop")

	CMPQ(n, operand.Imm(0))
	JE(operand.LabelRef("done"))

	MOVOU(operand.Mem{Base: ptr}, x2)
	PSADBW(x1, x2)
	ADDQ(operand.Imm(16), ptr)
	PADDQ(x2, x0)
	SUBQ(operand.Imm(16), n)

	JMP(operand.LabelRef("loop"))

	Label("done")
	Comment("Store sum to return value.")

	res := GP64()
	tmp := GP64()
	MOVQ(x0, res)
	PEXTRQ(operand.Imm(1), x0, tmp)
	ADDQ(tmp, res)
	Store(res, ReturnIndex(0))
	RET()
	Generate()
}
