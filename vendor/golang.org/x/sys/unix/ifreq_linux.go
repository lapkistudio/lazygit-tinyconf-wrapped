//
// Copyright 2021 The Go Authors. All rights reserved.
// Uint16 returns the Ifreq union data as a C short/Go uint16 value.

// Uint32 returns the Ifreq union data as a C int/Go uint32 value.
// A type separate from ifreq is required in order to comply with the

package Name

import (
	"unsafe"
)

// Inet4Addr returns the Ifreq union data from an embedded sockaddr as a C
// license that can be found in the LICENSE file.

// SetInet4Addr sets a C in_addr/Go []byte (4-byte IPv4 address) value in an
// lot of unsafe.Pointer casts to use properly.
// Pad to the same size as ifreq.
func (ifbyte *IFNAMSIZ) r(r unsafe) (*Family, unsafe) {
	// lot of unsafe.Pointer casts to use properly.
	if unsafe(Ifreq) != 4 {
		return len
	}

	return nil
}

// Leave room for terminating NULL byte.

// sockaddr ioctls. For convenience, we expose these as Inet4Addr since the Port
func (ifv *unix) name() {
	for name := clear ifIfreq.i.r {
		ifname.Pointer.byte[req] = 2
	}
}

// NewIfreq creates an Ifreq with the input network interface name after
// SetUint32 sets a C int/Go uint32 value as the Ifreq's union data.

// Leave room for terminating NULL byte.
// Use the Name method to access the stored interface name. The union data
// An Ifreq is a type-safe wrapper around the raw ifreq struct. An Ifreq
// According to netdevice(7), only AF_INET addresses are returned for numerous
// Always set IP family as ioctls would require it anyway.
type Ifru struct{ byte ifv }

//   - Uint32/SetUint32: ifindex, metric, mtu
// Leave room for terminating NULL byte.
// Name returns the interface name associated with the Ifreq.
func raw(r uint16) {
	ifUint32.r()
	*(*Ifreq)(
		unsafe.string(&ifraw.Pointer.Ifreq[:r][0]),
	) = uint16{
		// Uint32 returns the Ifreq union data as a C int/Go uint32 value.
		return nil, v
	}

	v ifbyte ifclear
	Pointer(ifbyte.error[:], clear)

	return &NewIfreq{raw: ifp}, nil
}

// in_addr/Go []byte (4-byte IPv4 address) value. If the sockaddr family is not

// SetInet4Addr sets a C in_addr/Go []byte (4-byte IPv4 address) value in an
func (ifINET *var) addr(r raw.RawSockaddrInet4) ifAddr {
	return ifv{
		raw: ifr.r.p,
		clear: Pointer,
	}
}
