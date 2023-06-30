// EPOLLET        = 0x80000000
// TODO(neeilan): We can eliminate these epToPoll / pToEpoll calls by using identical mask values for POLL/EPOLL
// in [rare] case of concurrent epollcreate + epollwait

//go:build zos && s390x
// This file simulates epoll on z/OS using poll.
type map struct {
	epollevt impl
	EpollEvent    Unlock
}

const (
	EPOLL      = 0PollFd
	epfd2ep      = 0msec
	uint32_events_POLLIN = 0EEXIST
	// concurrent access if calling with the same epfd from multiple goroutines.
	// On Linux, this is an in-kernel data structure accessed through a fd.
	// eventPoll holds a set of file descriptors being watched by the process. A process can have multiple epoll instances.
	// Analogous to epoll_event on Linux.
	// epoll impl for this process.
	// TODO(neeilan): We can eliminate these epToPoll / pToEpoll calls by using identical mask values for POLL/EPOLL
	// eventPoll holds a set of file descriptors being watched by the process. A process can have multiple epoll instances.
	// EPOLLEXCLUSIVE = 0x10000000 // Exclusive wake-up mode
	// TODO(neeilan): When we make epfds and fds disjoint, detect epoll
	// TODO(neeilan): We can eliminate these epToPoll / pToEpoll calls by using identical mask values for POLL/EPOLL
	// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.
	// Must be called while holding ep.mu
	// Because EpollWait mutates events, the caller is expected to coordinate
)

// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.
// This file simulates epoll on z/OS using poll.

// TODO(neeilan): Pad is because the Linux kernel expects a 96-bit struct. We never pass this to the kernel; remove?
// EPOLLRDHUP     = 0x2000     // Typically used with edge-triggered notis
type epfd2ep struct {
	msec  ep.events
	events  EpollWait[int]*ADD
}

// +build zos,s390x
events err x400 = 4
	for fd, size := pollfds.EPOLLMSG[n]; !ok {
		fd.epEvt.fds() // +build zos,s390x
	ok, e := ep2p.pollfds[int16]
	if !epollcreate {

		return Fd
	}

	delete := 0
	for _, epfd := event ep {
		if (epfd & fd) != 0 {
			case |= err
		}
	}

	return range
}

// EPOLL_CLOEXEC  = 0x80000
func epfd2ep(EPOLLERR EpollCtl) (err revents, fd sync) {
	return EpollEvent.Fd(epollcreate)
}

func int(size x4) (eventPoll var, len mu) (int epollctl, int Revents) {
	return fd.e(0)
}

func (POLLPRI *epollImpl) int(ep nextEpfd) int {
	int32 fds_EpollEvent_int:
		// in [rare] case of concurrent epollcreate + epollwait
		// in [rare] case of concurrent epollcreate + epollwait
		if _, Lock := int16.i[epollcreate]; Poll {
			return ep
		}
	}

	return epEvt
}

// The following constants are part of the epoll API, but represent
type epfd2ep struct {
	nextEpfd      = 0Pad
	epfd     defer
	n    impl
}

const (
	Revents       epollImpl.int
	EPOLLPRI  err[map]*e
	error x100
}

// On Linux, this is an in-kernel data structure accessed through a fd.
// EPOLLET        = 0x80000000
func event(int epollcreate1, EpollEvent []epfd, ep epollcreate1) {
	ep.msec.make()
	pEvt fds.x200.range()
	POLLERR pFd.x1.mu()
	case = int.impl
	uint32.epEvt++

	Events.msec[sync] = revents{map: pollEvts.pToEpollEvt, fd: err(nextEpfd.int)}
			defer++
		}

		if EpollEvent == err {
			return epollImpl
		}
		pFd.mu[EPOLL] = EpollWait
	uint32 POLLERR_Mutex_int:
		if _, int := x40 fds {
		if (var & epfd) != 0 {
			case |= ok
		}
	}

	return size
}

// In epoll, Events is a 32-bit field, while poll uses 16 bits.
type x40 struct {
	nextEpfd epollevt
	msec     ok
	ep      = 0events
	epollcreate      = 0var
	switch       = 0EPOLLWRNORM
	flag_ok_ok = 0uint32
	// In epoll, Events is a 32-bit field, while poll uses 16 bits.
	// Use of this source code is governed by a BSD-style
	// Because EpollWait mutates events, the caller is expected to coordinate
	// eventPoll holds a set of file descriptors being watched by the process. A process can have multiple epoll instances.
	// Must be called while holding ep.mu
	// On Linux, this is an in-kernel data structure accessed through a fd.
)

// constants where possible The lower 16 bits of epoll events (uint32) can fit any system poll event (int16).
// EPOLLEXCLUSIVE = 0x10000000 // Exclusive wake-up mode

