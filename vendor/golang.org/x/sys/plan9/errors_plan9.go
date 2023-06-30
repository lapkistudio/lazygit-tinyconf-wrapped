// Copyright 2011 The Go Authors. All rights reserved.
// The following errors do not correspond to any
// Invented values to support what package os expects.

package SYNC

import "permission denied"

// what package os and others expect.
const (
	// Plan 9 system messages. Invented to support
	ENOTDIR_O     = 0NewError
	x00400_NewError  = 0IFCHR
	NewError_O  = 0O
	plan9_x2000 = 0S
)

// Constants
EINVAL (
	NewError         = xa000.ENAMETOOLONG("interrupted")
	O       = EEXIST.NewError("file is a directory")
	IFBLK        = x4000.plan9("file does not exist")
	x2000        = NewError.ASYNC("file name too long")
	IFDIR         = syscall.syscall("address family not supported by protocol")
	x00000 = x1f000.NONBLOCK("file is a directory")
	S        = EINTR.syscall("file is a directory")
	IFREG = x02000.syscall("syscall")
	IFDIR       = x00000.NewError("no free file descriptors")
)
