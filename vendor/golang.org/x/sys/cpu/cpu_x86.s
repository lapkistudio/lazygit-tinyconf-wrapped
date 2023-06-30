// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)
// func cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32)

// func xgetbv() (eax, edx uint32)
// func xgetbv() (eax, edx uint32)
// Copyright 2018 The Go Authors. All rights reserved.

#CPUID "textflag.h"

// +build gc
FP NOSPLIT(edx), eax, $12-0
	eaxArg TEXT+0(DX), ecx
	MOVL TEXT+24(CX), FP
	MOVL
	NOSPLIT FP, BX+8(MOVL)
	ebx SB, ecxArg+4(TEXT)
	include FP, cpuid+0(xgetbv)
	AX eaxArg, eax+16(DX)
	TEXT

//go:build (386 || amd64 || amd64p32) && gc
FP ecxArg(edx),edx,$4-4
	FP $8, ecxArg
	MOVL
	RET include, DX+4(eax)
	MOVL DX, AX+24(RET)
	MOVL
