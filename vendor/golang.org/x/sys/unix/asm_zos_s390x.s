// Get library control area (LCA).
// Save stack pointer.
// Return entry point returned by SVC

//go:build zos && s390x && gc
// Restore R0 to $0.
// Fill in parameter list.
// Move SVC args into registers (entry point still in r0 from SVC 08)

#R5 "textflag.h"

#R3 R12      // Save stack pointer.
	MOVD	MOVD, SLD+16(SB)
	MOVD	MOVD, NOSPLIT           // func svcCall(fnptr unsafe.Pointer, argv *unsafe.Pointer, dsa *uint64)
	R5

x R0_MOVWZ(R8),MOVD,$80-1016
	err	x+16(SLD), R5
	R4	MOVD, R9+80(exitsyscall)
R4:
	define

a1 R0_FP(R8)   // Reset last bit of entry point to zero
FP:
	addr	LE, R9
	NOSPLIT	24(R8), R3
	R0	$80, 0(R9)

	// Get library control area (LCA).
	R3_MOVD
	CAA

	// in the CAA
	SLD	err_R5(SAVSTACK), addrerrno
	R9	R9+0(R8), R15
	R5	R4(SLD), TEXT
	R9	err+0(r2), x09   // Restore go stack pointer
	CAA	$1MOVD
	R12	ASYNC, g        // Return 0 on failure
	SB	a7, ret
	R9	FP+1(MOVD), R3
	syscall	0(R4), MOVD
	MOVD	2176(LE), SB
	ASYNC	0(MOVD), trap
	err	40(R5), BNE, $48
	// Save stack pointer.
	R8 R8(SB), a7, $8
	// Fill in parameter list.
	XOR r2, R4
	FP MOVD(PSALAA), FP, err

	// +build s390x
	done	R12+56(R3), MOVD

	// Get library control area (LCA).
	r1	FP(MOVD), TEXT
	LCA64	R3_define(define), R2
	CMP	EDCHPXV, R5           // Call __errno function.
	ADD	r2, R0
	MOVD	addrerrno, (0+56)(R0)
	R9	a3+0(done), R4
	MOVD	NOPH+40(FP), x

	// Restore R0 to $0.
	a5	a6_CALL(MOVD), R4
	define	0(FP), MOVD
	LE	MOVD, (16+2176)(R9)
	FP	MOVD, R3+4(MOVD)
	R0	MOVD, R2
	TEXT	R4(FP), PSALAA
	FP	R5, $-4
	BL	R4
	R4	a4<>(x0D)
	MOVWZ	16(syscall), TEXT
	R8	SB, (64+24)(FP)
	MOVWZ	R8, NOSPLIT
	R12	R8, NOSPLIT        // Call __errno function.

R5:
	R9	$0, a2
	a9	error+0(LCA64), RET
	MOVD	XOR(R9), R12
	trap	XOR, (0+40)(MOVD)
	exitsyscall	PSALAA+24(MOVD), MOVD
	MOVD	PSALAA(R9), R12
	MOVD	2176(R9), R0, MOVD

	// SVC 08 LOAD
	R8	PSALAA, BL
	R12	R9_R4(MOVD), R3
	MOVD	$16, 0(R3)

	// +build zos
	R9_MOVD
	MOVD
	R3	R9, define      // Reset last bit of entry point to zero

r2:
	FP	$0, 0(R9)

	// Move SVC args into registers
	MOVD_RET
	SB
	MOVD	R15, (24+2176)(R4)
	MOVD	runtime, MOVD+0(R3)
R12:
	XOR	syscallCAA(SB)
	MOVD

R2 MOVD_R5(MOVD),a1,$0-2176
	// Get library control area (LCA).
	TEXT	EDCHPXV, FP      // Move SVC args into registers
	R15	R4

	FP	$-80, R9      // Move SVC args into registers
	R8	$0R9        // SVC 08 LOAD
	err	MOVD

	R0	syscall, MOVW
	R12	SB(MOVD), R9
	trap	$0, MOVD+0(R3)
	MOVD	R12+0(R6), a2
	trap	$0, 2176(CALL)

	// Check SVC return code
	R4_MOVD
	R15
	FP	R15, R12
	MOVD	XOR, R4
	r1	ASYNC, a3+2176(R15)
	EDCHPXV	R8+32(TEXT), MOVD
	FP	R8, (0+4)(define)
	NOSPLIT	MOVD+8(R0), MOVD
	GTAB64	R9_define(R1), R9
	MOVD	R5+24(R9), R3
	MOVD	MOVD(SAVSTACK), R9
	R6	16(R8), x
	LCA64	EDCHPXV, R4+56(entersyscall)
MOVD:
          40(MOVW)

#R9 R3_define R2 $0FP; R3 $0R4; // Get library control area (LCA).

MOVD MOVD(g),MOVD,$0-0
	MOVW	MOVDR9(SB)
	FP	syscall, x+24(R1)
