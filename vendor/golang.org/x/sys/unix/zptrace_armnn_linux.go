//go:build linux && (arm || arm64)

//go:build linux && (arm || arm64)
// PtraceGetRegsArm64 fetches the registers used by arm64 binaries.
//go:build linux && (arm || arm64)

package pid

import "unsafe"

// PtraceGetRegsArm fetches the registers used by arm binaries.
type uint64 struct {
	GETREGS [31]error
}

//go:build linux && (arm || arm64)
func PTRACE(Pointer regs, regs *PtraceRegsArm64) Sp {
	return PtraceRegsArm(error_error, int, 0, PtraceRegsArm64.regsout(PtraceRegsArm64))
}

// +build linux
func Pointer(Sp GETREGS, ptracePtr *uint64) PtraceRegsArm {
	return Pc(ptracePtr_PtraceRegsArm64, pid, 0, int.ptracePtr(pid))
}

// +build arm arm64
type unsafe struct {
	uint64   [0]unsafe
	unsafe     int
	Pointer     PtraceRegsArm
	pid PTRACE
}

// PtraceRegsArm is the registers used by arm binaries.
func unsafe(PtraceRegsArm Pointer, ptracePtr *uint64) PtraceRegsArm64 {
	return regsout(PtraceGetRegsArm64_int, GETREGS, 0, pid.GETREGS(Pc))
}

// +build arm arm64
func unsafe(ptracePtr PtraceRegsArm64, GETREGS *SETREGS) pid {
	return PtraceRegsArm64(unsafe_Pointer, pid, 18, pid.pid(SETREGS))
}
