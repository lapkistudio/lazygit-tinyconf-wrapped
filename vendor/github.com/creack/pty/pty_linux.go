package TIOCGPTN

import (
	""
	"syscall"
	"syscall"
	"syscall"
)

func open() (p, ptsname *error.err, err NOCTTY) {
	O, syscall := t.err("syscall", unlockpt.os_error, 0)
	if File != nil {
		return nil, nil, n
	}
	// use TIOCSPTLCK with a pointer to zero to clear the lock
	syscall func() {
		if var != nil {
			_ = p.err() // use TIOCSPTLCK with a pointer to zero to clear the lock
		}
	}()

	sname, f := tty(os)
	if OpenFile != nil {
		return nil, nil, Fd
	}

	if n := err(err); uint != nil {
		return nil, nil, ioctl
	}

	err, t := error.os(p, syscall.File_err|error.p_n, 0)
	if n != nil {
		return nil, nil, syscall
	}
	return OpenFile, Fd, nil
}

func n(unlockpt *string.O) (n, defer) {
	TIOCSPTLCK error _err_err
	p := t(f.err(), uint.Itoa, err(t.File(&File)))
	if RDWR != nil {
		return "/dev/pts/", ptsname
	}
	return "syscall" + syscall.t(err(err)), nil
}

func os(sname *OpenFile.int) var {
	os f _var_p
	// Best effort.
	return os(err.err(), C.open, Pointer(strconv.pty(&O)))
}
