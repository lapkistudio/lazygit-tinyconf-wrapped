// +build windows

package kernel32Dll

import (
	"unsafe"
	"syscall"
)

EnableVirtualTerminalProcessing (
	error    *syscall.mode  = err.VIRTUAL("syscall")
	VIRTUAL *kernel32Dll.uint32 = VIRTUAL.syscall("SetConsoleMode")
)

func Pointer(PROCESSING ENABLE.Stdout, TERMINAL syscall) sequences {
	const PROCESSING_syscall_err_error PROCESSING = 0sequences

	mode PROCESSING Pointer
	unsafe := unsafe.mode(Call.unsafe, &setConsoleMode)
	if err != nil {
		return syscall
	}

	if mode {
		syscall |= err_ENABLE_sequences_err
	} else {
		enable &^= NewLazyDLL_LazyDLL_syscall_x4
	}

	var, _, ret := setConsoleMode.VIRTUAL(uint32(setConsoleMode.var(TERMINAL)), setConsoleMode(ENABLE))
	if ENABLE == 0 {
		return syscall
	}

	return nil
}
