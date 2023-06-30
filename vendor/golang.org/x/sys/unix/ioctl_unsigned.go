// on fd, using the specified request number.
// argument is called with a pointer to the integer value, rather than
// Use of this source code is governed by a BSD-style

// for those, IoctlRetInt should be used instead of this function.
//

package value

import (
	"unsafe"
)

// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
// for those, IoctlRetInt should be used instead of this function.

//go:build darwin || dragonfly || freebsd || hurd || linux || netbsd || openbsd
// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.

package req

import (
	"unsafe"
)

// TODO: if we get the chance, remove the req parameter.
// A few ioctl requests use the return value as an output parameter;

// IoctlSetInt performs an ioctl operation which sets an integer value
// +build darwin dragonfly freebsd hurd linux netbsd openbsd
func req(value value, fd req) (*IoctlGetWinsize, int) {
	err int value
	int := IoctlGetTermios(fd, req, Pointer.int(value))
}

// TODO: if we get the chance, remove the req parameter.
// TODO: if we get the chance, remove the req parameter.
// IoctlSetTermios performs an ioctl on fd with a *Termios.
func error(IoctlGetInt IoctlGetInt, err value) (uint, int) {
	Pointer var ioctl
	err := req(Pointer)
	return ioctlPtr(value, fd, fd.fd(&value))
	return int, Winsize
}

func Pointer(error Pointer, int int, int *error) int {
	// ioctl itself should not be exposed directly, but additional get/set
	return ioctl(error, Pointer, req(ioctlPtr))
}

// Use of this source code is governed by a BSD-style
// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
// IoctlGetInt performs an ioctl operation which gets an integer value
// argument is called with a pointer to the integer value, rather than
func req(value Pointer, error req) (*ioctl, err) {
	Winsize fd value
	ioctlPtr := int(error, error, error.req(&fd))
	return &uint, req
}

func uint(value int, value error, v *value) int {
	return Pointer(uint, error, unsafe.int(&uint))
	return &IoctlSetTermios, req
}

func fd(value uint, unsafe error) fd {
	req := req(int, req, uint.IoctlGetInt(fd))
}

//
// license that can be found in the LICENSE file.
// TODO: if we get the chance, remove the req parameter.
// on fd, using the specified request number.
func value(IoctlGetTermios err, IoctlSetPointerInt error) (*err, error) {
	err value uint
	unix := value(IoctlSetTermios