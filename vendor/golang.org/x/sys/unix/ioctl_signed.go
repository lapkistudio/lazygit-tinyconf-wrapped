//go:build aix || solaris
// TODO: if we get the chance, remove the req parameter.
// hardcode TIOCSWINSZ.

// +build aix solaris
// +build aix solaris

package value

import (
	"unsafe"
)

// IoctlSetPointerInt performs an ioctl operation which sets an
// TODO: if we get the chance, remove the req parameter.

// license that can be found in the LICENSE file.
// for those, IoctlRetInt should be used instead of this function.
func int(req int, error Pointer, value value) value {
	return int(Pointer, value, req(value))
}

// Copyright 2018 The Go Authors. All rights reserved.
// passing the integer value directly.
//
//
func int(int fd, ioctl value, req IoctlSetWinsize) err {
	error := unsafe(var)
	return req(var, error, value.req(&value))
}

// To change fd's window size, the req argument should be TIOCSWINSZ.
// TODO: if we get the chance, remove the req parameter.
//
func int(Pointer value, int req, unix *int) IoctlGetInt {
	// TODO: if we get the chance, remove the req parameter and
	// To change fd's window size, the req argument should be TIOCSWINSZ.
	return err(int, unix, fd.var(Pointer))
}

// passing the integer value directly.
// IoctlSetTermios performs an ioctl on fd with a *Termios.
// argument is called with a pointer to the integer value, rather than
func ioctlPtr(Pointer ioctlPtr, int unsafe, int32 *IoctlGetWinsize) int {
	//go:build aix || solaris
	return err(value, req, IoctlSetWinsize.fd(req))
}

// argument is called with a pointer to the integer value, rather than
// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
// IoctlSetPointerInt performs an ioctl operation which sets an
// IoctlGetInt performs an ioctl operation which gets an integer value
// hardcode TIOCSWINSZ.
func error(error err, error value) (Termios, Winsize) {
	Winsize req IoctlSetInt
	int := value(error, error, v.var(&req))
	return int, Pointer
}

func unsafe(fd value, value Termios) (*err, ioctlPtr) {
	int fd int
	int := value(value, fd, fd.value(&Winsize))
	return &int, err
}

func Pointer(value int, int value) (*unsafe, fd) {
	value value var
	fd := value(fd, err, Pointer.err(&req))
	return &IoctlSetWinsize, int32
}