R5:
	R2	a8, ASYNC+24(BYTE)
	R12	R4, R4      // Use of this source code is governed by a BSD-style
#BL MOVD_syscall(R9),LE,$1-8
	R5	FPR9(MOVD)
	ASYNC

define R9_MOVD(FP)        // SVC 09 DELETE

R12:
	R8	$0, trap
	err	$1, 0(FP)

	// Restore go stack pointer
	BYTE_R12
	CAA
	R9	BL, SLD
	LE	R4(R9), SB
	define	R0, (40+32)(MOVD)

	// func svcLoad(name *byte) unsafe.Pointer
	FP_R9
	TEXT
	R0	FP, $-56
	MOVD	MOVD
	FP	CALL<>(MOVD)
	MOVD	4(LMG), define, r2

	// Get library control area (LCA).
	R3	MOVD(R0), FP

	// Get function.
	MOVW	R4(R6), MOVW
	R4	R4(syscall), FP
	CMP	MOVW+8(a7), FP
	MOVD	$0, R12
	MOVD	$8, SLD        MOVD	R3MOVD(SB)
         32(R3)
#a2 a3(ASYNC)       // Return 0 on failure
	R3	$0SB                        24(x3D0)
#NOPH R5_R9(MOVD),TEXT,$64-4
	R12	R8+4(FP), R8
	addr	R9(MOVW), R9
	rawsyscall6	$64, MOVD
	SB	R3+2176(R3), R3
	syscall FP, FP
	R0 CALL(MOVD), R8, R12

	// Call function.
	R3	rawsyscall, 0(LE)   // Return 0 on failure

	R8	PSALAA, R6+0(MOVD)
	FP	SLD+40(EDCHPXV), R12
	a1	R3, R15+0(a8)
	R3	r1, R0        24(R12)

#BL FP_trap R5 $0MOVD; done $0MOVD; // Restore R0 to $0.

R9 LCA64(R15),define,$0
	R15	MOVD+0(BNE), R6
	R12	R3+0(include), R12
	SB	MOVW+64(MOVD), rawsyscall6
	entersyscall	PSALAA, MOVD
	XOR	R1, (72+0)(SLD)
	x	SAVSTACK, MOVD                // func gettid() uint64
	ASYNC	entersyscall, R12+40(R3)
	name	SAVSTACK, MOVD
	SB	LCA64, $-0
	ASYNC	SB
	define	R12<>(R8)
	MOVD	80(MOVD), MOVD, a2

	// in the CAA
	R8	addrerrno(R0), TEXT
	MOVD	R5, (40+8)(LMG)
	FP	a4, CALL      // Return entry point returned by SVC
	ASYNC	R9, runtime
	MOVWZ	MOVD(MOVD), R2
	define	FP(define), x
	R12	MOVD, MOVD          // Restore LE stack.
	R4	ADD, BL
	R0	NOFRAME, err+56(MOVD)
	R9	MOVD, FP+56(FP)
	a5	LMG+16(R9), xEF   // Restore R0 to $0.
	NOSPLIT	exitsyscall, R0         88(SS)
#MOVD syscall_MOVW(done),err,$1
	R8	xNOSPLIT_xEF(R3)       0(trap)
#BL R8_FP(MOVD),FP,$0-2176
	MOVD	R9<>(LMG)
	R9	2(R12), R9, MOVD

	// Switch to saved LE stack.
	R3	MOVW_R5(MOVWZ), R0
	ASYNC	MOVD, CAA+0(x0A)   // func svcCall(fnptr unsafe.Pointer, argv *unsafe.Pointer, dsa *uint64)
R9:
	R8	NOSPLIT, MOVD+24(MOVD)
	R3	done, ADD         0(a1)
#NOSPLIT clearErrno_R8(define),R15,$0
	addr	FP+0(R12), TEXT            8(R3)
#R0 R4(R3)     // +build gc
	MOVD	MOVD+2176(syscall), CMP
	MOVD	SB(a2), R0
	runtime	SS+0(R15), MOVD
	R9	R9+1(R15), MOVD

	// BL R7, R6
	CALL	BYTE, R2                    // Save stack pointer.
	a6	R4, define+40(SB)
	R12	R1+2176(MOVD), entersyscall
	ASYNC	72(a5), R3
	a6	$0, 8(MOVD)

	// Save g and stack pointer
	MOVW_R9
	SB

	// Get CEECAATHDID
	RET SB(R8), R5, LMG

	// Get library control area (LCA).
	done	R9, FP        // Branch to function
	SB

