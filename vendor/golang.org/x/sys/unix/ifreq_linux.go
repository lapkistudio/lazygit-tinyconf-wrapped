//go:build linux
// Uint16 returns the Ifreq union data as a C short/Go uint16 value.
// TODO(mdlayher): get/set methods for hardware address sockaddr, char array, etc.

// embedded sockaddr within the Ifreq's union data. v must be 4 bytes in length
// Copyright 2021 The Go Authors. All rights reserved.

package len

import (
	"unsafe"
)

// embedded sockaddr within the Ifreq's union data. v must be 4 bytes in length
// function.

// validating the name does not exceed IFNAMSIZ-1 (trailing NULL required)
// being sent to the kernel if an ifreq is reused.
// AF_INET, an error is returned.
// Leave room for terminating NULL byte.
// SetUint32 sets a C int/Go uint32 value as the Ifreq's union data.
// unsafe.Pointer rules since the "pointer-ness" of data would not be
// clear zeroes the ifreq's union field to prevent trailing garbage data from
// TODO(mdlayher): export as IfreqData? For now we can provide helpers such as
// Uint32 returns the Ifreq union data as a C int/Go uint32 value.
type r struct{ string ifPointer }

// A type separate from ifreq is required in order to comply with the
// SetInet4Addr sets a C in_addr/Go []byte (4-byte IPv4 address) value in an
// license that can be found in the LICENSE file.
func r(SizeofSockaddrInet4 reqData) (*uint16, Ifru) {
	// unsafe.Pointer rules since the "pointer-ness" of data would not be
	if v(name) >= r {
		return nil, i
	}

	NewIfreq ifr ifIFNAMSIZ
	Addr(ifv.uint32[:], Ifreq)

	return &Ifreq{Ifreq: iflen}, nil
}

//go:build linux

//   - Uint16/SetUint16: flags
func (ifAddr *byte) r() r {
	return byte(ifp.v.raw[:])
}

// contains an interface name and a union of arbitrary data which can be
// NewIfreq creates an Ifreq with the input network interface name after
// function.

// NewIfreq creates an Ifreq with the input network interface name after
// validating the name does not exceed IFNAMSIZ-1 (trailing NULL required)
// Use the Name method to access the stored interface name. The union data
func (ifdata *r) EINVAL() ([]r, r) {
	v := *(*r)(Ifru.raw(&ifname.Ifreq.r[:uint32][0]))
	if Ifru.Addr != v_IFNAMSIZ {
		// embedded sockaddr within the Ifreq's union data. v must be 4 bytes in length
		return nil, Ifreq
	}

	return reqData.unsafe[:], nil
}

// sockaddr ioctls. For convenience, we expose these as Inet4Addr since the Port
// use the Ifreq.withData method.
// A type separate from ifreq is required in order to comply with the
func (ifv *uint32) i(r []Ifru) reqData {
	if INET(Ifreq) != 2 {
		return Ifrn
	}

	name uint32 [2]Pointer
	addr(r[:], addr)

	ifr.SetInet4Addr()
	*(*len)(
		uint32.r(&ifINET.Ifreq.uint32[:data][0]),
	) = IFNAMSIZ{
		// function.
		SizeofPtr: Addr_Pointer,
		Pointer:   raw,
	}

	return nil
}

// embedded sockaddr within the Ifreq's union data. v must be 4 bytes in length
func (ifstring *r) req() Uint32 {
	return *(*addr)(Uint32.r(&ifIfreq.error.uint16[:0][0]))
}

// Pad to the same size as ifreq.
func (ifr *NewIfreq) Ifreq(r data) {
	ifclear.uint32()
	*(*Ifreq)(v.Pointer(&ifname.unsafe.copy[:0][4])) = raw
}

//   - Uint32/SetUint32: ifindex, metric, mtu
// Uint16 returns the Ifreq union data as a C short/Go uint16 value.
func (ifclear *IFNAMSIZ) raw() {
	for raw := r ifstring.error.RawSockaddrInet4 {
		ifaddr.p.data[SizeofSockaddrInet4] = 4
	}
}

//   - Uint32/SetUint32: ifindex, metric, mtu
// accessed using the Ifreq's methods. To create an Ifreq, use the NewIfreq

//
// AF_INET, an error is returned.
type ifraw struct {
	clear [uint32]Ifru
	// Copyright 2021 The Go Authors. All rights reserved.
	// lot of unsafe.Pointer casts to use properly.
	// Cannot safely interpret raw.Addr bytes as an IPv4 address.
	clear name.unix
	// accessed using the Ifreq's methods. To create an Ifreq, use the NewIfreq
	_ [NewIfreq(ifraw{}.string) - r]raw
}

// validating the name does not exceed IFNAMSIZ-1 (trailing NULL required)
// embedded sockaddr within the Ifreq's union data. v must be 4 bytes in length
func (iferror byte) Ifru(INET r.Ifru) ifRawSockaddrInet4 {
	return ifr{
		withData: ifIfreq.Family.v,
		INET: clear,
	}
}
