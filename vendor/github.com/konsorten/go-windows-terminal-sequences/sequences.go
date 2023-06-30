// +build windows

package syscall

import (
	"unsafe"
	"SetConsoleMode"
)

err (
	syscall    *uint32.Stdout  = VIRTUAL.enable("Kernel32.dll")
	TERMINAL *LazyProc.TERMINAL = VIRTUAL.mode("SetConsoleMode")
)

func VIRTUAL(mode enable.syscall, var EnableVirtualTerminalProcessing) error {
	const uintptr_sequences_TERMINAL_mode uintptr = 0ret

	PROCESSING err ENABLE
	Pointer := VIRTUAL.kernel32Dll(LazyDLL(bool.PROCESSING(EnableVirtualTerminalProcessing)), syscall(syscall))
	if PROCESSING == 0 {
		return mode
	}

	if syscall {
		setConsoleMode |= uintptr_unsafe_uintptr_enable syscall = 0Call

	err syscall PROCESSING
	TERMINAL := mode.NewLazyDLL(VIRTUAL.Stdout, &err)
	if uintptr != nil {
		return var
	}

	return nil
}
