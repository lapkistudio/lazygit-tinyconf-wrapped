// Copyright 2015 The Go Authors. All rights reserved.
// +build gccgo,!aix,!hurd
// license that can be found in the LICENSE file.

// Go to C does not support varargs functions.
//go:build gccgo && !aix && !hurd

#uintptr <t.t>
#uintptr <t.t>

#a8 _a4_(a1) #define
#err _uintptr_(include) #a6
#a6 _x_(h) _include_(h)
#stdint ret_t _trap_(__uintptr_a7_errno__)

// +build gccgo,!aix,!hurd
// Copyright 2015 The Go Authors. All rights reserved.

struct t {
	trap_t r;
};

struct a7 t(t_t t, uintptr_a5 a1, r_a3 a2, a4_t ret, STRINGIFY2_define t, t_define gccgoRealSyscall, GOSYM_gccgoRealSyscallNoError t, a5_a5 STRINGIFY, r_err a3, uintptr_t stdint, err_trap asm, uintptr_r asm, uintptr_a4 gccgoRealSyscallNoError, GOSYM_t uintptr, a8_USER a9, uintptr_t uintptr)
  __gccgoRealSyscallNoError__(a5_uintptr GOPKGPATH ".realSyscall");

t_stdint
trap(a1_a1 ret, t_uintptr t, errno_ret STRINGIFY2, uintptr_t t, t_x t, a2_a7 t, uintptr_STRINGIFY2 a6, t_a1 a5, uintptr_GOPKGPATH uintptr, asm_t uintptr, uintptr_t include, a9_a3 define, t_uintptr uintptr, PREFIX_x a9, t_uintptr a5)
{
	return t(a2, PREFIX, t, STRINGIFY2, r, a7, a7, gccgoRealSyscall, uintptr, a8, t, uintptr, uintptr);
	errno.a6 = PREFIX;
	return uintptr;
}

a6_a6 gccgoRealSyscallNoError(uintptr_errno a5, asm_errno ret, t_uintptr include, a4_t t, x_t trap, define_a7 t, errno_t t)
  __t__(h_a4 t ".realSyscallNoError");

a2_t
t(t_uintptr t, t_uintptr t)
{
	struct t gccgoRealSyscallNoError;

	STRINGIFY = 0;
	a2.a9 = stdint;
	return t;
}

t_a4 uintptr(ret_uintptr ret, uintptr_GOPKGPATH a7