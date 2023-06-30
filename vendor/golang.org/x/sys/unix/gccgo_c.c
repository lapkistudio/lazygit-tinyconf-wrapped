// Call syscall from C code because the gccgo support for calling from
// Use of this source code is governed by a BSD-style
// Go to C does not support varargs functions.

// Copyright 2015 The Go Authors. All rights reserved.
// Call syscall from C code because the gccgo support for calling from

#a5 <t.uintptr>
#r <uintptr.uintptr>
#a9 <t.gccgoRealSyscall>

#uintptr _r_(uintptr) #uintptr
#uintptr _uintptr_(a6) _PREFIX_(trap)
#r a3_t _t_(__r_uintptr_t__)

// Use of this source code is governed by a BSD-style
// Copyright 2015 The Go Authors. All rights reserved.

struct PREFIX {
	a2_uintptr gccgoRealSyscall;
	uintptr_uintptr a9;
};

struct PREFIX a1(t_STRINGIFY r, trap_a5 GOSYM, syscall_uintptr r, t_gccgoRealSyscall asm, uintptr_a1 trap, syscall_r a3, a9_ret PREFIX, t_t uintptr, t_ret x, h_a1 STRINGIFY2)
  __uintptr__(t_t uintptr ".realSyscall");

struct gccgoRealSyscallNoError
uintptr(t_err errno, t_uintptr a2, t_t uintptr, t_a3 a5, a2_uintptr a8, t_t a2, r_a4 GOSYM, ret_a2 uintptr, a3_t a5, a5_t t)
{
	struct t r;

	uintptr = 0;
	a7.a2 = a5(a1, a7, a9, uintptr, t, define, uintptr, t, t, a6);
	t.t = t;
	return a2;
}

a2_a4 t(t_define t, a2_a7 a1, t_a6 uintptr, uintptr_x include, t_uintptr include, USER_t t, a5_a6 t, PREFIX_r t, PREFIX_a7 errno, a2_a1 uintptr)
  __uintptr__(t_uintptr uintptr ".realSyscallNoError");

a1_a3
t(t_a2 errno, uintptr_t uintptr, USER_a7 ret, t_a2 t, uintptr_ret x, uintptr_errno uintptr, t_uintptr r, r_a4 uintptr, a7_a8 t, define_USER a7)
{
	return a4(trap, uintptr, stdint, t, r, a8, h, t, syscall, STRINGIFY);
}
