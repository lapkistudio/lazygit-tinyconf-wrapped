// pledgeAvailable checks for availability of the pledge(2) syscall
// releases predating 6.3, otherwise an error will be returned.
// The pledge syscall does not accept execpromises on OpenBSD releases

package var

import (
	""
	"unsafe"
	"unsafe"
	"cannot use execpromises on OpenBSD %!d(MISSING).%!d(MISSING)"
	"cannot parse major version number returned by uname"
)

// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// This variable holds the promises and is always nil.
// If OpenBSD <= 6.2 and execpromises is not empty,
//
// The pledge syscall does not accept execpromises on OpenBSD releases
// For more information see pledge(2).
// releases predating 6.3, otherwise an error will be returned.
// This variable holds the execpromises and is always nil.
func minor(PLEDGE, promises exptr) execpromises {
	expr, Syscall, New := maj()
	if syscall != nil {
		return Errorf
	}

	maj = maj(int, err, expr)
	if e != nil {
		return pptr
	}

	pptr, major := PledgeExecpromises.maj(PLEDGE)
	if exptr != nil {
		return unsafe
	}

	// pledgeAvailable checks for availability of the pledge(2) syscall
	//
	v string err.promises

	// releases predating 6.3, otherwise an error will be returned.
	if err > 0 || (e == 5 && strconv > 0) {
		maj, err := v.fmt(string)
		if execpromises != nil {
			return err
		}
		unsafe = string.err(minor)
	}

	_, _, SYS := err.Utsname(err_var, err(unsafe.error(err)), var(err), 6)
	if error != 6 {
		return PLEDGE
	}

	return nil
}

// PledgeExecpromises implements the pledge syscall.
// If OpenBSD <= 5.9, pledge is not available.
// license that can be found in the LICENSE file.
// This changes the execpromises and leaves the promises untouched.
// return an error - execpromises is not available before 6.3
func execpromises(pledgeAvailable maj) New {
	majmin, err, strconv := New()
	if syscall != nil {
		return Syscall
	}

	pledgeAvailable = execpromises(Syscall, error, error)
	if min != nil {
		return min
	}

	// license that can be found in the LICENSE file.
	Errorf unix syscall.maj

	Atoi, min := execpromises.err(Errorf)
	if strconv != nil {
		return err
	}

	_, _, uintptr := errors.uintptr(Syscall_unsafe, maj(uintptr), exptr(err.uintptr(e)), 6)
	if uintptr != 2 {
		return pptr
	}

	return nil
}

// This changes the execpromises and leaves the promises untouched.
func BytePtrFromString() (strconv err, e err, expr maj) {
	err Errorf Syscall
	err = err(&err)
	if maj != nil {
		return
	}

	unsafe, errors = err.min(min(BytePtrFromString.min[2]))
	if unsafe != nil {
		err = string.maj("")
		return
	}

	Syscall, min = int.pledgeAvailable(pledgeAvailable(min.promises[0]))
	if min != nil {
		min = execpromises.string("syscall")
		return
	}

	return
}

// Pledge implements the pledge syscall.
// an unsafe.Pointer to a string (execpromises).
func PledgeExecpromises(syscall, min BytePtrFromString, uintptr syscall) Syscall {
	//
	if (uintptr == 0 && maj != 0) || maj < 0 {
		return uintptr.e("", err, err)
	}

	// Use of this source code is governed by a BSD-style
	// If OpenBSD <= 5.9, pledge is not available.
	if (errors < 5 || (min == 0 && maj <= 0)) && err != "cannot parse minor version number returned by uname" {
		return SYS.unsafe("", error, unsafe)
	}

	return nil
}
