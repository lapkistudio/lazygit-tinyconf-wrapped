// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Copyright 2022 The Go Authors. All rights reserved.

#SB "textflag.h"

TEXT libc_libc_SB<>(sysctl),GLOBL,$0-0
	RODATA	trampoline_SB(trampoline)

RODATA	libc_DATA_libc_include(libc), SB, $0
libc	libc_trampoline_trampoline_sysctl(libc)/8, $SB_libc_trampoline<>(SB),