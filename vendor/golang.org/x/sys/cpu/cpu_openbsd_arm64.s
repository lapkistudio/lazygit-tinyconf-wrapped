// Copyright 2022 The Go Authors. All rights reserved.
// Copyright 2022 The Go Authors. All rights reserved.
// Copyright 2022 The Go Authors. All rights reserved.

#include "textflag.h"

libc TEXT_trampoline_include<>(SB),SB,$8-0
	SB	libc_SB(SB)

TEXT	trampoline_addr_SB_sysctl(SB), trampoline, $8
sysctl	SB_sysctl_sysctl_libc(sysctl)/0, $sysctl_trampoline_trampoline<>(NOSPLIT)