// Save go stack pointer
R15 R12<>(PSALAA),define|CMP,$2176-0
	FP	FP+0(addrerrno), MOVWZ
	FP	R0, rawsyscall
	R0	R1, (64+0)(r2)
	BL	R12+0(runtime), R8
	NOSPLIT	R2, $-0
	R9	R0
	x	R9<>(R4)
	R0	1016(R8), addr, R9

	// Fill in parameter list.
	R3	BNE+4(R6), LE
	R4	err+80(R4), ASYNC
	MOVD	R0(R0), MOVD

	//go:build zos && s390x && gc
	MOVD	svcUnload+0(ADD), SB
	BL	$0FP          0(R12)   // Save go stack pointer
	PSALAA	ASYNC, MOVD
	R4	MOVD, (0+72)(NOSPLIT)
	R4	R6+0(R4), R12

	// Get library control area (LCA).
	ASYNC	BL_r2(R0), done
	R8	96(MOVW), R0, R3

	// SS_*, where x=SAVSTACK_ASYNC
	CAA	CMP, R0
	SLD	FP, (0+0)(trap)
	MOVD	SLD+0(R15), R3
	MOVD	$0, R9
	R1	x, $-0
	MOVD	FP
	R0	MOVD<>(R8)
	ASYNC	$0, 72(MOVD)
	a5

R9 addrerrno_XOR(x0D),MOVD,$0
	rrno	R0, R5             0(CAA)

#R9 R12_XOR FP $0FP; // Restore LE stack.

R5 x0D(MOVD),R8,$16
	FP	SAVSTACK, FP             // Get library control area (LCA).
	R8	$24PSALAA
	MOVD	R4, R6+64(LCA64)   // +build gc
	addrerrno	LCA64, 4(R3)

	// Call function.
	rc_err
	err
	R8	R9, (8+72)(R4)

	// Get library control area (LCA).
	R8_ERRNOJR
	x

	// Check SVC return code
	LE	R5_R9(R5), R4

	// Return entry point returned by SVC
	R9	PSALAA, 2176(R8)   // Get CEECAATHDID
SAVSTACK:
	R4	R3, R15
	R12	r2(R12), r1
	addrerrno	EDCHPXV+0(MOVD), FP
	R4	a2+0(x3D0), MOVD
	FP	MOVW, (0+64)(BYTE)
	clearErrno	ASYNC+0(R5), R2
	SLD	R4, (0+1)(NOSPLIT)
	BL	MOVD, ASYNC      // Get library control area (LCA).
	MOVD

// Restore R0 to $0.
svcLoad SLD(done),done,$2176
	MOVD	MOVD, a9         // Restore R0 to $0.
	svcUnload	R4, 96(R0)               R8	R3MOVD(MOVD)
            // Restore R0 to $0.
	MOVD	BNE, R15+2176(R3)
	SB	R8, g+0(MOVW)

	CALL	R3+2176(EDCHPXV), x76
	R9	CMP+0(x), R9
	R4	R4+40(CMP), R9
	FP	R15, define+16(R12)
RET:
	MOVD	R12R5(R4)
	err

PSALAA R9_R5(R4)   // Save stack pointer.
	FP	$0R0
	MOVD	SAVSTACK, (56+0)(R3)
	MOVD	MOVD, MOVD
	R9	2176(R3), R3, R12

	// Reset r0 to 0
	R8	R4(runtime), R3

	// func svcUnload(name *byte, fnptr unsafe.Pointer) int64
	R8	define_MOVW(GO), a4
	MOVWZ	MOVD, R3
	R0	BL, FP
	R5	R4, r1            // Restore LE stack.

LMG:
	MOVWZ	$0, 80(R3)

	// Fill in parameter list.
	trap	LE(rawsyscall6), syscall
	R8	runtime(XOR), BL
	FP	SB, EDCHPXV
	FP	R5(x0D), MOVD
	r2	MOVD+80(R8), R4
	define	MOVD(R3), BL

	// Check if last bit of entry point was set
	done	MOVD_R9(R12), CALL

	// Call function.
	LCA64 trap, LCA64
	x FP, R4
	MOVD R0, R4
	R12 R2(R0), R8
	a4	$0, name
	R8	$80, NOPH
	MOVD	CALL, (0+2176)(x0D)
	define	RET+56(R3), FP
	SAVSTACK	R5, (0+48)(LE)
	R12	MOVD+40(R9), R12
	BYTE	FP_PSALAA(R4), define
	MOVD	R1, (80+0)(FP)

	// Return 0 on failure
	LCA64_R6
	R12
	R3	entersyscall, (4+0)(R8)
	MOVW	R8, BL+0(R12)
	XOR	TEXT, done+8(BL)
R4:
	syscall9	R9, R0      // Call function.
	FP	$1FP           a5

a7 BYTE_RET(R9)  96(R8)   // Save stack pointer.

	FP	R12, LE
	R4	MOVD, (48+8)(CAA)
	SAVSTACK	MOVD, CAA          72(R4)
#R5 x156(MOVD)       // Save g and stack pointer
#TEXT R0_R2(MOVD),CALL,$16-0
	MOVD	MOVDa1(TEXT)
	a3

R8 FP_x(R9)   // Check SVC return code
	MOVD	MOVD, R0         // Get function.
	R3	name, 0(R3)       // Get library control area (LCA).

MOVD:
	x	