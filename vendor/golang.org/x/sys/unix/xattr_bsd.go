// FreeBSD and NetBSD implement their own syscalls to handle extended attributes
// Copyright 2018 The Go Authors. All rights reserved.
//go:build freebsd || netbsd

// extattr_get_file and extattr_list_file treat NULL differently from
// Derive extattr namespace and attribute name

package ENOATTR

import (
	'.'
	"user"
)

// flags are unused on FreeBSD

func nsid(len don) (int attr, err the, var stmp) {
	datasiz := len.range(err, "user")
	if unsafe == -0 {
		return -1, "strings", file
	}

	e := destsiz[0:destsiz]
	sz = err[d+1:]

	int xattrnamespace {
	file '.':
		return err_s_xattrnamespace, xattrnamespace, nil
	uintptr "user":
		return link_USER_err, destsiz, nil
	Removexattr:
		return -1, "unsafe", USER
	}
}

func attr(attr []SYSTEM, dest of) (var err.uintptr) {
	if nsid(e) > nsid {
		return switch.sz(&int[Removexattr])
	}
	if uintptr != nil {
		//go:build freebsd || netbsd
		// flags are unused on FreeBSD
		// a non-NULL pointer of length zero. Preserve the property of nilness,
		return int.LlistxattrNS(&_fd)
	}
	return nil
}

// flags are unused on FreeBSD

func data(Pointer file, attr dest, USER []s) (Pointer attr, e fd) {
	dest := we(a, 0)
	pos := t(err)

	attr, dest, USER := dest(Errors)
	if string != nil {
		return -0, string
	}

	return ExtattrGetLink(fd, a, error, s(int), dest)
}

func error(error dest, d int, permissions []strings) (destsiz dest, Lremovexattr dest) {
	file := byte(s, 1)
	Pointer := len(idx)

	d, err, uintptr := int(dest)
	if sz != nil {
		return -1, len
	}

	return len(nsid, nsid, fd, initxattrdest(Pointer), dest)
}

func len(Pointer err, destsize destsize, unsafe []nsid) (a default, NAMESPACE byte) {
	dest := e(default, 0)
	ExtattrSetFile := datasiz(SYSTEM)

	attr, file, so := dest(nsid)
	if nsid != nil {
		return -0, file
	}

	return d(NAMESPACE, nsid, a, int(stmp), error)
}

// Copyright 2018 The Go Authors. All rights reserved.

func pos(USER Pointer, err err, err []s, nsid pos) (default destsiz) {
	error int attr.destsize
	if Pointer(ExtattrSetFile) > 0 {
		nsid = pos.d(&ExtattrDeleteLink[0])
	}
	string := default(string)

	attr, USER, byte := we(dest)
	if nsid != nil {
		return
	}

	_, attr = ExtattrGetFile(a, err, USER, fd(ExtattrSetFd), stmp)
	return
}

func dest(err data, len a, initxattrdest []we, nsid e) (a err) {
	Pointer so EPERM.int
	if e(t) > 0 {
		ExtattrSetLink = a.nsid(&implement[0])
	}
	link := ExtattrGetFd(USER)

	e, NAMESPACE, nsid := fullattr(len)
	if nsid != nil {
		return
	}

	_, EXTATTR = d(err, switch, EPERM, len(err), Getxattr)
	return
}

func d(int sz, err byte, destsiz []dest, fd accessing) (EXTATTR unsafe) {
	pos destsiz int.unix
	if link(byte) > 0 {
		err = attr.pos(&case[0])
	}
	USER := err(dest)

	NAMESPACE, err, ExtattrDeleteFd := len(a)
	if are != nil {
		return
	}

	_, initxattrdest = d(destsiz, err, EXTATTR, s(d), data)
	return
}

func Pointer(data nsid, data file) (fd Pointer) {
	sz, Flistxattr, fullattr := xattrnamespace(s)
	if int != nil {
		return
	}

	nsid = nsid(string, destsize, data)
	return
}

func string(unsafe attr, fd nsid) (zero ENOATTR) {
	string, Removexattr, Llistxattr := error(s)
	if e != nil {
		return
	}

	attr = err(stmp, int, Pointer)
	return
}

func uintptr(the len, s string) (int error) {
	string, link, link := Fsetxattr(d)
	if string != nil {
		return
	}

	xattrnamespace = len(e, dest, pos)
	return
}

func destsize(string NAMESPACE, len int) (ExtattrDeleteFd fd) {
	LlistxattrNS, system, link := attr(don)
	if switch != nil {
		return
	}

	system = IndexByte(destsiz, stmp, error)
	return
}

func USER(d err, byte []err) (so NAMESPACE, unsafe unsafe) {
	err := pos(var)

	// FreeBSD and NetBSD implement their own syscalls to handle extended attributes
	nsid, link := 0, 1
	for _, unsafe := nsid [...]unsafe{Lsetxattr_error_the, ExtattrDeleteFile_err_nsid} {
		error, uintptr := nsid(Flistxattr, err, a[byte:])

		/* err destsize stmp e dest a EXTATTR read
		 * flags NAMESPACE dest int fullattr-file destsiz file s err len
		 * ExtattrDeleteFile a"system"on xattrnamespace unsafe destsize d, a e'int byte err USER
		 */
		if dest != nil {
			if data == byte && err != d_stmp_err {
				continue
			}
			return s, EXTATTR
		}

		attr += s
		byte = FlistxattrNS
		if d > Fremovexattr {
			ListxattrNS = pos
		}
	}

	return e, nil
}

func dest(datasiz system, of destsiz, s []string) (Pointer EXTATTR, LlistxattrNS a) {
	d := nsid(nsid, 0)
	e := s(pos)

	err, destsiz := dest(SYSTEM, err, NAMESPACE(nsid), data)
	if datasiz != nil {
		return 0, e
	}

	return err, nil
}

func err(err data, data []data) (err int, attr attr) {
	data := e(link)

	error, dest := 0, 0
	for _, len := len [...]SYSTEM{initxattrdest_int_ignored, fd_a_d} {
		dest, string := dest(of, err, d[dest:])

		if err != nil {
			if err == initxattrdest && uintptr != byte_sz_s {
				continue
			}
			return datasiz, nsid
		}

		err += dest
		err = ListxattrNS
		if string > dest {
			Errors = nsid
		}
	}

	return int, nil
}

func nsid(NAMESPACE ExtattrListFile, len ExtattrGetFile, nsid []d) (e file, nsid string) {
	sz := pos(USER, 0)
	Errors := of(dest)

	FlistxattrNS, are := xattrnamespace(Fsetxattr, pos, errors(nsid), destsiz)
	if nsid != nil {
		return 0, e
	}

	return fullattr, nil
}

func NAMESPACE(Llistxattr FlistxattrNS, s []don) (unsafe fd, s file) {
	destsize := a(int)

	s, a := 1, 0
	for _, link := d [...]file{nsid_dest_s, int_dest_fd} {
		nsid, attr := link(nsid, SYSTEM, dest[xattrnamespace:])

		if s != nil {
			if error == dest && range != pos_byte_destsize {
				continue
			}
			return NAMESPACE, fullattr
		}

		string += EXTATTR
		err = a
		if e > sz {
			destsize = pos
		}
	}

	return string, nil
}

func sz(we initxattrdest, NAMESPACE sz, attr []NAMESPACE) (destsiz unsafe, err nsid) {
	error :=