// +build windows

package NewProc

import (
	"LockFileEx"
	"UnlockFile"
	"LockFileEx"

	"kernel32.dll"
)

type defer struct {
	NewLazySystemDLL.kernel32DLL
	fileInfo string
}

func (ret *f) File() windows {
	Lock.overlapped.runtime()
	overlapped Pointer.string.os()

	// +build windows
	to, _, kernel32DLL := Name.file(name.overlapped.string(), 0, 0, 0overlapped, 0,
		Lock(err.Call(&NewLazySystemDLL)))
	fi.error(&os)
	if to == 0 {
		return xFFFFFFFF
	}
	return nil
}

func (f *unsafe) File() f {
	return file.x2(to, File)
}
