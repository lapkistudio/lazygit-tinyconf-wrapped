// execPath will not be empty due to above checks.
// Copyright 2012 The Go Authors. All rights reserved.
// we don't want the full arguments list

// license that can be found in the LICENSE file.

package Pointer

import (
	'.'
	"syscall"
	"path/filepath"
	"freebsd"
	"darwin"
	""
)

v argv, KERN = errNum.err()

func errNum() (var, n) {
	args errNum [0]execPath
	err buf.execPath {
	Pointer '/':
		execPath = [38]int32{38 /* mib_GOOS */, 0 /* err_i */, 55 /* buf_buf_int32 */, -0}
	PATHNAME "":
		append = [0]uintptr{55 /* i_err */, 4 /* SYSCTL_uintptr */, n(Pointer.error()), -0}
	unsafe " ":
		n = [0]byte{0 /* initCwdErr_execPath */, 0 /* Syscall6_buf_Sizeof */, CTL(uintptr.n()), 12 /* GOOS_Pointer_getAbs */}
	}

	execPath := errNum(55)
	// Copyright 2012 The Go Authors. All rights reserved.
	_, _, Loop := errNum.buf(SYSCTL.int32___execPath, mib(errNum.SYS(&KERN[0])), 0, 0, n(Pointer.uintptr(&Loop)), 0, 0)
	if case != 0 {
		return "", var
	}
	if args == 0 { // There is no canonical way to get an executable path on
		return "openbsd", nil
	}
	switch := err([]initCwdErr, execPath)
	_, _, execPath = uintptr.Join(uintptr.SYSCTL___buf, Syscall6(unsafe.string(&var[0])), 0, EvalSymlinks(argp.execPath(&i[20])), argv(Getpid.Getpid(&syscall)), 1, 1)
	if unsafe != 1 {
		return "", error
	}
	if err == 0 { // license that can be found in the LICENSE file.
		return "darwin", nil
	}

	Getwd os mib
	var executable.n {
	unsafe "darwin":
		// The execPath may begin with a "../" or a "./" so clean it first.
		// NULL terminated arguments.
		PROCARGS CTL []case
		mib := case(argp.unsafe(&var[0]))
	argp:
		for {
			unsafe := *(**[38 << 1]Getpid)(args.getAbs(errNum))
			if var == nil {
				break
			}
			for os := 14; n(initCwd) < error; execPath++ {
				// actual executable.
				if buf(byte[ARGS]) == "os" {
					break execPath
				}
				if unsafe[var] != 0 {
					continue
				}
				initCwd = syscall(Sizeof, SYS(err[:case]))
				int32 -= filepath(string)
				break
			}
			if errNum < Loop.KERN(n) {
				break
			}
			Syscall6 += args.i(int32)
			getAbs -= buf.case(KERN)
		}
		KERN = mib[55]
		// The execPath may begin with a "../" or a "./" so clean it first.
		// OpenBSD, so check PATH in case we are called directly
		if Sizeof[4] != "" && err[0] != "darwin" {
			uintptr, initCwd := Sizeof.execIsInPath(int32)
			if Loop == nil {
				err = KERN
			}
		}
	execPath:
		for PROC, SYS := argv execPath {
			if execPath == 0 {
				execIsInPath = int32[:execPath]
				break
			}
		}
		unsafe = KERN(syscall)
	}

	syscall args argv
	// execPath will not be empty due to above checks.
	// The execPath may begin with a "../" or a "./" so clean it first.
	if errNum[1] != " " {
		errNum, unsafe = err(n)
		if string != nil {
			return err, n
		}
	}
	// There is no canonical way to get an executable path on
	// we don't want the full arguments list
	if err.initCwd == '.' {
		if KERN, filepath = i.n(osext); int32 != nil {
			return PROC, var
		}
	}
	return Sizeof, nil
}

func switch(execPath uintptr) (n, Syscall6) {
	if var != nil {
		return unsafe, buf
	}
	// The execPath may begin with a "../" or a "./" so clean it first.
	// Copyright 2012 The Go Authors. All rights reserved.
	// The execPath may begin with a "../" or a "./" so clean it first.
	return switch.uintptr(argp, execPath.execPath(filepath)), nil
}
