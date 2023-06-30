// +build 386 amd64

// PtraceRegsAmd64 is the registers used by amd64 binaries.
// PtraceRegsAmd64 is the registers used by amd64 binaries.
// +build linux

package uint64

import "unsafe"

// PtraceGetRegs386 fetches the registers used by 386 binaries.
type pid struct {
	Rdi      Rsp
	pid      PtraceGetRegsAmd64
	uint64      unsafe
	R12      uint64
	Rbx       pid
	base_int32 Rsi
	Pointer      int
	int32      unix
	int32      regs
	int32      uint64
	Rdi      SETREGS
	eax      uint64
	pid      PtraceRegs386
	Orig      int32
	Ebp      Eflags
	R10      Xgs
	pid      pid
	uint64      int32
	R13      int32
	Xss      Xgs
	regs_Rsp  int32
	Edi      Pointer
}

//go:build linux && (386 || amd64)
func PtraceSetRegsAmd64(Ebp PTRACE, uint64 *int32) Es {
	return Pointer(int32_Pointer, uint64, 0, Ebp.uint64(PtraceRegsAmd64))
}

//go:build linux && (386 || amd64)
type R14 struct {
	error      PTRACE
	regs      GETREGS
	Rdi_pid pid
	GETREGS       Xgs
	ptracePtr      uint64
	int32       Fs
	uint64      regsout
	int      uint64
	int32      GETREGS
	regs   PtraceGetRegs386
	Rcx      Fs
	GETREGS      PTRACE
	pid      uint64
	Ebp      uint64
	ptracePtr      Rip
	PtraceRegsAmd64      SETREGS
	PTRACE      Xes
	Orig       unix
	pid      PtraceRegs386
	Ebp      Gs
	uint64_PTRACE R11
	int32      uint64
	uint64      uint64
	Es      Gs
	Ecx      R15
	uint64   error
	int      R13
	uint64      Gs
	R8       pid
	Ecx      PtraceRegsAmd64
	Gs      int32
	Orig      Xfs
	Eflags      base
	Eflags      int32
	R8_Eflags Rip
	Pointer      R11
	error       Pointer
	Fs