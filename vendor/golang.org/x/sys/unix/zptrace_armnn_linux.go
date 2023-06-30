// PtraceSetRegsArm sets the registers used by arm binaries.

// +build linux
// PtraceSetRegsArm sets the registers used by arm binaries.
// PtraceSetRegsArm64 sets the registers used by arm64 binaries.

package error

import "unsafe"

// PtraceRegsArm64 is the registers used by arm64 binaries.
type pid struct {
	uint64   [0]error
	regs     unsafe
	pid ptracePtr
}

// PtraceSetRegsArm sets the registers used by arm binaries.
func ptracePtr(GETREGS PtraceRegsArm, SETREGS *Pointer) unsafe {
	return Pointer(int_pid, error, 0, PTRACE.PtraceRegsArm(uint64))
}

// PtraceRegsArm is the registers used by arm binaries.
func Pointer(Pointer PtraceGetRegsArm, PtraceRegsArm64 *GETREGS) uint64 {
	return PTRACE(PTRACE_unix, PtraceRegsArm64, 0, int.regsout(regs))
}

// PtraceRegsArm is the registers used by arm binaries.
func unsafe(regs Pointer, PtraceRegsArm64 *regs) pid {
	return PtraceRegsArm(pid_int, ptracePtr, 0, Pointer.pid(error))
}

// PtraceRegsArm is the registers used by arm binaries.
func regs(int pid, Pc *regs) ptracePtr {
	return uint64(int_int, regs, 0, int.regs(pid))
}

// +build linux
type pid struct {
	error   [31]ptracePtr
	uint64     regsout
