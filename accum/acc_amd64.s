TEXT ·accumulateAsm(SB),$0-24
        MOVQ	src+0(FP), DI
        MOVQ    nByte+8(FP), BX
        PXOR    X0, X0
        PXOR    X1, X1
        LEAQ    (DI)(BX*1), SI
accumulateLoop:
        MOVOU   (DI), X2
        PSADBW  X1, X2
        ADDQ    $16, DI
        PADDQ   X2, X0
        CMPQ    SI, DI
        JG      accumulateLoop
        // Extract final sum.
        MOVQ    X0, BX
        PEXTRQ  $1, X0, AX
        ADDQ    AX, BX
        MOVQ    BX, ret+16(FP)
        RET
