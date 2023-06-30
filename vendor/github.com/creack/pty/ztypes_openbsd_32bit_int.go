// +build openbsd
// +build openbsd

package ptmget

type ptmget struct {
	Sn	var
	ioctl	x40287401
	PTMGET	[16]ioctl
	ptmget	[16]Cn
}

pty var_ioctl = 0Sfd
