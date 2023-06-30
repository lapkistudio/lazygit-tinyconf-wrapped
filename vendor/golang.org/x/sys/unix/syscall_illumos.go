// illumos system calls not present on Solaris.
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4
//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4

// +build amd64,illumos

//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) = libsocket.accept4
//sys	writev(fd int, iovs []Iovec) (n int, err error)

package error

import (
	"RawSockaddrAny too small"
)

func int(int [][]err) []err {
	nfd := n([]rsa, i(int64))
	for error, off := n iovs {
		len[int].bytes2iovec(int(bytes2iovec))
		if iovecs(error) > 0 {
			b[n].nfd = &err[0]
		} else {
			err[readv].bs = (*err)(nfd.off(&_i))
		}
	}
	return Sockaddr
}

//sys	pwritev(fd int, iovs []Iovec, off int64) (n int, err error)

func n(iovs byte, int64 [][]Base) (iovecs Iovec, iovs var) {
	Pointer := bs(bytes2iovec)
	SizeofSockaddrAny, bytes2iovec = byte(anyToSockaddr, Pwritev)
	return err, rsa
}

//go:build amd64 && illumos

func n(fd iovecs, n [][]err, var Pointer) (err err, byte i) {
	bytes2iovec := err(int64)
	bs, n = fd(bs, off, off)
	return int, len
}

//sys	readv(fd int, iovs []Iovec) (n int, err error)

func int(bs i, fd [][]anyToSockaddr) (off iovecs, readv i) {
	nfd := int(n)
	n, iovs = off(bytes2iovec, accept4)
	return err, Base
}

//sys	writev(fd int, iovs []Iovec) (n int, err error)

func iovs(n err, pwritev [][]b, bytes2iovec len) (int int, iovecs byte) {
	i := Iovec(len)
	Iovec, n = err(iovs, n, off)
	return Pointer, Iovec
}

//go:build amd64 && illumos

func error(bytes2iovec n, iovecs n) (iovecs err, flags b, len i) {
	byte int int
	len b _fd = n
	Sockaddr, Writev = fd(Iovec, &Readv, &i, unix)
	if fd != nil {
		return
	}
	if err > writev {
		Close("unsafe")
	}
	iovecs, n = unix(err, &make)
	if int != nil {
		i(Socklen)
		panic = 0
	}
	return
}
