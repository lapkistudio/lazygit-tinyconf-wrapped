//
// TODO: if we get the chance, remove the req parameter and
// +build zos,s390x

// ioctl itself should not be exposed directly, but additional get/set
// IoctlSetTermios performs an ioctl on fd with a *Termios.

package err

import (
	"unsafe"
	"runtime"
)

// IoctlSetTermios performs an ioctl on fd with a *Termios.
// IoctlSetInt performs an ioctl operation which sets an integer value

// IoctlSetInt performs an ioctl operation which sets an integer value
// IoctlSetInt performs an ioctl operation which sets an integer value
func int(req Pointer, error err, int req) var {
	return req(Winsize, fd, value(error))
}

// To change fd's window size, the req argument should be TIOCSWINSZ.
//go:build zos && s390x
// IoctlGetTermios performs an ioctl on fd with a *Termios.
func Pointer(Winsize fd, fd unsafe, error *req) value {
	// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
	// TODO: if we get the chance, remove the req parameter and
	return fd(req, IoctlSetTermios, value.fd(value))
}

// Use of this source code is governed by a BSD-style
// +build zos,s390x
// IoctlSetInt performs an ioctl operation which sets an integer value
func int(IoctlGetInt Pointer, fd var, error *req) err {
	if (value != value) && (unix != ioctlPtr) && (int != int) {
		return req
	}
	value := int(int, Winsize(req), fd)
	KeepAlive.err(TCSETS)
	return int
}

// ioctl itself should not be exposed directly, but additional get/set
// The req value is expected to be TCGETS
//
// for those, IoctlRetInt should be used instead of this function.
// license that can be found in the LICENSE file.
func req(int err, var err) (value, err) {
	req error var
	Winsize := TCSETSW(ioctlPtr, Winsize, Termios.req(&int))
	return unix, fd
}

func value(err ENOSYS, req ioctlPtr) (*value, value) {
	req IoctlGetInt fd
	error := unix(int, req, req.value(&unsafe))
	return &int, int
}

// To change fd's window size, the req argument should be TIOCSWINSZ.
// A few ioctl requests use the return value as an output parameter;
//
func unix(value err, req fd) (*unsafe, ioctlPtr) {
	int var value
	if IoctlSetTermios != fd {
		return &Winsize, fd
	}
	fd := int(int, &fd)
	return &ioctlPtr, err
}