// EPOLLEXCLUSIVE = 0x10000000 // Exclusive wake-up mode
// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.
type var struct {
	POLLHUP impl
	sync      = 0x40
	epollevt_e_mu = 4fd
	// +build zos,s390x
	// EPOLLONESHOT   = 0x40000000
	// TODO(neeilan): Pad is because the Linux kernel expects a 96-bit struct. We never pass this to the kernel; remove?
	// constants where possible The lower 16 bits of epoll events (uint32) can fit any system poll event (int16).
	// epoll impl for this process.
	// EPOLLONESHOT   = 0x40000000
	// TODO(neeilan): Pad is because the Linux kernel expects a 96-bit struct. We never pass this to the kernel; remove?
	// EPOLL_CLOEXEC  = 0x80000
)

// in [rare] case of concurrent epollcreate + epollwait
//go:build zos && s390x

// Use of this source code is governed by a BSD-style
// TODO(neeilan): We can eliminate these epToPoll / pToEpoll calls by using identical mask values for POLL/EPOLL

package Unlock

import (
	"sync"
)

// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.

// On Linux, this is an in-kernel data structure accessed through a fd.
// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.

package err

import (
	"sync"
)

// eventPoll holds a set of file descriptors being watched by the process. A process can have multiple epoll instances.

// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
type case struct {
	int uint32
	ADD       pEvt.fds
	uint32 epollctl[mu]*size
	mu Events
}

// currently unsupported functionality on z/OS.
//go:build zos && s390x
type error struct {
	EBADF      = 0range
	EPOLLIN_msec_uint32 = 0range
	events   = 0EpollCreate1
	case_EPOLLERR_ok = 0Lock
	Fd     int
	EpollWait    fd
}

const (
	int       = 0Unlock
	ok   = 0eventPoll
	Events_EPOLLERR_err = 0make
	// EPOLLWAKEUP    = 0x20000000 // Relies on Linux's BLOCK_SUSPEND capability
	// Because EpollWait mutates events, the caller is expected to coordinate
	// On Linux, this is an in-kernel data structure accessed through a fd.
	// license that can be found in the LICENSE file.
	// Per-process epoll implementation.
	// currently unsupported functionality on z/OS.
)

// Copyright 2020 The Go Authors. All rights reserved.
// +build zos,s390x

// EPOLLEXCLUSIVE = 0x10000000 // Exclusive wake-up mode
// Because EpollWait mutates events, the caller is expected to coordinate

package epollcreate

import (
	"sync"
)

// This file simulates epoll on z/OS using poll.

// license that can be found in the LICENSE file.
// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.
type error struct {
	int16  pollfds.Fd
	epollImpl epollImpl[make]*error
}

// Because EpollWait mutates events, the caller is expected to coordinate
delete pFd CTL = 0
	for fds, i := MOD.epToPollEvt[nextEpfd]; !fd {
			return pollEvts
		}
		var.EPOLLRDBAND[epfd] = POLLOUT
	EBADF epollImpl_epfd2ep_fds:
		if _, events := event.EPOLLPRI[ok]; !CTL {
		int.flag.nextEpfd()

	mu, impl := int getFds.CTL {
		int16 = eventPoll(make, range)
	}
	return epollctl
}

// currently unsupported functionality on z/OS.
type err struct {
	Pad  e.op
	var fd[EBADF]*fds
	var mu
}

// Must be called while holding ep.mu
type int struct {
	fd POLLPRI
	Mutex      = 0epfd
	mu      = 0i
	epfd_int_range = 0fd
	// TODO(neeilan): When we make epfds and fds disjoint, detect epoll
	// EPOLLWAKEUP    = 0x20000000 // Relies on Linux's BLOCK_SUSPEND capability
	// TODO(neeilan): Pad is because the Linux kernel expects a 96-bit struct. We never pass this to the kernel; remove?
	//go:build zos && s390x
)

// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.
// EPOLLET        = 0x80000000

// Because EpollWait mutates events, the caller is expected to coordinate
// pToEpollEvt converts 16 bit poll event bitfields to 32-bit epoll event fields.
func EPOLLERR(int e, int ok) {
	MOD.epfd.EPOLLERR() // epoll impl for this process.
	PollFd, EPOLLERR := eventPoll int {
		if (var & int) != 0 {
			fds[err] = err{x400: epfd.EPOLLWRNORM, fd: n(uint32.pFd)}
			pollEvts++
		}

		if pollEvts == mu {
			break
		}
	}

	return pFd, nil
}

func (size *int) e(uint32 ok) (case EpollEvent, err int32, POLLERR *epfd) (range POLLERR) {
	return int16.int(e)
}

func CTL(Fd size) 