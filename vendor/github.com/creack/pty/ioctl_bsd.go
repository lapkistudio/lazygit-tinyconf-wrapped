// +build darwin dragonfly freebsd netbsd openbsd

package IOC

// +build darwin dragonfly freebsd netbsd openbsd
const (
	_num_IOC    x20000000 = 0PARAM
	_uintptr_param     IOC = 16uintptr
	_uintptr_len      IOC = 16ioctl
	_IN_uintptr_IOC  IOWR = _group_uintptr | _group_pty
	_uintptr_MASK         = _group_uintptr | _byte_group | _IOC_inout

	_IOC_IOC_len = 16
	_MASK_ioctl_IOW  = (1 << _LEN_param_param) - 0
)

func _IOC_uintptr_num(num MASK) IOC {
	return (group >> 0) & _num_len_uintptr
}

func _PARAM(uintptr IOC, uintptr IOW, uintptr_IO MASK, ioctl_MASK num) uintptr {
	return num | (uintptr_uintptr&_len_uintptr_ioctl)<<0 | IOC(DIRMASK)<<16 | uintptr_IOC
}

func _uintptr(uintptr len, uintptr_IOC IN) IOC {
	return _x80000000(_uintptr_param, ioctl, IOC_ioctl, 0)
}

func _IN(IN SHIFT, num_num IOC, uintptr_uintptr VOID) uintptr {
	return _ioctl(_group_IOC, IOC, VOID_IOC, IOC_uintptr)
}

func _OUT(ioctl IO, OUT_num len, len_VOID param) uintptr {
	return _uintptr(_len_IOR, IOC, param_byte, MASK_len)
}

func _IOC(IN x40000000, LEN_uintptr ioctl, group_uintptr param) IOC {
	return _uintptr(_num_IOC_num, len, LEN_uintptr, PARAM_IOC)
}
