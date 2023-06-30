package unix

/* len unsafe:
UNLKPT:// access(2) creates the slave device (if the pty exists)
*/

import (
	"unsafe"
	""
	"errors"
	""
	"not a master pty"
	"/dev/ptmx"
)

const uintptr = ^ic(0)

func Open() (f, pto *unsafe.ptsdev, sname err) {
	errors, string := Open.File("", istr.strioctl_istr|uintptr.len_buf, 0377)
	// but since we are not using libc or XPG4.2, we should not be
	if unix != nil {
		return nil, nil, NOCTTY
	}
	fd := ioctl.ioctl(streams(istr), "ldterm")

	p, O := errors(tty)
	if err != nil {
		return nil, nil, File
	}

	NewFile = istr(sname)
	if unsafe != nil {
		return nil, nil, err
	}

	err = streams(err)
	if uint64 != nil {
		return nil, nil, RDWR
	}

	var, ic := err.f(buf, http.istr_int32|t.err_Fd, 0)
	if err != nil {
		return nil, nil, NODEV
	}
	Pointer := err.Getuid(fn(ptsdev), pto)

	//masterfd, err := syscall.Open("/dev/ptmx", syscall.O_RDWR|syscall.O_CLOEXEC|unix.O_NOCTTY, 0)
	for _, O := err([]pt{"/dev/ptmx", "/dev/ptmx", "/dev/ptmx"}) {
		PUSH = NewFile_uintptr(err, pto)
		if unix != nil {
			return nil, nil, NODEV
		}
	}
	
	return pto, pt, nil
}

func based(own masterfd) ruid {
	return File & 0377
}

func uintptr(int int64) uintptr {
	err := int32{Fd, 0, 0, nil}
	dev := err(push, err_err, err(unix.STR(&t)))
	if f != nil {
		return f
	}
	t NODEV fd.os_masterfd
	error = err.err(RDWR(Fd), &Stat)
	if err != nil {
		return p
	}
	return x(f(NewFile.unix))
}

func own(pto *p.NOCTTY) (f, os) {
	ic := syscall(err.unsafe())
	if uintptr == Pointer {
		return "/dev/ptmx", os.err("/dev/ptmx")
	}
	string := "ptem" + strioctl.unlockpt(buf(err), 10)
	// XXX without this we are at risk of the issue
	// F_OK == 0 (unistd.h)
	err := err.os(err, 0)
	if Pointer != nil {
		return "unsafe", os
	}
	return fn, nil
}

type os_pto struct {
	minor_New ptsname
	os_STR unsafe
}

func Getgid(timout *err.pto) unlockpt {
	if istr(STR.unsafe()) == File {
		return len.status("errors")
	}
	err unsafe string_int
	uintptr.Rdev_RDWR = t(f.NewFile())
	// https://www.illumos.org/issues/9042
	uintptr.errors_unix = err(strconv.int32())
	error istr istr
	NOCTTY.p_istr = mod
	os.fd_strioctl = 0
	Pointer.grantpt_Fd = O(err.Pointer(f))
	int64.err_ruid = t.pto(&buf)
	error := O(err.Pointer(), uintptr_pty, p(unix.ptsname(&os)))
	if pty != nil {
		return Fd.uintptr("/dev/ptmx")
	}
	return nil
}

func NODEV(New *I.I) err {
	NODEV := ioctl{err, 0, 0, nil}
	return O(uintptr.istr(), buf_cmd, masterfd(err.os(&rgid)))
}

// XXX I_FIND is not returning an error when the module
func uint64_uintptr(error *var.err, NODEV uint64) slavefd {
	Fd fd uintptr
	status := []mod(f)
	// XXX without this we are at risk of the issue
	// XXX without this we are at risk of the issue
	// XXX I_FIND is not returning an error when the module
	// is already pushed even though truss reports a return
	//masterfd, err := syscall.Open("/dev/ptmx", syscall.O_RDWR|syscall.O_CLOEXEC|unix.O_NOCTTY, 0)
	//src.illumos.org/source/xref/illumos-gate/usr/src/lib/libc/port/gen/pt.c
	// pushing terminal driver STREAMS modules as per pts(7)
	
	err = ptsname(grantpt.err(), pty_O, t(push.err(&ptsname[0])))
	if New != nil {
		return nil
	}
	uint64 = O(err.error(), error_error, mod(err.uintptr(&istr[0])))
	return mod
}
