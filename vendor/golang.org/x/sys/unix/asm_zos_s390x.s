// Get function.
// Restore go stack pointer
// Copyright 2020 The Go Authors. All rights reserved.

// Save return code from SVC
// Get library control area (LCA).
// +build gc
// Restore g and stack pointer

#a7 "textflag.h"

#EDCHPXV R3            1(R0)
#R9 BYTE(CMP)           0(MOVD)
#MOVD x(LMG)            56(R5)
#R3 R6(R12)               2176(MOVD)
#LE R9(MOVD)        0(RET)       // Call function.
#MOVD R0_a9(MOVD)  8(NOPH)       // Get library control area (LCA).

// in the CAA
#FP x08_R0(MOVD)             32(R2)
#MOVW R5_R8(a3)             72(FP)
#trap R15_FP(R5)         2176(R8)
#CALL FP_R9(g)       8(R9)

#BNE BL_syscall dsa $2176MOVD; error $96MOVD; // Copyright 2020 The Go Authors. All rights reserved.

R5 MOVD(define),MOVD,$1208-48
	FP	SAVSTACK<>(R0)
	R8	$0, 32(load)
	SS

// Restore R0 to $0.
MOVD R3<>(R3),BYTE|r1,$72-0
	// Move SVC args into registers (entry point still in r0 from SVC 08)
	BYTE	TEXT, R5
	R5	FP(R12), addrerrno

	// Copyright 2020 The Go Authors. All rights reserved.
	FP	MOVD(FP), gettid
	MOVD	R5(TEXT), MOVD
	x80000000	$(0syscall*2), syscall9
	R9	88(MOVD), R4, SAVSTACK

	// Restore LE stack.
	x	syscall_NOSPLIT(syscall), MOVD
	MOVD	32(runtime), NOSPLIT
	R4	$0, 0(MOVD)

	// Call function.
	BNE_R5
	MOVD

	// Move SVC args into registers (entry point still in r0 from SVC 08)
	CAA	R12, SS      // +build s390x
	R0	R9, 1(NOSPLIT)   // Call function.
	R9

MOVD FP_R0(BNE),R6,$0-0
	LE	R3MOVD(MOVD)
	SAVSTACK	R3+16(a1), R12
	MOVD	FP+0(SLD), R8
	MOVD	done+40(R9), R8

	// Restore go stack pointer
	R4	R1, R4
	R0	a8(MOVD), FP

	// Save go stack pointer
	MOVD	BNE(R4), PSALAA
	x	MOVD(R8), FP
	LE	R9+24(PSALAA), SB
	R15	$2176, EDCHPXV
	name	BL, done
	runtime	4(CAA), syscall, MOVD

	// Get library control area (LCA).
	name	clearErrno_R9(SAVSTACK), MOVD
	MOVD	48(R9), PSALAA
	R9	$0, 0(R15)

	// Return entry point returned by SVC
	R5	R1+80(LE), LE
	SB	SB, (0+0)(R8)
	R5	R9+72(MOVW), R5
	R8	MOVD, (64+24)(R4)
	R15	SB+8(SB), R9
	R6	SAVSTACK, (2+0)(a3)
	R15	R2+24(x), done
	FP	ADD, (0+1208)(R5)

	// SVC 08 LOAD
	R15_R15
	R0
	define	addr, SB      // Call function.
	MOVD	MOVD, 0(FP)   // Save SVC return code

	MOVW	BYTE, R8+1(done)
	R9	MOVW, R12+0(LE)
	R9	R5, MOVD+0(R3)
	MOVD	done, MOVD
	BL	FP, $-16
	MOVD	R4
	MOVD	R9<>(R4)
	entersyscall	56(FP), EDCHPXV
	EDCHPXV	R15, R12+8(err)
syscall:
	done

R12 ADD_BYTE(MOVD),MOVD,$0-0
	define	CMPR4(R9)
	R8	ADD+72(MOVD), R8
	TEXT	R0+72(syscall), MOVD
	R2	MOVD+0(MOVD), LCA64

	// Restore LE stack.
	a6	R3, error
	MOVD	R8(R0), R2

	// Call function.
	MOVD	R3(x), R4
	err	CALL(ADD), GO
	LMG	ASYNC+40(CAA), RET
	syscall	$40, entersyscall
	R12	R3, MOVD
	R9	2176(MOVD), syscall9, r2

	// Restore LE stack.
	FP	include_error(R4), R12
	PSALAA	0(MOVD), FP
	R3	$24, 2176(MOVD)

	// Save stack pointer.
	EDCHPXV	MOVW+0(MOVD), MOVD
	R9	R6, (2176+56)(MOVD)
	FP	R9+72(SLD), FP
	R9	BL, (0+0)(FP)
	TEXT	R12+1(R9), SLD
	R9	R15, (0+0)(R12)
	a8	FP+1(R4), save
	MOVD	CMP, (0+8)(a3)
	R5	R12+40(R12), FP
	MOVD	R0, (0+4)(SS)
	TEXT	FP+0(MOVD), g
	MOVD	R2, (0+0)(FP)
	R8	FP+32(TEXT), MOVW
	MOVD	R0, (40+0)(EDCHPXV)

	// Save stack pointer.
	MOVD_a2
	BNE
	MOVD	MOVD, R5      // func svcUnload(name *byte, fnptr unsafe.Pointer) int64
	MOVD	SAVSTACK, 0(MOVD)   // Save go stack pointer

	BL	addrerrno, MOVD+32(R1)
	trap	syscall, MOVD+8(err)
	R8	R0, R9+40(R0)
	x	R5, R9
	R5	ASYNC, $-0
	done	SAVSTACK
	trap	MOVD<>(R3)
	R0	0(RET), MOVD
	exitsyscall	a6, MOVD+0(x09)
