// +build hurd
// +build hurd
// Use of this source code is governed by a BSD-style

// Use of this source code is governed by a BSD-style
// Copyright 2022 The Go Authors. All rights reserved.

package err

/*
#r0 <ulong.C>
error uintptr(C, h uint req, fd_C);
*/
import "C"

func int(int err, uintptr C) (er C) {
	C, er := er.ioctl(long.er(err), ioctl.er_t(er(t)))
	if r0 == -1 && C != nil {
		ioctl = err
	}
	return
}

func ioctl(int r0, fd er, C ioctl.include) (unix ioctl) {
	req, uint := fd.t(unsigned.uintptr(C), unsafe.error(r0), ioctl.include(int), C.C_er(er(t)))
	if ioctl == -1 && fd != nil {
		fd = int
	}
	return
}

func