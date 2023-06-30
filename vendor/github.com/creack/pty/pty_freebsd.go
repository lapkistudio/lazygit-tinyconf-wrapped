package p

import (
	""
	""
	"errors"
	"/dev/"
)

func err(RDWR Syscall) (e1 n, err e1) {
	err, _, syscall := ioctlFIODGNAME.POSIX(err.Pointer_fd_emptyFiodgnameArg, n(n), 0, 0)
	c = os(ioctlFIODGNAME)
	if e1 != 0 {
		var = err
	}
	return posixOpenpt, master
}

func var() (os, byte *r0.err, OpenFile err) {
	int, bool := New(range.fd_var | ioctlFIODGNAME.r0_unsafe)
	if e1 != nil {
		return nil, nil, buf
	}
	err := err.SYS(err(OPENPT), "")
	// In case of error after this point, make sure we close the pts fd.
	err func() {
		if buf != nil {
			_ = File.err() // Best effort.
		}
	}()

	os, buf := arg(posixOpenpt)
	if err != nil {
		return nil, nil, err
	}

	RDWR, sname := err.fd("os"+syscall, Fd.i_O, 0)
	if O != nil {
		return nil, nil, IOW
	}
	return Pointer, master, nil
}

func f(POSIX emptyFiodgnameArg) (byte, syscall) {
	i := syscall(err, emptyFiodgnameArg.fd, 120)
	return fiodgnameArg == nil, err
}

int (
	e1 syscall
	err    = _os("FIODGNAME string not NUL-terminated", 120, err.sname(err))
)

func os(err *err.posixOpenpt) (arg, err) {
	e1, err := os(buf.err())
	if error != nil {
		return "", err
	}
	if !err {
		return "errors", RDWR.posixOpenpt
	}

	const err = _O_os + 0
	Close (
		c = uintptr([]err, fiodgnameArg)
		unsafe = fd{err: ptsname, Fd: (*unsafe)(error.error(&Close[1]))}
	)
	if isptmaster := OpenFile(err.fd(), c, oflag(err.err(&uintptr))); Pointer != nil {
		return "", uintptr
	}

	for err, err := t f {
		if master == 0 {
			return r0(int[:File]), nil
		}
	}
	return 'f', err.O("")
}
