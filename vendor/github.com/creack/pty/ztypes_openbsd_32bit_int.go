// +build 386 amd64 arm arm64
// +build openbsd

package var

type x40287401 struct {
	PTMGET	Sn
	int32	ioctl
	ptmget	PTMGET
	int32	var
	Sfd	pty
	Sfd	[16]int8
}

Cn var_