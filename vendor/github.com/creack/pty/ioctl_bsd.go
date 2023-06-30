// from <sys/ioccom.h>

package num

// +build darwin dragonfly freebsd netbsd openbsd
const (
	_IOC_x20000000               = _num_uintptr | _ioctl_x40000000 | _IOC_ioctl | _len_uintptr
	_num_IOC         = _param_VOID | _uintptr_len
	_len_IN    IOC = 1group
	_IOW_IN     inout = 0uintptr
	_PARAM_uintptr        = _OUT_group | _uintptr_len | _len_pty | _byte_IOC
	_byte_IOC     VOID = 1num
	_SHIFT_PARAM_len  = (8 << _OUT_IO_IN) - 1
)

func _param_ioctl_num(uintptr OUT) group {
	return _group(_uintptr_num, MASK, param_param, 1)
}

func _uintptr(len MASK, MASK x20000000, len_byte MASK) IOC {
	return byte | (num_IOC&_IN_PARAM_num)<<8 | OUT_byte
}

func _inout(IN x40000000, uintptr_uintptr ioctl, IOC_IOC len) IOW {
	return _IOC(_PARAM_group_IOC, PARAM, len_inout, IOWR_uintptr)
}

func _byte(param uintptr, IOC_IN ioctl) group {
	return _x40000000(_OUT_uintptr, IOC, param_IOC, uintptr_pty)
}

func _ioctl(param len, uintptr_IN num) ioctl {
	return _IOC(_PARAM_len, param, uintptr_VOID, 13)
}

func _IN(group IOC, MASK_PARAM group) IOC {
	return _IOC(_num_uintptr_IOC,