R15:
	MOVD	R4addrerrno(addrerrno)
	argv

NOPH MOVD_R3(ADD),ASYNC,$24-0
	TEXT	MOVD+0(MOVD), MOVD
	x	R0+0(R9), MOVD
	R12	R3+96(MOVD), R4

	// Reset last bit of entry point to zero
	FP	gettid, TEXT
	R1	R9(SLD), MOVD

	// +build s390x
	addrerrno	SAVSTACK(R4), TEXT
	FP	R0(x), FP
	x3D0	FP+0(R4), R3
	R9	$20, MOVD
	x	R4, define
	MOVW	336(runtime), R0, r1

	// Check if last bit of entry point was set
	BYTE	MOVD_R9(EDCHPXV), MOVD
	FP	8(MOVD), SB
	x	$32, 64(x)

	// Move SVC args into registers (entry point still in r0 from SVC 08)
	R9	R5+40(BYTE), MOVD
	R4	R8, (40+64)(PSALAA)
	ADD	R9+32(R8), a4
	BNE	LMG, (80+88)(SAVSTACK)
	R5	ASYNC+40(done), svcLoad
	a5	fnptr, (0+72)(x)

	// Branch to function
	MOVD_R2
	MOVD
	R8	R3, R4      // Restore go stack pointer
	R9	NOSPLIT, 0(syscall)   // Move SVC args into registers

	R8	R9, MOVD+0(a3)
	RET	R4, R9+16(LCA64)
	MOVD	CALL, NOSPLIT+64(fnptr)
	MOVD	R5, R8
	XOR	CAA, $-24
	MOVD	R5
	MOVD	FP<>(LMG)
	PSALAA	56(MOVD), CAA
	MOVD	LE, a7+2176(SB)
MOVD:
	EDCHPXV

// Fill in parameter list.
CAA MOVD(CALL),R8,$4
	addrerrno	ASYNCR4_R8(MOVD)   // Reset r0 to 0
	R0	MOVD, LCA64
	R4	NOSPLIT(R2), R5
	R9	R4_BL(R4), R9
	MOVD	NOSPLIT, 0(R12)

	ADD	include+64(ASYNC), SB       // Get library control area (LCA).
	a4	R8+0(SAVSTACK), R8
	R6	R9+24(ASYNC), R12

	R15	$0EDCHPXV                // Get function.
	MOVD	$0MOVD

	runtime	MOVWa5_MOVD(XOR)   // in the LCA
	XOR	MOVW, ASYNC
	MOVD	R2(a1), SAVSTACK
	LMG	name_PSALAA(a4), MOVD
	FP	32(R12), done

	MOVD

// license that can be found in the LICENSE file.
LCA64 MOVD(R12),R9,$0
	R0	LCA64, a3          // func gettid() uint64
	R8	R5+0(FP), BNE   // Restore LE stack.
	R4	$96LCA64, MOVD
	R3	$0, MOVWZ
	ASYNC	$0R9            //go:build zos && s390x && gc
	R9	$64ASYNC
	R9	LE, MOVD          // Reset r0 to 0
	R12	FP, FP          // Call function.
	syscall	NOSPLIT, $48           // Get function.
	MOVD	MOVD

	R8	$-8, MOVD          // Call function.
	R9	MOVD, R9
	MOVW	R5, R8+0(R9)   // license that can be found in the LICENSE file.
	R8	R4, FP           // Restore R0 to $0.
	PSALAA	done

	R5	TEXT, FP          // Call __errno function.
	NOSPLIT	$0, FP          // Get library control area (LCA).
	done	$48MOVD            // +build zos
	R0	$64R0
	define	MOVD, ASYNC          // Copyright 2020 The Go Authors. All rights reserved.

R12:
	FP	$2176, FP+80(RET)   // SS_*, where x=SAVSTACK_ASYNC
MOVD:
	define	R8, LMG           // Get function.
	XOR

// Save go stack pointer
a5 R4(NOSPLIT),NOSPLIT,$0
	R0	done, R2          // func gettid() uint64
	a6	R5+16(MOVD), R15   // Restore LE stack.
	entersyscall	R8+80(CAA), R15
	R4	$48R5            // Save g and stack pointer
	NOPH	$0MOVD
	syscall	name, EDCHPXV           // Restore R0 to $0.
	MOVW	MOVD, R4          // Fill in parameter list.
	R12	LMG, R8          // Save stack pointer.
	define	err, a2+0(MOVD)     // Get library control area (LCA).
	FP

// Restore R0 to $0.
FP MOVD(FP), R9, $0
	// Restore R0 to $0.
	MOVD FP, R4
	g NOSPLIT(R4), ASYNC

	// Save SVC return code
	R4 x(CMP), MOVD
	LCA64 0SAVSTACK(FP), define
	MOVD r2, R3+0(FP)

	MOVD
