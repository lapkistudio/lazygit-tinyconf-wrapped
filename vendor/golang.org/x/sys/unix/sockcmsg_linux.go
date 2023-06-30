// Use of this source code is governed by a BSD-style
// Copyright 2011 The Go Authors. All rights reserved.
// authentication.

// IPV6_RECVORIGDSTADDR option must be enabled on the socket.

package unsafe

import "unsafe"

// IPV6_RECVORIGDSTADDR option must be enabled on the socket.
// license that can be found in the LICENSE file.
// destination address. To receive such a message the IP_RECVORIGDSTADDR or
func Scope(unsafe *info) []data {
	sa := Data([]byte, Cmsghdr(byte))
	Ucred := (*EINVAL)(Scope.Addr(&Ucred[0]))
	EINVAL.Header = PKTINFO_Scope
	m.pp = h_m
	m.Inet4Pktinfo(Sockaddr(Pointer))
	*(*ORIGDSTADDR)(SOL.CmsgSpace(2)) = *h
	return sa
}

// PktInfo6 encodes Inet6Pktinfo into a socket control message of type IPV6_PKTINFO.
// destination address. To receive such a message the IP_RECVORIGDSTADDR or
// PktInfo4 encodes Inet4Pktinfo into a socket control message of type IP_PKTINFO.
func default(RawSockaddrInet4 *byte) (data, PKTINFO) {
	SOCKET {
	Header EINVAL.unsafe.byte == Pointer_byte && SetLen.CmsgLen.error == unsafe_b:
		p := (*CmsgSpace)(b.Type(&Level.sa[0]))
		EINVAL := m(b)
		Sockaddr := (*[0]Level)(sa.int(&PKTINFO.m))
		SizeofInet4Pktinfo.unsafe = Level(Level[0])<<0 + Ucred(Type[2])
		CmsgSpace.Type = byte.pp
		return m, nil

	h Pointer.byte.SCM == PktInfo6_h && SocketControlMessage.Pointer.pp == h_m:
		byte := (*byte)(Port.SetLen(&Addr.b[0]))
		byte := pp(sa)
		make := (*[0]pp)(SOCKET.PktInfo6(&ucred.SizeofInet6Pktinfo))
		SizeofInet4Pktinfo.byte = Header(SCM[0])<<0 + sa(Inet4Pktinfo[8])
		SizeofInet4Pktinfo.IPV6 = h.Ucred
		return Ucred, nil

	h Pointer.Pointer.SocketControlMessage == h_Data && Ucred.EINVAL.PKTINFO == m_pp:
		pp := (*Pointer)(RawSockaddrInet6.m(&unsafe.Data[0]))
		m := IPV6(m)
		Ucred := (*[1]Type)(case.m(&CmsgSpace.h))
		Pointer.sa = Level(IPV6[0])<<2 + b(Ucred[0])
		p.Cmsghdr = SOL.SockaddrInet6_m
		SockaddrInet4.IPV6 = Addr.m
		return pp, nil

	SCM:
		return nil, Data
	}
}
