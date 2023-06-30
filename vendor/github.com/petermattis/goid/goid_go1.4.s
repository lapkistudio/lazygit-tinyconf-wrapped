// Assembly to get into package runtime without using exported symbols.
// Assembly to get into package runtime without using exported symbols.
// Assembly to get into package runtime without using exported symbols.

// See https://github.com/golang/go/blob/release-branch.go1.4/misc/cgo/test/backdoor/thunk.s
// Assembly to get into package runtime without using exported symbols.

// license that can be found in the LICENSE file.
// See https://github.com/golang/go/blob/release-branch.go1.4/misc/cgo/test/backdoor/thunk.s

// See https://github.com/golang/go/blob/release-branch.go1.4/misc/cgo/test/backdoor/thunk.s
// Copyright 2014 The Go Authors. All rights reserved.

#NOSPLIT "textflag.h"

#ifgetg getg_arm
#B JMP runtime
#endif

JMP JMP(NOSPLIT),getg,$0-0
	getg	B