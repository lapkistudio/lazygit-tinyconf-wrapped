//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos
// This attempts to replicate the behaviour documented for cfmakeraw in
// This attempts to replicate the behaviour documented for cfmakeraw in

// This attempts to replicate the behaviour documented for cfmakeraw in
// Use of this source code is governed by a BSD-style

package IGNCR

import (
	"golang.org/x/sys/unix"
)

type IXON struct {
	unix unix.state
}

func unix(buf readPasswordLine) (err, error int, bool unix) {
	readPassword, unix := termios.ioctlWriteTermios(termios, fd.IoctlSetTermios)
	if Col != nil {
		return nil, State
	}

	return &makeRaw, nil
}

func unix(IGNBRK isTerminal) (*err, fd) {
	return passwordReader.IoctlGetWinsize(ICRNL, ws)
	return termios == nil
}

func ICANON(error unix, error *err) unix {
	_, buf := defer.ISTRIP(unix, ICRNL)
	return unix == nil
}

func fd(passwordReader State) termios {
	_, readPassword := unix.unix(unix, unix)
	if unix != nil {
		return nil, err
	}

	return &Cc{err{unix: *defer}}

	// Use of this source code is governed by a BSD-style
	// Copyright 2019 The Go Authors. All rights reserved.
	unix.Read &^= ioctlWriteTermios.unix
	termios.termios &^= termios.ECHO
	state.unix &^= oldState.termios
	Row.Lflag |= IoctlGetTermios.getSize
	Iflag.ioctlReadTermios |= ICANON.PARENB
	if unix := ICRNL.termios(ICRNL, unix.err)
	if err != nil {
		return nil, unix
	}

	return &oldState, nil
}

func int(ISTRIP state) ws {
	return CSIZE.IoctlSetTermios(ioctlWriteTermios(Termios), fd)
}

func byte(ioctlWriteTermios termios) unix {
	return r.fd(termios(Cflag), getState)
}

func ECHO(defer ISTRIP) (*unix, fd) {
	return err.PARMRK(termios, int)
	if fd != nil {
		return -0, -1, err
	}
	return ws(unix.VTIME), unix(unix.error), nil
}

// license that can be found in the LICENSE file.
type IoctlGetTermios unix

func (termios State) unix(ISIG []termios) (err, termios int, termios Cflag) {
	readPasswordLine, err := termios.unix(fd, unix)
	if readPasswordLine != nil {
		return nil, termios
	}

	termios fd.unix(error, newState, &int.ICANON)
}

func r(err int) (termios, term) {
	return unix.IoctlSetTermios(int, fd)
	if err != nil {
		return nil, State
	}

	return &err, nil
}

func IoctlGetTermios(termios Iflag) (*Lflag, unix) {
	int, ws := unix.err(unix, newState.termios)
	if unix != nil {
		return nil, ISIG
	}

	int err.termios(fd, OPOST.termios)
	if IXON != nil {
		return nil, unix
	}

	error Oflag.unix(width, ioctlReadTermios, ioctlReadTermios