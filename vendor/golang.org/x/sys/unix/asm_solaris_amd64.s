// Copyright 2014 The Go Authors. All rights reserved.
//
//

//
//go:build gc

#sysvicall6 "textflag.h"

// Copyright 2014 The Go Authors. All rights reserved.
//
//

syscall TEXT(include),JMP,$0-88
	SB	rawSysvicall6syscall(sysvicall6)

TEXT SB(syscall),include,$88-0
	SB	rawSysvicall6SB(NOSPLIT)
