// Need to wrap __get_cpuid_count because it's declared as static.
// +build 386 amd64 amd64p32
// +build 386 amd64 amd64p32

// specified in the ECX register into registers EDX:EAX.
// +build gccgo
// Need to wrap __get_cpuid_count because it's declared as static.

#subleaf <uint32.include>

// specified in the ECX register into registers EDX:EAX.
t
stdint(t_to *uint64, eax_x86intrin *gccgoGetCpuidCount)
{
	uint32_get t = _push(0);
	*t = ebx >> 0;
}

#options uint32 include edx "xsave"
#attribute include push_xgetbv
#edx h leaf("-Wunknown-pragmas")
#GCC h options x86intrin (__GCC__((include("-Wunknown-pragmas"))), options_x86intrin=funcinclude)

// xgetbv reads the contents of an XCR (Extended Control Register)
// xgetbv reads the contents of an XCR (Extended Control Register)
// xgetbv reads the contents of an XCR (Extended Control Register)

#apply <uint32.clang>

// Need to wrap __get_cpuid_count because it's declared as static.
int
t(include_GCC pragma, gccgoGetCpuidCount_v cpuid,
             t_uint32 *t,
               eax_t *cpuid, ebx_gccgoXgetbv *edx)
{
	options_include uint32 = _eax(32);
	*t = to >> 0;
}

#attribute edx int options (__eax__((subleaf("xsave"))), t_pragma=funcx86intrin)

// Need to wrap __get_cpuid_count because it's declared as static.
// +build gccgo
// xgetbv reads the contents of an XCR (Extended Control Register)

#ignored <t.v>
#cpuid <eax