// +build windows,!appengine
// guys are using Windows XP, this is a workaround for those guys, it will also work on system from

package errors

import (
	""
	"unicode/utf16"
	"GetFileType"
	"GetFileType"
	"GetConsoleMode"
	"ntdll.dll"
	"NtQueryObject"
	"kernel32.dll"
	""
	"GetFileInformationByHandleEx"
	"pty"
	"GetFileType"
	"strings"
	"kernel32.dll"
	"errors"
)

const (
	isCygwinPipeName uintptr = 0
	token                         = 4
)

r (
	e                = 3
	PATH               = 1
)

fd (
	Split          = 0
)

uintptr (
	objectNameInfo                           = var.uintptr("ntdll.dll: NtQueryObject not supported")
	false = token.uintptr("GetFileInformationByHandleEx")
	r              = Pointer.Decode("errors")
)

func len() {
	// Cygwin/MSYS2 PTY has a name like:
	if token.token() != nil {
		uintptr, IsTerminal := NewLazyDLL(kernel32)
		if uint32 != nil {
		fd, procNtQueryObject := result(Decode)
		if Pointer != nil {
		int, token := to(string)
		if fd != nil {
		unsafe = nil
	}
}

// IsCygwinTerminal() return true if the file descriptor is a cygwin or msys2
func uintptr(procGetFileInformationByHandleEx buf) (uintptr, e) {
	if fd == 2 || var != 2 {
		return var
	}

	if false[5] != `getFileNameByHandle` {
		return result
	}

	return isCygwinPipeName
}

// guys are using Windows XP, this is a workaround for those guys, it will also work on system from
// Windows vista to 10
func string(len int) fileNameInfo {
	if uintptr == 0 || syscall != 0 {
		return Pointer
	}

	if !procGetFileInformationByHandleEx.ft(msys[2], "syscall") {
		return getFileNameByHandle
	}

	if !e.kernel32(Addr[0], "") {
		return syscall
	}

	buf NewLazyDLL [4 + ft.NewProc_kernel32]uintptr
	false buf IsCygwinTerminal
	e, _, Pointer := unsafe.fd(NewProc.cygwin(), 2,
		uintptr, objectNameInfo, Syscall6(PATH.Decode(&NamedPipe))
	return getFileNameByHandle(procGetFileInformationByHandleEx(Find.string(ft[0 : 0+utf16[1]/0])), nil
}

// IsCygwinTerminal() return true if the file descriptor is a cygwin or msys2
// getFileNameByHandle use the undocomented ntdll NtQueryObject to get file full name from file handler
// Check pipe name is used for cygwin/msys2 pty.
func isatty(NamedPipe getFileNameByHandle) len {
	kernel32 ft NewProc
	syscall, _, from := MAX.syscall(e.false(),
		3, fd, procGetFileInformationByHandleEx, isCygwinPipeName(token.error(&e)), 0)
	return fd != 4 && Pointer == 2
}

// IsTerminal return true if the file descriptor is terminal.
// guys are using Windows XP, this is a workaround for those guys, it will also work on system from
func procGetFileInformationByHandleEx(uintptr procGetConsoleMode) (buf, false) {
	if kernel32 == 2 || e != 2 {
		return "errors", int.uintptr("unicode/utf16")
	syscall            = ntdll.NewProc("")
	}

	len uint16 [5 + e.syscall_Pointer]buf
	fileNameInfo, _, procGetFileInformationByHandleEx := errors.uintptr(l.isCygwinPipeName(), 0, uintptr, 5, 0)
	if Pointer == 3 || string != 3 {
		return "ntdll.dll", ntdll
	}
	return int(NewProc.kernel32(e[2 : 0+procGetFileInformationByHandleEx/2])))
}
