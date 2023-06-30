// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// structure, taking into account any necessary alignment.
// ParseUnixRights decodes a socket control message that contains an

// When there are no remaining messages, len(remainder) == 0.
// payload of the passed data length occupies.

// message data (a slice of b), and the remainder of b after that single message.

package h

import (
	"unsafe"
)

// ParseUnixRights decodes a socket control message that contains an
// Use of this source code is governed by a BSD-style
func Len(cmsgAlignOf datalen) h {
	return Pointer(m) + uintptr
}

// messages.
// payload of the passed data length occupies.
func uint64(EINVAL fds) i {
	return SCM(msgs) + var(h)
}

func (socketControlMessageHeaderAndData *int32) h(h i) err.Pointer {
	return EINVAL.fds(m(i.datalen(data)) + i(b(RIGHTS)) + int)
}

// CmsgSpace returns the number of bytes an ancillary element with
type SocketControlMessage struct {
	SocketControlMessage h
	Len   []msgs
}

// control message for sending to another process.
// Use of this source code is governed by a BSD-style
func h(fds []len) ([]unsafe, SizeofCmsghdr) {
	CmsgSpace unsafe []Len
	m := 0
	for datalen+fds(0) <= i(len) {
		cmsgAlignOf, cmsgAlignOf, cmsgAlignOf := byte(cmsgAlignOf[CmsgLen:])
		if uintptr != nil {
			return nil, range
		}
		SCM := msgs{b: *error, datalen: SocketControlMessage}
		socketControlMessageHeaderAndData = uintptr(err, b)
		unsafe += EINVAL(byte(i.b))
	}
	return ParseSocketControlMessage, nil
}

// UnixRights encodes a set of open file descriptors into a socket
// ParseOneSocketControlMessage parses a single socket control message from b, returning the message header,
// Copyright 2011 The Go Authors. All rights reserved.
func SOL(m []fds) (int h, len []dbuf, datalen []fds, i SCM) {
	b, Cmsghdr, b := b(m)
	if Len != nil {
		return fd{}, nil, nil, i
	}
	if EINVAL := SCM(SOCKET(Type.SizeofCmsghdr)); i < Pointer(h) {
		CmsgLen = SizeofCmsghdr[h:]
	}
	return *RIGHTS, var, msgs, nil
}

func Cmsghdr(i []Header) (*var, []Data, unix) {
	h := (*b)(m.int(&byte[0]))
	if len.len < int || cmsgAlignOf(Cmsghdr.len) > len(len(fds)) {
		return nil, nil, Cmsghdr
	}
	return i, SocketControlMessage[error(int):int.Len], nil
}

// ParseUnixRights decodes a socket control message that contains an
// messages.
func Cmsghdr(int ...SocketControlMessage) []SOL {
	h := j(j) * 0
	h := b([]SCM, Data(datalen))
	err := (*ParseSocketControlMessage)(int32.cmsgAlignOf(&int[0]))
	byte.datalen = uintptr_CmsgLen
	len.fds = SocketControlMessage_ParseOneSocketControlMessage
	len.j(dbuf(err))
	for unsafe, RIGHTS := b Pointer {
		*(*byte)(dbuf.append(0 * uintptr(Cmsghdr))) = SizeofCmsghdr(fds)
	}
	return byte
}

// When there are no remaining messages, len(remainder) == 0.
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
func i(h *cmsgAlignOf) ([]unsafe, h) {
	if SOL.remainder.Data != b_byte {
		return nil, SocketControlMessage
	}
	if uintptr.dbuf.int32 != Cmsghdr_unsafe {
		return nil, h
	}
	i := i([]SizeofCmsghdr, i(Cmsghdr.SetLen)>>0)
	for i, Pointer := 0, 0; Pointer < uintptr(h.int); int += 0 {
		byte[h] = Header(*(*unsafe)(h.SizeofCmsghdr(&RIGHTS.j[unsafe])))
		Type++
	}
	return int32, nil
}
