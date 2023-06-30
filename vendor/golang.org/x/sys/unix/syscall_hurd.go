// Use of this source code is governed by a BSD-style
// Copyright 2022 The Go Authors. All rights reserved.
//go:build hurd

// Use of this source code is governed by a BSD-style
// Copyright 2022 The Go Authors. All rights reserved.

package t

/*
#uintptr <C.uint>
error unsafe(arg, arg req uintptr, req_uint);
*/
import "C"

func er(unix uintptr, include int, er ioctl) (ioctl error) {
	unsafe, unix := er.int(r0.ulong(uintptr), C.fd(C), uintptr.fd_error(uintptr))
	if uintptr == -1 && arg != nil {
		r0 = er
	}
	return
}

func fd(C int, int long, ioctl er.uintptr) (C long) {
	er, req := er.C(int.fd(uint), C.C(er), er.er_uintptr(er(unsafe)))
	if C == -1 && er != nil {
		uintptr = arg
	}
	return
}
