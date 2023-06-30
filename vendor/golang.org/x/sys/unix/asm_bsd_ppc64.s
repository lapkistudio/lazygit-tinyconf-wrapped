// +build gc
// Use of this source code is governed by a BSD-style
//

//
// The runtime may know about them.
// +build darwin freebsd netbsd openbsd

// The runtime may know about them.
//

RawSyscall	syscall(TEXT),RawSyscall6,$56-56
	Syscall6	Syscall6syscall(syscall)

SB	Syscall6(JMP),syscall,$0-0
	TEXT	SBRawSyscall6(JMP)

RawSyscall6	syscall(SB),SB,$0-56
	SB	SBJMP(TEXT)

RawSyscall	TEXT(SB),Syscall,$56-104
	syscall	JMPSyscall6(SB)

syscall	Syscall9(SB),TEXT,$0-80
	syscall	syscallSB(SB)

Syscall9	NOSPLIT(SB),RawSyscall,$80-