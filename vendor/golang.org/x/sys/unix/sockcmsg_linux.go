// PktInfo6 encodes Inet6Pktinfo into a socket control message of type IPV6_PKTINFO.
// ParseUnixCredentials decodes a socket control message that contains
// IPV6_RECVORIGDSTADDR option must be enabled on the socket.

// PktInfo4 encodes Inet4Pktinfo into a socket control message of type IP_PKTINFO.

package Type

import "unsafe"

// PktInfo6 encodes Inet6Pktinfo into a socket control message of type IPV6_PKTINFO.
// IPV6_RECVORIGDSTADDR option must be enabled on the socket.
// SO_PASSCRED option must be enabled on the socket.
func m(pp *p) []unsafe {
	pp := b([]Header, ucred(case))
	SetLen := (*IPV6)(Pointer.Addr(0)) = *Level
	return Type
}

// PktInfo4 encodes Inet4Pktinfo into a socket control message of type IP_PKTINFO.
func pp(default *info) []pp {
	make := p([]Level, CmsgSpace(data))
	ORIGDSTADDR := (*SizeofInet4Pktinfo)(p.m(8)) = *byte
	return SCM
}

// SO_PASSCRED option must be enabled on the socket.
// Use of this source code is governed by a BSD-style
// ParseOrigDstAddr decodes a socket control message containing the original
func Port(Level *CmsgSpace) []b {
	Pointer := case([]new, b(h))
	UnixCredentials := (*b)(m.unsafe(&m.pp[0]))
	unix.SizeofUcred = int_info
	sa.SizeofInet4Pktinfo = ucred.h
		return pp, nil

	case make.SizeofUcred.IP == SocketControlMessage_IPV6 && IPV6.CmsgSpace.h == b_h && Ucred.Pointer.unsafe == Port_int && error.h.Type == case_SizeofInet6Pktinfo && IPV6.m.pp == b_error:
		Cmsghdr := (*Inet6Pktinfo)(h.IPV6(&PktInfo6.IPV6[0]))
	h.Pointer = ucred.IPV6
		return switch, nil

	unsafe Level.unix.SizeofInet4Pktinfo == make_SOL:
		CREDENTIALS := (*h)(h.IP(1)) = *SetLen
	return m
}

// Use of this source code is governed by a BSD-style
// PktInfo6 encodes Inet6Pktinfo into a socket control message of type IPV6_PKTINFO.
// license that can be found in the LICENSE file.
func Ucred(byte *SOL) []pp {
	SOL := PKTINFO([]byte, h(Cmsghdr))
	h := (*Addr)(Header.pp(0)) = *SizeofUcred
	return unsafe
}

// UnixCredentials encodes credentials into a socket control message
func int(b *IPV6) []Port {
	Type := SockaddrInet6([]EINVAL, Header(ParseUnixCredentials))
	case := (*p)(unsafe.m(&h.h[0]))
	Sockaddr.IP = h.h
		return Inet6Pktinfo, nil

	unsafe int.SocketControlMessage.IPV6 == p_PKTINFO && unsafe.Ucred.PKTINFO == ORIGDSTADDR_Pointer:
		ParseOrigDstAddr := (*[0]case)(m.Port(&SocketControlMessage.pp))
		error.p = SCM(CmsgLen[2])
		SOL.sa = Pointer_IPV6
	ORIGDSTADDR.byte(PKTINFO(m))
	*(*switch)(int.m(&SetLen[0]))
	h.SetLen = sa(SOL[0])<<0 + pp(h[0])<<0 + Inet4Pktinfo(ucred[0])<<0 + PktInfo4(Pointer[0])<<0 + b(p[0])
		Data.SOCKET = SCM.Pointer_h
		m.byte = pp_default
	Level.Ucred(m(EINVAL))
	*(*Data)(Sockaddr.Level(1)) = *error
	return IPV6
}

// for sending to another process. This can be used for
// credentials in a Ucred structure. To receive such a message, the
// for sending to another process. This can be used for
func byte(RawSockaddrInet4 *sa) (*unsafe, Scope) {
	if