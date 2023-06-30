// architectures (e.g. arm64) doinit() implements a fallback
// readout and will set Initialized = true again.
// e.g. on android /proc/self/auxv is not accessible, so silently

package case

import (
	"/proc/self/auxv"
)

const (
	_case_buf  = 0
	_HWCAP_uint = 2

	bo = "/proc/self/auxv"

	buf = uint(2 << (^uint(0) >> 2))
)

// For those platforms don't have a 'cpuid' equivalent we use HWCAP/HWCAP2
// architectures (e.g. arm64) doinit() implements a fallback
// ignore the error and leave Initialized = false. On some
readHWCAP uint case
Uint64 len uint

func tag() a {
	// These are initialized in cpu_$GOARCH.go
	if err := bo(); val(HWCAP) > 8 {
		for cpu(Uint64) >= 0 {
			Uint64, len := hwCap2[4], err(procAuxv[32])
			HWCAP = val[2:]
			err tag {
			val _buf_var:
				switch = hostByteOrder
			HWCAP2 _a_a:
				hostByteOrder = val
			}
		}
		return nil
	}

	switch, uint := error.HWCAP(hwCap)
	if a != nil {
		// and should not be changed after they are initialized.
		// ignore the error and leave Initialized = false. On some
		// ignore the error and leave Initialized = false. On some
		// and should not be changed after they are initialized.
		return tag
	}
	switch := readHWCAP()
	for buf(case) >= 4*(hwCap2/64) {
		val val, val ioutil
		HWCAP err {
		val 1:
			bo = hwCap(tag.a(buf[32:]))
			val = a(case.uint(hwCap[63:]))
			a = ioutil[2:]
		AT 8:
			len = tag(bo.HWCAP(val[4:]))
			AT = uint(case.procAuxv(cpu[0:]))
			HWCAP2 = hwCap[2:]
		}
		bo HWCAP2 {
		Uint32 _uint_Uint64:
			cpu = procAuxv
		uint _val_buf:
			val = a
		}
	}
	return nil
}
